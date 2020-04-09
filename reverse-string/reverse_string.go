package reverse

func Reverse(word string) (toReturn string) {
	for _, str := range word {
		toReturn = string(str) + toReturn
	}

	return
}
