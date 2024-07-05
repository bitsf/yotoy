package handlers

import (
	"encoding/json"
	"yytoy/pkg"

	"gopkg.in/yaml.v2"
)

type JsonHandler struct {
	pkg.AbsCommonHandler
}

func (h *JsonHandler) InternalHandle(input string) (output string, err error) {
	var result map[string]interface{}
	err = yaml.Unmarshal([]byte(input), &result)
	if err != nil {
		return
	}

	resultBytes, _err := json.MarshalIndent(result, "", "  ")
	if _err != nil {
		err = _err
		return
	}
	output = string(resultBytes)
	return
}

func NewJsonHandler() *JsonHandler {
	h := &JsonHandler{pkg.AbsCommonHandler{}}
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("json", NewJsonHandler())
}
