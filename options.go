package envconfig

type Option func(options *Options)

type Options struct {
	AutoSplitWords bool
}

func WithAutoSplitWords(options *Options) {
	options.AutoSplitWords = true
}
