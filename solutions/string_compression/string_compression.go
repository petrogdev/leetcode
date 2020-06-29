package main

import (
	"fmt"
	"strconv"
)

func main() {
	originalStringExample := "aabcccccaaa"

	compressedString := compressString(originalStringExample)
	fmt.Println("compressed string", compressedString)
}

func compressString(originalString string) string {
	if len(originalString) == 0 {
		return ""
	}

	var (
		compressedString string
		prevElement      string
		count            int = 0
	)

	for i := 0; i < len(originalString); i++ {
		if string(originalString[i]) == prevElement {
			count++
		} else {
			if prevElement != "" {
				compressedString += strconv.Itoa(count)
			}
			count = 1
			compressedString += string(originalString[i])
		}

		prevElement = string(originalString[i])
	}

	compressedString += strconv.Itoa(count)
	fmt.Println("compressedString: ", compressedString)

	if len(compressedString) >= len(originalString) {
		return originalString
	}

	return compressedString
}

func isNumber(val string) bool {
	if _, err := strconv.Atoi(val); err == nil {
		return true
	} else {
		return false
	}
}

// decompressString
func decompressString(compressedString string) string {
	return ""
}
