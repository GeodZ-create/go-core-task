package main_1

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	var a int = 1
	var b float64 = 1.0
	var c string = "Golang"
	var d bool = true
	var e complex64 = 0 + 0i
	getType(a)
	getType(b)
	getType(c)
	getType(d)
	getType(e)
	fullString := valueToString(a, b, c, d, e)
	fmt.Println(fullString)
	fullStringRune := stringToRune(fullString)
	fmt.Println(string(putSalt(fullStringRune)))
	fullStringWithSalt := putSalt(fullStringRune)
	fmt.Println(fullStringWithSalt)
	fullStringHash := getHash(fullStringWithSalt)
	fmt.Println(fullStringHash)
}

func getType(value any) string {
	result := fmt.Sprintf("%T", value)
	fmt.Println(result)
	return result
}

func valueToString(a int, b float64, c string, d bool, e complex64) string {
	aString := fmt.Sprint(a)
	bString := fmt.Sprint(b)
	cString := fmt.Sprint(c)
	dString := fmt.Sprint(d)
	eString := fmt.Sprint(e)
	fullString := aString + bString + cString + dString + eString
	return fullString
}

func stringToRune(fullString string) []rune {
	return []rune(fullString)
}

func putSalt(fullString []rune) []rune {
	salt := []rune("go-2024")
	middle := len(fullString) / 2
	left := fullString[:middle]
	right := fullString[middle:]
	var result []rune
	result = append(result, left...)
	result = append(result, salt...)
	result = append(result, right...)
	return result
}

func getHash(fullString []rune) string {
	fullStringS := string(fullString)
	fullStringB := []byte(fullStringS)
	hash := sha256.Sum256([]byte(fullStringB))
	return hex.EncodeToString(hash[:])

}
