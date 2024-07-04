package handlers

import (
	"godevtoy/pkg"
	"strings"
)

type TrimHandler struct {
	pkg.BaseHandler
}

func (s TrimHandler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	output = strings.TrimSpace(input)
	return output, 0, nil
}

func init() {
	pkg.Register("trim", &TrimHandler{})
}
