package handlers

import (
	"godevtoy/pkg"
	"io"
	"os"
)

type StdinHandler struct {
	pkg.BaseHandler
}

func (s StdinHandler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	bytes, _err := io.ReadAll(os.Stdin)
	if _err != nil {
		err = _err
		return
	}
	output = string(bytes)
	return
}

func (s StdinHandler) NeedInput() bool {
	return false
}

func init() {
	pkg.Register("stdin", &StdinHandler{})
}
