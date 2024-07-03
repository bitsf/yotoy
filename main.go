package main

import (
	"fmt"
	"godevtoy/pkg"
	"os"
	"strings"

	_ "godevtoy/pkg/handlers"
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
			result, inc_i, err = handler[cmd].Handle(result, params, i)
			if err != nil {
				panic(err)
			}
			i += inc_i
		}
	}
	fmt.Print(result)
}

func initHandler() map[string]pkg.Handler {
	return pkg.GetFactory()
}
