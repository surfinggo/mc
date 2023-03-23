package mc

// Must check a value to be nil, panics if not
func Must(e any) {
	if e != nil {
		panic(e)
	}
}
