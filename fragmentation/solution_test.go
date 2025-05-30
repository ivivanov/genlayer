package fragmentation

import (
	"errors"
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
			desc:   "Successful_Hashing",
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

func TestReconstructData(t *testing.T) {
	testCases := []struct {
		desc       string
		fragments  map[int]fragment
		expOut     string
		shouldFail bool
		expErr     error
	}{
		{
			desc:      "Successful_Reconstruction",
			fragments: initTestInput(),
			expOut:    "HelloWorld!",
		},
		{
			desc: "TamperedFragments_ShouldFailWith_ErrTamperedData",
			fragments: func() map[int]fragment {
				fragments := initTestInput()
				fragments[0] = fragment{"tampered", "000011001100000001000001111000"}
				return fragments
			}(),
			shouldFail: true,
			expErr:     ErrTamperedData,
		},
		{
			desc:      "NilFragments_ShouldReturn_EmptyString",
			fragments: nil,
			expOut:    "",
		},
	}
	for _, tc := range testCases {
		tc := tc

		t.Run(tc.desc, func(t *testing.T) {
			data, err := reconstructData(tc.fragments)
			if tc.shouldFail {
				assertNotNilError(t, err)
				assertCorrectError(t, err, tc.expErr)
			} else {
				assertNilError(t, err)
			}

			if data != tc.expOut {
				assertEqualStrings(t, data, tc.expOut)
			}
		})
	}
}

func initTestInput() map[int]fragment {
	fragments := make(map[int]fragment)
	fragments[3] = fragment{data: "!", hash: simpleHash("!")}
	fragments[2] = fragment{data: "World", hash: simpleHash("World")}
	fragments[1] = fragment{data: "Hello", hash: simpleHash("Hello")}
	return fragments
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

func assertNotNilError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("expected error to be not nil, but got nil")
	}
}

func assertNilError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected error to be nil, but got not nill")
	}
}

func assertCorrectError(t *testing.T, act, exp error) {
	t.Helper()
	if !errors.Is(act, exp) {
		t.Errorf("exp: %v, act: %v", exp, act)
	}
}
