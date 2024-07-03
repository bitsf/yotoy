package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	handler := initHandler()
	result := ""
	var err error
	var inc_i int
	if len(os.Args) < 2 {
		panic("no params")
	}
	params := os.Args[1:]
	if params[0] != "-stdin" && params[0] != "-file" && params[0] != "-str" {
		params = append([]string{"-stdin"}, params...)
	}
	for i := 0; i < len(params); i++ {
		if strings.HasPrefix(params[i], "-") {
			cmd := strings.TrimPrefix(params[i], "-")
			result, inc_i, err = handler[cmd](result, params, i)
			if err != nil {
				panic(err)
			}
			i += inc_i
		}
	}
	fmt.Print(result)
}

type Handler func(input string, params []string, current_ind int) (output string, inc_i int, err error)

func initHandler() map[string]Handler {
	stdinHandler := func(input string, params []string, current_ind int) (output string, inc_i int, err error) {
		bytes, _err := io.ReadAll(os.Stdin)
		if _err != nil {
			err = _err
			return
		}
		output = string(bytes)
		return
	}

	strHandler := func(input string, params []string, current_ind int) (output string, inc_i int, err error) {
		output, inc_i = params[current_ind+1], 1
		return
	}

	jsonHandler := func(input string, params []string, current_ind int) (output string, inc_i int, err error) {
		var result map[string]interface{}
		err = json.Unmarshal([]byte(input), &result)
		if err != nil {
			return
		}
		// If you want to return the map as a string, you can marshal it back to JSON
		resultBytes, _err := json.MarshalIndent(result, "", "  ")
		if _err != nil {
			err = _err
			return
		}
		output = string(resultBytes)
		return
	}

	return map[string]Handler{
		"stdin": stdinHandler,
		"str":   strHandler,
		"json":  jsonHandler,
	}

}
