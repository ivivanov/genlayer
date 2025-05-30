package fragmentation

import (
	"math"
	"strings"
	"testing"

	th "developers-challenge/pkg/testhelpers"
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
			actHash := SimpleHash(tc.input)
			th.AssertEqualInts(t, len(actHash), tc.expLen)
			th.AssertEqualStrings(t, actHash, tc.expOut)
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
				SimpleHash(tc.input)
			}
		})
	}
}

func TestReconstructData(t *testing.T) {
	testCases := []struct {
		desc       string
		fragments  map[int]Fragment
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
			fragments: func() map[int]Fragment {
				fragments := initTestInput()
				fragments[0] = Fragment{"tampered", "000011001100000001000001111000"}
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
			data, err := ReconstructData(tc.fragments)
			if tc.shouldFail {
				th.AssertNotNilError(t, err)
				th.AssertCorrectError(t, err, tc.expErr)
			} else {
				th.AssertNilError(t, err)
			}

			if data != tc.expOut {
				th.AssertEqualStrings(t, data, tc.expOut)
			}
		})
	}
}

func initTestInput() map[int]Fragment {
	fragments := make(map[int]Fragment)
	fragments[3] = Fragment{Data: "!", Hash: SimpleHash("!")}
	fragments[2] = Fragment{Data: "World", Hash: SimpleHash("World")}
	fragments[1] = Fragment{Data: "Hello", Hash: SimpleHash("Hello")}
	return fragments
}
