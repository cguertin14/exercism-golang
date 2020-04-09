package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func reverse(word string) (toReturn string) {
	for _, str := range word {
		toReturn = string(str) + toReturn
	}

	return
}

func containsNonDigitsChars(code string) bool {
	for _, char := range code {
		if !unicode.IsDigit(char) {
			return true
		}
	}

	return false
}

func doubleSecondDigits(code *string) {
	var sb strings.Builder

	for index, digit := range reverse(*code) {
		if index%2 != 0 {
			value, _ := strconv.Atoi(string(digit))
			doubled := value * 2

			if doubled > 9 {
				doubled -= 9
			}

			sb.WriteString(strconv.Itoa(doubled))
		} else {
			sb.WriteString(string(digit))
		}
	}

	*code = sb.String()
}

func calculateSum(code string) (sum int) {
	doubleSecondDigits(&code)

	for _, digit := range code {
		value, _ := strconv.Atoi(string(digit))
		sum += value
	}

	return
}

func Valid(code string) bool {
	code = strings.Replace(code, " ", "", -1)
	if len(code) <= 1 {
		return false
	}

	if containsNonDigitsChars(code) {
		return false
	}

	sum := calculateSum(code)

	return sum%10 == 0
}
