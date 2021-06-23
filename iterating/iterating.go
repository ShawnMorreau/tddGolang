package iterating

func Iterate(char string, amount int) string {
	var returnString string
	for i := 0; i < amount; i++ {
		returnString += char
	}
	return returnString
}
