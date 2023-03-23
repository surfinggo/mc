package mc

func SliceContains[V comparable](slice []V, v V) bool {
	for _, s := range slice {
		if s == v {
			return true
		}
	}
	return false
}

// PointerTo creates a pointer to bool
func PointerTo[V any](v V) *V {
	return &v
}

func PtrTo[V any](v V) *V {
	return PointerTo(v)
}

// ValueOr returns the default value if the value is empty
func ValueOr[V comparable](v V, defaultValue V) V {
	zero := new(V)
	if v == *zero {
		return defaultValue
	}
	return v
}

func VarOr[V comparable](v V, defaultValue V) V {
	return ValueOr(v, defaultValue)
}
