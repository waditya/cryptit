package encrypt

func Nimbus(str string) string {
	encryptedString := ""

	for _, c := range str {
		// Fetch the ASCII code of the character
		asciiCode := int(c)
		// Create a character which is 3 places forward than the existing character
		character := string(asciiCode + 3)
		encryptedString += character
	}

	return encryptedString
}
