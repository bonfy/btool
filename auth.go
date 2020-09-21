package btool

func CheckRight(inputCode, rightCode int) bool {
	return inputCode&rightCode == rightCode
}
