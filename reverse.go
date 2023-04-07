package gotils

func Reverse[S ~[]E, E any](arr S) {
	for i, j := 0, len(arr)-1; i < len(arr)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
