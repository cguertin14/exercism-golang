package sublist

// equal, unequal, sublist or superlist
type Relation string

func isSublist(listOne, listTwo []int) bool {
	lenA, lenB := len(listOne), len(listTwo)

	for i := 0; i <= lenB-lenA; i++ {
		if areListsEqual(listOne, listTwo[i:i+lenA]) {
			return true
		}
	}

	return false
}

func areListsEqual(listOne, listTwo []int) bool {
	if len(listOne) != len(listTwo) {
		return false
	}

	for i := 0; i < len(listOne); i++ {
		if listOne[i] != listTwo[i] {
			return false
		}
	}

	return true
}

func Sublist(listOne, listTwo []int) (relation Relation) {
	lenA, lenB := len(listOne), len(listTwo)

	switch {
	case lenA == lenB:
		areEqual := areListsEqual(listOne, listTwo)
		if areEqual {
			relation = "equal"
		} else {
			relation = "unequal"
		}
		break
	case lenA > lenB:
		isSuperlist := isSublist(listTwo, listOne)
		if isSuperlist {
			relation = "superlist"
		} else {
			relation = "unequal"
		}
		break
	case lenA < lenB:
		isSublistOfB := isSublist(listOne, listTwo)
		if isSublistOfB {
			relation = "sublist"
		} else {
			relation = "unequal"
		}
		break
	}

	return
}
