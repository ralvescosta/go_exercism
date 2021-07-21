package sublist

type Relation string

func Sublist(a, b []int) Relation {
	r := "sublist"
	aLen := len(a)
	bLen := len(b)

	if aLen == bLen {
		r = "equal"
		for i, number := range a {
			if number != b[i] {
				return Relation("unequal")
			}
		}
	}

	if aLen > bLen {
		if bLen == 0 {
			return Relation("superlist")
		}
		initialIndex := findIndex(a, b[0])
		if initialIndex == -1 {
			return Relation("unequal")
		}
		finalIndex := findIndex(a, b[len(b)-1])
		if finalIndex == -1 {
			return Relation("unequal")
		}
		aSlice := a[initialIndex:finalIndex]
		r = "superlist"
		for index, number := range aSlice {
			if number != b[index] {
				r = "unequal"
				break
			}
		}
	}

	if bLen > aLen {
		if aLen == 0 {
			return Relation("sublist")
		}

		initialIndex := findIndex(b, a[0])
		if initialIndex == -1 {
			return Relation("unequal")
		}
		finalIndex := findIndex(b, a[len(a)-1])
		if finalIndex == -1 {
			return Relation("unequal")
		}
		bSlice := b[initialIndex : finalIndex+1]
		r = "sublist"
		for index, number := range bSlice {
			if number != a[index] {
				r = "unequal"
				break
			}
		}
	}

	return Relation(r)
}

func findIndex(slice []int, search int) int {
	for i := len(slice) - 1; i > 0; i-- {
		if slice[i] == search {
			return i
		}
	}
	return -1
}
