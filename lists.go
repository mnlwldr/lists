package lists

// Head return the first value of the given list
func Head[T any](list []T) []T {
	return list[0:1]
}

// Tail return all values of the given list, 
// except the first one.
func Tail[T any](list []T) []T {
	return list[1:]
}
