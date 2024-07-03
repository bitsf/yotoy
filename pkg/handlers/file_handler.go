package handlers

import (
	"godevtoy/pkg"
	"io"
	"os"
)

type FileHandler struct {
	pkg.BaseHandler
}

func (s FileHandler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	h, _err := os.OpenFile(params[current_ind+1], os.O_RDONLY, 0644)
	if _err != nil {
		err = _err
		return
	}
	bytes, _err := io.ReadAll(h)
	if _err != nil {
		err = _err
		return
	}
	inc_i = 1
	output = string(bytes)
	return
}

func (s FileHandler) NeedInput() bool {
	return false
}

func init() {
	pkg.Register("file", &FileHandler{})
}
