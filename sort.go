package sort

// Collection ... is a data structure interface defines the behaviours required to sort the data structure
type Collection interface {
	// Len ... returns the length of data structure
	Len() int
	// Greater ... tells if the (i)th element is greater than the (j)th element
	Greater(i, j int) bool
	// Swap ... swaps the elements present in i and j positions
	Swap(i, j int)
}
