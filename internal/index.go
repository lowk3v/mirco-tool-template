package internal

type Options struct {
	Verbose bool
	Output  string
}

func New(opt *Options) *Options {
	return opt
}

func (opt *Options) Run() {
	// TODO
}
