package collections

// HasInt32 returns true if list contains search
func HasInt32(list []int32, search int32) bool {
	for _, item := range list {
		if item == search {
			return true
		}
	}

	return false
}

// UniqueIntegers32 returns list without any duplicates in it.
func UniqueIntegers32(list []int32) []int32 {
	index := map[int32]bool{}
	result := []int32{}
	for _, item := range list {
		if !index[item] {
			index[item] = true
			result = append(result, item)
		}
	}

	return result
}

// CompareInts32 returns true if the values of both lhs & rhs are the same.
func CompareInts32(lhs, rhs []int32) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for index := range lhs {
		if lhs[index] != rhs[index] {
			return false
		}
	}
	return true
}
