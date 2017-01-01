package main

// ChangeType typed constants
type ChangeType string

const (
	// Added "added"
	Added ChangeType = "added"
	// Deleted "deleted"
	Deleted ChangeType = "deleted"
)

// CompareResult represents compare result
type CompareResult struct {
	StoreName string
	Change    ChangeType
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// CompareAvailability compares availability
func CompareAvailability(currentStoreNames []string, newStoreNames []string) []CompareResult {
	res := []CompareResult{}
	for _, n := range newStoreNames {
		if !contains(currentStoreNames, n) {
			res = append(res, CompareResult{n, Added})
		}
	}
	for _, n := range currentStoreNames {
		if !contains(newStoreNames, n) {
			res = append(res, CompareResult{n, Deleted})
		}
	}
	return res
}
