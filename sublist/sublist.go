package sublist

type Relation string

func Sublist(a, b []int) Relation {
	isASublist := verifyWhichIsSublist(a, b)
	isBSuperlist := verifyWhichIsSublist(b, a)
	if isASublist && isBSuperlist {
		return "equal"
	} else if isASublist {
		return "sublist"
	} else if isBSuperlist {
		return "superlist"
	} else {
		return "unequal"
	}
}

func verifyWhichIsSublist(a, b []int) bool {
	if len(b) < len(a) {
		return false
	} else if len(a) == 0 {
		return true
	}

	index1 := 0
	index2 := 0
	cursor2 := 0
	match := false
	for index1 < len(a) {
		v1 := a[index1]
		if index2 == len(b) && index1 < len(a) {
			match = false
			break
		}

		for index2 < len(b) {
			v2 := b[index2]
			if v1 == v2 {
				index1++
				index2++
				match = true
				break
			} else if match {
				index1 = 0
				cursor2++
				index2 = cursor2
				break
			} else {
				index2++
			}
		}

		if !match {
			break
		}

	}
	return match
}
