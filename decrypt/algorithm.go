package decrypt

func Nimbus(str string) string {
	dencryptedString := ""

	for _, c := range str {
		// Fetch the ASCII code of the character
		asciiCode := int(c)
		// Create a character which is 3 places forward than the existing character
		character := string(asciiCode - 3)
		dencryptedString += character
	}

	return dencryptedString
}
