package handlers

import (
	"godevtoy/pkg"
	"strings"
)

type UpperHandler struct {
	pkg.AbsCommonHandler
}

func (h *UpperHandler) InternalHandle(input string) (output string, err error) {
	output = strings.ToUpper(input)
	return
}

func NewUpperHandler() *UpperHandler {
	h := &UpperHandler{pkg.AbsCommonHandler{}}
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("upper", NewUpperHandler())
}
