package btool

func AuthCheckRight(inputCode, rightCode int) bool {
	return inputCode&rightCode == rightCode
}
