package handlers

import (
	"strings"
	"yotoy/pkg"
)

type LowerHandler struct {
	pkg.AbsCommonHandler
}

func (h *LowerHandler) InternalHandle(input string) (output string, err error) {
	output = strings.ToUpper(input)
	return
}

func NewLowerHandler() *LowerHandler {
	h := &LowerHandler{pkg.AbsCommonHandler{}}
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("lower", NewLowerHandler())
}
