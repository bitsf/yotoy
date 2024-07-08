package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"

	"yotoy/pkg"
	_ "yotoy/pkg/handlers"
)

func main() {
	factory := initHandlerFactory()
	result := ""
	var err error
	var inc_i int
	if len(os.Args) < 2 {
		panic("no params")
	}
	params := os.Args[1:]
	for i := 0; i < len(params); i++ {
		if strings.HasPrefix(params[i], "-") {
			cmd := strings.TrimPrefix(params[i], "-")
			handler, ok := factory[cmd]
			if !ok {
				panic(errors.Errorf("Bad command %s", cmd))
			}
			if i == 0 && handler.NeedInput() && result == "" {
				cmd = "stdin"
				handler = factory[cmd]
				i--
			}
			result, inc_i, err = handler.Handle(result, params, i)
			if err != nil {
				panic(err)
			}
			i += inc_i
		} else {
			panic(errors.Errorf("Bad command %s", params[i]))
		}
	}
	fmt.Print(result)
}

func initHandlerFactory() map[string]pkg.Handler {
	return pkg.GetFactory()
}
