package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"godevtoy/pkg"
)

type Md5Handler struct {
	pkg.AbsCommonHandler
}

func (h *Md5Handler) InternalHandle(input string) (output string, err error) {
	hash := md5.Sum([]byte(input))
	output = hex.EncodeToString(hash[:])
	return
}

func NewMd5Handler() *Md5Handler {
	h := &Md5Handler{pkg.AbsCommonHandler{}}
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("md5", NewMd5Handler())
}
