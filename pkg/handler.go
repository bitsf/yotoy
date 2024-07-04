package pkg

type Handler interface {
	Handle(input string, params []string, current_ind int) (output string, inc_i int, err error)
	NeedInput() bool
}

type BaseHandler struct {
}

func (BaseHandler) NeedInput() bool {
	return true
}

type IEncodeDecodeHandler interface {
	Encode(input string) (output string, err error)
	Decode(input string) (output string, err error)
}

type EncodeDecodeHandler struct {
	BaseHandler
	IEncodeDecodeHandler
}

func (h EncodeDecodeHandler) Handle(input string, params []string, current_ind int) (output string, inc_i int, err error) {
	if current_ind+1 < len(params) {
		if params[current_ind+1] == "-decode" {
			output, err = h.Decode(input)
			return output, 1, nil
		}
	}
	output, err = h.Encode(input)
	return output, 0, nil
}
