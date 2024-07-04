package handlers

import (
	"godevtoy/pkg"
	"strings"
)

type TrimHandler struct {
	pkg.AbsCommonHandler
}

func (h *TrimHandler) InternalHandle(input string) (output string, err error) {
	output = strings.TrimSpace(input)
	return
}

func NewTrimHandler() *TrimHandler {
	h := &TrimHandler{pkg.AbsCommonHandler{}}
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("trim", NewTrimHandler())
}
