package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	testCases := []struct {
		desc   string
		args   []string
		expect string
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
	}

	var buf bytes.Buffer
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			os.Args = tc.args
			originalStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			outC := make(chan struct{})

			main()

			go func() {
				io.Copy(&buf, r)
				outC <- struct{}{}
			}()
			w.Close()
			<-outC
			os.Stdout = originalStdout
			if buf.String() != tc.expect {
				t.Errorf("expected %s, got %s", tc.expect, buf.String())
			}
		})
	}
}
