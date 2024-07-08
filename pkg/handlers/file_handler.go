package handlers

import (
	"io"
	"os"
	"yotoy/pkg"
)

type FileHandler struct {
	pkg.AbsHasparamHandler
}

func (h FileHandler) InternalHandle(input string) (output string, err error) {
	f, _err := os.OpenFile(h.Params[0], os.O_RDONLY, 0644)
	if _err != nil {
		err = _err
		return
	}
	bytes, _err := io.ReadAll(f)
	if _err != nil {
		err = _err
		return
	}
	output = string(bytes)
	return
}

func (h FileHandler) NeedInput() bool {
	return false
}

func init() {
	pkg.Register("file", &FileHandler{})
}
