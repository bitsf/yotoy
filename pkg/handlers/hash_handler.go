package handlers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"godevtoy/pkg"
)

type HashHandler struct {
	pkg.AbsHasparamHandler
}

func (h HashHandler) InternalHandle(input string) (output string, err error) {
	switch h.Params[0] {
	case "md5":
		hash := md5.Sum([]byte(input))
		output = hex.EncodeToString(hash[:])
	case "sha1":
		hash := sha1.Sum([]byte(input))
		output = hex.EncodeToString(hash[:])
	case "sha256", "sha":
		hash := sha256.Sum256([]byte(input))
		output = hex.EncodeToString(hash[:])
	case "sha512":
		hash := sha512.Sum512([]byte(input))
		output = hex.EncodeToString(hash[:])
	default:
		err = pkg.ErrInvalidParam
	}

	return
}

func NewHashHandler() *HashHandler {
	h := &HashHandler{pkg.AbsHasparamHandler{}}
	h.ParamNum = 1
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("hash", NewHashHandler())
}
