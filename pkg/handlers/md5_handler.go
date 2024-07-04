package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"godevtoy/pkg"
)

type Md5Handler struct {
	pkg.BaseHandler
}

func (s Md5Handler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	hash := md5.Sum([]byte(input))
	output = hex.EncodeToString(hash[:])
	return
}

func init() {
	pkg.Register("md5", &Md5Handler{})
}
