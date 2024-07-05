package handlers

import (
	"yytoy/pkg"

	"gopkg.in/yaml.v2"
)

type YamlHandler struct {
	pkg.AbsCommonHandler
}

func (h *YamlHandler) InternalHandle(input string) (output string, err error) {
	var data map[string]interface{}
	err = yaml.Unmarshal([]byte(input), &data)
	if err != nil {
		return
	}

	result, _err := yaml.Marshal(data)
	if _err != nil {
		err = _err
		return
	}
	output = string(result)
	return
}

func NewYamlHandler() *YamlHandler {
	h := &YamlHandler{pkg.AbsCommonHandler{}}
	h.IHasparamHandler = h
	return h
}

func init() {
	pkg.Register("yaml", NewYamlHandler())
}
