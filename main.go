package main

import (
	"fmt"
	"strings"
)

const originalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetters(key int, letters string) (result string) {
	runes := []rune(letters)

	fh := string(runes[:len(runes)-key])
	sh := string(runes[len(runes)-key:])
	// result = append(result, sh...)
	// result = sh+fh
	result = fmt.Sprintf("%s%s", sh, fh)
	return
}
func encrypt(key int, plainText string) string {
	hashLetters := hashLetters(key, originalLetters)
	var hashedString = ""

	findOne := func(r rune) rune {
		pos := strings.Index(originalLetters, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetters)) % len(originalLetters)
			hashedString = hashedString + string(hashLetters[letterPosition])
		}
		return r
	}

	strings.Map(findOne, plainText)
	return hashedString

}
func decrypt(key int, encryptedText string) string {
	hashLetters := hashLetters(5, originalLetters)
	var hashedString = ""
	findOne := func(r rune) rune {
		pos := strings.Index(hashLetters, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetters)) % len(originalLetters)
			hashedString = hashedString + string(originalLetters[letterPosition])
		}
		return r
	}

	strings.Map(findOne, encryptedText)
	return hashedString

}

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain Text: ", plainText)
	encrypted := encrypt(5, plainText)
	fmt.Println("Encrypted Text: ", encrypted)
	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted Text: ", decrypted)
}
