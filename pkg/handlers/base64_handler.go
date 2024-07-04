package handlers

import (
	"encoding/base64"
	"godevtoy/pkg"
)

type Base64Handler struct {
	pkg.EncodeDecodeHandler
}

func (s Base64Handler) Encode(input string) (output string, err error) {
	output = base64.StdEncoding.EncodeToString([]byte(input))
	return
}

func (s Base64Handler) Decode(input string) (output string, err error) {
	ss, _err := (base64.StdEncoding.DecodeString(input))
	if _err != nil {
		err = _err
		return
	}

	return string(ss), nil
}

func NewBase64Handler() Base64Handler {
	h := Base64Handler{pkg.EncodeDecodeHandler{}}
	h.IEncodeDecodeHandler = h
	return h
}

func init() {
	pkg.Register("base64", NewBase64Handler())
}
