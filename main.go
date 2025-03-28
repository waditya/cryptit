package main

import (
	"fmt"

	"github.com/waditya/cryptit/decrypt"
	"github.com/waditya/cryptit/encrypt"
)

func main() {
	encryptedString := encrypt.Nimbus("KodeKloud") // Encrypt and save it
	fmt.Println("Encrypted String is : ", encryptedString)
	fmt.Println("Decrypted String is : ", decrypt.Nimbus(encryptedString))
}
