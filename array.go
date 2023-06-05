package gotils

func Reverse[S ~[]E, E any](arr S) {
	for i, j := 0, len(arr)-1; i < len(arr)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func Contains[S comparable](arr []S, item S) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}
	return false
}

func ContainsAny[S comparable](arr []S, items ...S) bool {
	for _, i := range items {
		if Contains(arr, i) {
			return true
		}
	}
	return false
}

func ContainsSameElements[S comparable](a, b []S) bool {
	if len(a) != len(b) {
		return false
	}

	tmp := make(map[S]struct{})
	for _, el := range a {
		tmp[el] = struct{}{}
	}

	for _, el := range b {
		if _, ok := tmp[el]; !ok {
			return false
		}
	}

	return true

}
