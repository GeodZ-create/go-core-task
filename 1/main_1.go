package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64
	getType(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fullString := valueToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println(fullString)
	fullStringRune := stringToRune(fullString)
	fmt.Println(string(putSalt(fullStringRune)))
	fullStringWithSalt := putSalt(fullStringRune)
	fmt.Println(fullStringWithSalt)
	fullStringHash := getHash(fullStringWithSalt)
	fmt.Println(fullStringHash)
}

func getType(value ...any) []string {
	var result []string
	for _, v := range value {
		result = append(result, fmt.Sprintf("%T", v))
	}
	for _, t := range result {
		fmt.Println(t)
	}
	return result
}

func valueToString(value ...any) string {
	var fullString string
	for _, v := range value {
		fullString += fmt.Sprint(v)
	}
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
