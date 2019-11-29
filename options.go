package rate

// Options Rate Package config
type Options struct {
	Max          int
}

// Option opt func
type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		Max: 1,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Max max
func Max(max int) Option {
	return func(o *Options) {
		o.Max = max
	}
}
