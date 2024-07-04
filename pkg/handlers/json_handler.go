package handlers

import (
	"encoding/json"
	"godevtoy/pkg"
	"gopkg.in/yaml.v2"
)

type JsonHandler struct {
	pkg.BaseHandler
}

func (s JsonHandler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
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

func init() {
	pkg.Register("json", &JsonHandler{})
}
