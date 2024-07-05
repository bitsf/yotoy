package handlers

import "yytoy/pkg"

type StrHandler struct {
	pkg.AbsHasparamHandler
}

func (h *StrHandler) InternalHandle(input string) (output string, err error) {
	output = h.Params[0]
	return
}

func (s StrHandler) NeedInput() bool {
	return false
}

func NewStrHandler() *StrHandler {
	h := &StrHandler{pkg.AbsHasparamHandler{}}
	h.ParamNum = 1
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("str", NewStrHandler())
}
