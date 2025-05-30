package fragmentation

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

const (
	HashLen  = 30
	PrimeNum = 29
)

var (
	ErrTamperedData = errors.New("data integrity verification failed")
)

type fragment struct {
	data string
	hash string
}

func (f *fragment) isValid() bool {
	return simpleHash(f.data) == f.hash
}

func reconstructData(input map[int]fragment) (string, error) {
	var sb strings.Builder

	// we need the sorted keys to reconstruct the data in proper order
	sortedKeys := getSortedKeys(input)

	for _, key := range sortedKeys {
		fragment := input[key]
		if !fragment.isValid() {
			return "", ErrTamperedData
		}
		sb.WriteString(fragment.data)
	}

	return sb.String(), nil
}

func getSortedKeys(input map[int]fragment) []int {
	keys := make([]int, len(input))

	// extract keys from map
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}

	slices.Sort(keys)

	return keys
}

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
