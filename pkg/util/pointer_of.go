package util

func PointerOf[T any](o T) *T {
	return &o
}
