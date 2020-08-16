package accumulate

//Accumulate iterates through a collection, applies the results of the passed converter function for each, and returns them in an array of the same size.
func Accumulate(collection []string, converter func(string) string) []string {
	accumulation := make([]string, len(collection))
	for i, value := range collection {
		accumulation[i] = converter(value)
	}
	return accumulation
}
