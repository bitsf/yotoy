package handlers

import (
	"encoding/base64"
	"godevtoy/pkg"
)

type Base64Handler struct {
	pkg.BaseHandler
}

func (s Base64Handler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	if current_ind+1 < len(params) {
		if params[current_ind+1] == "-decode" {
			ss, _err := (base64.StdEncoding.DecodeString(input))
			if _err != nil {
				err = _err
				return
			}
			return string(ss), 1, nil
		}
	}
	output = base64.StdEncoding.EncodeToString([]byte(input))
	return
}

func init() {
	pkg.Register("base64", &Base64Handler{})
}
