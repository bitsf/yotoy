package handlers

import "godevtoy/pkg"

type StrHandler struct {
}

func (h *StrHandler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	output, inc_i = params[current_ind+1], 1
	return
}

func init() {
	pkg.Register("str", &StrHandler{})
}
