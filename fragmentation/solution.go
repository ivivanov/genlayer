package fragmentation

import (
	"fmt"
	"strings"
)

const (
	HashLen  = 30
	PrimeNum = 29
)

func simpleHash(data string) string {
	result := 0

	// ignore potential int overflow
	// combines all chars
	for _, v := range data {
		result = result*PrimeNum + int(v)
	}

	// get the binary representation
	binary := fmt.Sprintf("%b", result)
	if len(binary) < HashLen {
		binary = fmt.Sprint(strings.Repeat("0", HashLen-len(binary)), binary)
	}

	// trim excess output up to the required length
	if len(binary) > HashLen {
		binary = binary[:HashLen]
	}

	return binary
}
