package fragmentation

import (
	"math"
	"strings"
	"testing"
)

func TestSimpleHash(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string // string to be hashed
		expOut string // expected hash
		expLen int    // expected hash length
	}{
		{
			desc:   "Successful_hashing",
			input:  "Hello",
			expOut: "000011001100000001000001111000",
			expLen: HashLen,
		},
		{
			desc:   "EmptyString_ShouldSucceed",
			input:  "",
			expOut: "000000000000000000000000000000",
			expLen: HashLen,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actHash := simpleHash(tc.input)
			assertEqualInts(t, len(actHash), tc.expLen)
			assertEqualStrings(t, actHash, tc.expOut)
		})
	}
}

// We can use the result to benchmark against other hashing algorithms
func BenchmarkSimpleHash(b *testing.B) {
	testCases := []struct {
		desc  string
		input string
	}{
		{
			desc:  "Regular_Input",
			input: "Benchmark",
		},
		{
			desc:  "MidSize_Input",
			input: strings.Repeat("xyz", math.MaxInt16),
		},
		{
			desc:  "LargeSize_Input",
			input: strings.Repeat("xyz", math.MaxInt32),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.desc, func(b *testing.B) {
			for b.Loop() {
				simpleHash(tc.input)
			}
		})
	}
}

func assertEqualStrings(t *testing.T, act, exp string) {
	t.Helper()
	if act != exp {
		t.Errorf("exp: %v, act: %v", exp, act)
	}
}

func assertEqualInts(t *testing.T, act, exp int) {
	t.Helper()
	if act != exp {
		t.Errorf("exp: %v, act: %v", exp, act)
	}
}
