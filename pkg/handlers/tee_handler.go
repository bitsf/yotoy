package handlers

import (
	"godevtoy/pkg"
	"os"
)

type TeeHandler struct {
	pkg.AbsHasparamHandler
}

func (h *TeeHandler) InternalHandle(input string) (output string, err error) {
	f, _err := os.OpenFile(h.Params[0], os.O_CREATE|os.O_WRONLY, 0644)
	if _err != nil {
		err = _err
		return
	}
	_, err = f.WriteString(input)
	f.Close()
	output = input
	return
}

func NewTeeHandler() *TeeHandler {
	h := &TeeHandler{pkg.AbsHasparamHandler{}}
	h.ParamNum = 1
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("tee", NewTeeHandler())
}
