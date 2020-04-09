package raindrops

import "strconv"

func Convert(input int) string {
	var toReturn string

	if input%3 == 0 {
		toReturn += "Pling"
	}

	if input%5 == 0 {
		toReturn += "Plang"
	}

	if input%7 == 0 {
		toReturn += "Plong"
	}

	if toReturn == "" {
		toReturn = strconv.Itoa(input)
	}

	return toReturn
}
