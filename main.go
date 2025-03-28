package main

import (
	"fmt"

	"github.com/waditya/cryptit/decrypt"
	"github.com/waditya/cryptit/encrypt"
)

func main() {
	encryptedString := encrypt.Nimbus("KodeKloud")
	fmt.Println("Encrypted String is : ", encryptedString)
	fmt.Println("Decrypted String is : ", decrypt.Nimbus(encryptedString))
}
