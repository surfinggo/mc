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
