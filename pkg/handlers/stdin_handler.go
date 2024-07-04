package handlers

import (
	"godevtoy/pkg"
	"io"
	"os"
)

type StdinHandler struct {
	pkg.AbsCommonHandler
}

func (h *StdinHandler) InternalHandle(input string) (output string, err error) {
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

func NewStdinHandler() *StdinHandler {
	h := &StdinHandler{pkg.AbsCommonHandler{}}
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("stdin", NewStdinHandler())
}
