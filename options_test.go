package envconfig

import "testing"

func TestWithSplitWords(t *testing.T) {
	options := Options{AutoSplitWords: false}

	WithAutoSplitWords(&options)

	if !options.AutoSplitWords {
		t.Errorf("expected %t, got %v", true, options.AutoSplitWords)
	}
}
