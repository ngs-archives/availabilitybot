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
	Store  Store
	Change ChangeType
}

func contains(s []Store, e Store) bool {
	for _, a := range s {
		if a.Name == e.Name {
			return true
		}
	}
	return false
}

// CompareAvailability compares availability
func CompareAvailability(currentStoreNames []Store, newStoreNames []Store) []CompareResult {
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
