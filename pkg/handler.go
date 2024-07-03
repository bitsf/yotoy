package pkg

type Handler interface {
	Handle(input string, params []string, current_ind int) (output string, inc_i int, err error)
}
