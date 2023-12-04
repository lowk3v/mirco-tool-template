package internal

type Options struct {
	Verbose bool
	Version bool
	Output  string
}

func New(opt *Options) *Options {
	return opt
}

func (opt *Options) Run() {
	// TODO
}
