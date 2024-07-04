package handlers

import (
	"godevtoy/pkg"
	"gopkg.in/yaml.v2"
)

type YamlHandler struct {
	pkg.BaseHandler
}

func (s YamlHandler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	var data map[string]interface{}
	err = yaml.Unmarshal([]byte(input), &data)
	if err != nil {
		return "", 0, err
	}

	result, _err := yaml.Marshal(data)
	if _err != nil {
		err = _err
		return
	}
	output = string(result)
	return
}

func init() {
	pkg.Register("yaml", &YamlHandler{})
}
