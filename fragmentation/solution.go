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

type Fragment struct {
	Data string
	Hash string
}

func (f *Fragment) isValid() bool {
	return SimpleHash(f.Data) == f.Hash
}

// ReconstructData rebuilds the original data string from a map of fragments.
// The input map should have fragment indices as keys and fragment values as values.
// The function returns the reconstructed data as a string, or an error if reconstruction fails.
//
// Parameters:
//   - input: a map where keys are fragment indices and values are fragment data.
//
// Returns:
//   - The reconstructed data as a string.
//   - An error if the reconstruction is unsuccessful (e.g., missing fragments or invalid input).
func ReconstructData(input map[int]Fragment) (string, error) {
	var sb strings.Builder

	// we need the sorted keys to reconstruct the data in proper order
	sortedKeys := getSortedKeys(input)

	for _, key := range sortedKeys {
		fragment := input[key]
		if !fragment.isValid() {
			return "", ErrTamperedData
		}
		sb.WriteString(fragment.Data)
	}

	return sb.String(), nil
}

// SimpleHash computes and returns a simple hash value for the provided data string.
// The specific hash algorithm used is implementation-defined and intended for non-cryptographic purposes.
//
// Parameters:
//   - data: the input string to hash.
//
// Returns:
//   - A string representing the hash value of the input data.
func SimpleHash(data string) string {
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

func getSortedKeys(input map[int]Fragment) []int {
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
