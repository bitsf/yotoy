package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	testCases := []struct {
		desc       string
		args       []string
		expect     string
		expectFunc func(result string) error
	}{
		{
			desc: "format raw string",
			args: []string{"main.go", "-str", `
{
  "a": "11",
"b": "22"
}
`, "-json"},
			expect: `{
  "a": "11",
  "b": "22"
}`,
		},
		{
			desc:   "base64 encode",
			args:   []string{"main.go", "-str", `abc`, "-base64"},
			expect: "YWJj",
		},
		{
			desc:   "base64 decode",
			args:   []string{"main.go", "-str", `YWJj`, "-base64", "-decode"},
			expect: "abc",
		},
		{
			desc:   "base64 decode then encode",
			args:   []string{"main.go", "-str", `YWJj`, "-base64", "-decode", "-base64"},
			expect: "YWJj",
		},
		{
			desc:   "base64 encode then decode",
			args:   []string{"main.go", "-str", `def`, "-base64", "-base64", "-decode"},
			expect: "def",
		},
		{
			desc:   "md5 encode",
			args:   []string{"main.go", "-str", `abc`, "-hash", "md5"},
			expect: "900150983cd24fb0d6963f7d28e17f72",
		},
		{
			desc: "yaml format",
			args: []string{"main.go", "-str", `
a: 1
b:
  c: 2
  d: 3
e:
  - 1
  - 2
`, "-yaml", "-trim"},
			expect: `a: 1
b:
  c: 2
  d: 3
e:
- 1
- 2`,
		},
		{
			desc: "json to yaml",
			args: []string{"main.go", "-str", `
{
  "a": "11",
  "b": "22"
}
`, "-yaml"},
			expect: `a: "11"
b: "22"
`,
		},
		{
			desc: "json to yaml to json",
			args: []string{"main.go", "-str", `
{
  "a": "11",
  "b": "22"
}
`, "-yaml", "-json"},
			expect: `{
  "a": "11",
  "b": "22"
}`,
		},
		{
			desc: "tee",
			args: []string{"main.go", "-str", `def`, "-tee", "/tmp/godevtoy.tmp"},
			expectFunc: func(result string) error {
				if f, err := os.OpenFile("/tmp/godevtoy.tmp", os.O_RDONLY, 0644); err != nil {
					return err
				} else {
					defer os.Remove("/tmp/godevtoy.tmp")
					defer f.Close()
					s, err := io.ReadAll(f)
					if err != nil {
						return err
					}
					if result != string(s) {
						return fmt.Errorf("expected %s, got %s", result, s)
					}
					return nil
				}

			},
		},
		{
			desc:   "upper lower",
			args:   []string{"main.go", "-str", `def`, "-upper", "-lower", "-upper"},
			expect: "DEF",
		},
	}

	for _, _tc := range testCases {
		tc := _tc
		t.Run(tc.desc, func(t *testing.T) {
			os.Args = tc.args
			originalStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			outC := make(chan struct{})

			main()
			var buf bytes.Buffer
			go func() {
				io.Copy(&buf, r)
				outC <- struct{}{}
			}()
			w.Close()
			<-outC
			os.Stdout = originalStdout
			if tc.expectFunc != nil {
				if err := tc.expectFunc(buf.String()); err != nil {
					t.Error(err)
				}
			} else if buf.String() != tc.expect {
				t.Errorf("expected %s, got %s", tc.expect, buf.String())
			}
		})
	}
}
