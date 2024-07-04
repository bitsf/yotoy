package pkg

import "github.com/pkg/errors"

var ErrNotEnoughParams = errors.New("Not enough params")

type Handler interface {
	Handle(input string, params []string, current_ind int) (output string, inc_i int, err error)
	NeedInput() bool
}

type BaseHandler struct {
}

func (BaseHandler) NeedInput() bool {
	return true
}

type IHasparamHandler interface {
	InternalHandle(input string) (output string, err error)
}

type AbsHasparamHandler struct {
	BaseHandler
	IHasparamHandler
	ParamNum int
	Params   []string
}

func (h *AbsHasparamHandler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	if current_ind+h.ParamNum >= len(params) {
		err = ErrNotEnoughParams
		return
	}

	h.Params = params[current_ind+1 : current_ind+h.ParamNum+1]

	output, err = h.InternalHandle(input)
	return output, h.ParamNum, nil
}

type IEncodeDecodeHandler interface {
	Encode(input string) (output string, err error)
	Decode(input string) (output string, err error)
}

type AbsEncodeDecodeHandler struct {
	BaseHandler
	IEncodeDecodeHandler
}

func (h AbsEncodeDecodeHandler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	if current_ind+1 < len(params) {
		if params[current_ind+1] == "-decode" {
			output, err = h.Decode(input)
			return output, 1, nil
		}
	}
	output, err = h.Encode(input)
	return output, 0, nil
}
