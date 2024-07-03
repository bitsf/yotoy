package handlers

import (
	"encoding/base64"
	"godevtoy/pkg"
)

type Base64Handler struct {
	pkg.BaseHandler
}

func (s Base64Handler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	output = base64.StdEncoding.EncodeToString([]byte(input))
	return
}

func init() {
	pkg.Register("base64", &Base64Handler{})
}
