package semver

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_isOlder(t *testing.T) {
	testCases := map[string]struct {
		semver1  string
		semver2  string
		expected bool
	}{
		"SameMayorAndMinorAndPatch": {
			semver1:  "1.2.4",
			semver2:  "1.2.4",
			expected: false,
		},
		"SameMayorAndMinorButOlderPatch": {
			semver1:  "1.2.3",
			semver2:  "1.2.4",
			expected: true,
		},
		"SameMayorAndMinorButNewerPatch": {
			semver1:  "1.2.5",
			semver2:  "1.2.4",
			expected: false,
		},
		"SameMayorAndMinorButPatchMissingInFirst": {
			semver1:  "1.2",
			semver2:  "1.2.4",
			expected: true,
		},
		"SameMayorAndMinorButPatchMissingInSecond": {
			semver1:  "1.2.4",
			semver2:  "1.2",
			expected: false,
		},
		"SameMayorAndMinor": {
			semver1:  "1.2",
			semver2:  "1.2",
			expected: false,
		},
		"SameMayorButOlderMinor": {
			semver1:  "1.1.7",
			semver2:  "1.2.4",
			expected: true,
		},
		"SameMayorButNewerMinor": {
			semver1:  "1.3.2",
			semver2:  "1.2.4",
			expected: false,
		},
		"SameMayorButMinorMissingInFirst": {
			semver1:  "1",
			semver2:  "1.2",
			expected: true,
		},
		"SameMayorButMinorMissingInSecond": {
			semver1:  "1.2",
			semver2:  "1",
			expected: false,
		},
		"SameMayor": {
			semver1:  "1",
			semver2:  "1",
			expected: false,
		},
		"OlderMayor": {
			semver1:  "1.4.7",
			semver2:  "2.2.4",
			expected: true,
		},
		"NewerMayor": {
			semver1:  "2.1.2",
			semver2:  "1.2.4",
			expected: false,
		},
		"ComparesMayorAsNumber": {
			semver1:  "2",
			semver2:  "10",
			expected: true,
		},
		"ComparesMinorAsNumber": {
			semver1:  "1.2",
			semver2:  "1.10",
			expected: true,
		},
		"ComparesPatchAsNumber": {
			semver1:  "1.1.2",
			semver2:  "1.1.10",
			expected: true,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			actual := isOlder(testCase.semver1, testCase.semver2)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func Test_IsValid(t *testing.T) {
	testCases := map[string]struct {
		semver   string
		expected bool
	}{
		"Mayor": {
			semver:   "1",
			expected: true,
		},
		"MayorWithMultipleDigits": {
			semver:   "21",
			expected: true,
		},
		"MayorEndsWithDot": {
			semver:   "1.",
			expected: false,
		},
		"MayorWithInvalidChar": {
			semver:   "2a",
			expected: false,
		},
		"MayorAndMinor": {
			semver:   "1.3",
			expected: true,
		},
		"MayorAndMinorWithMultipleDigits": {
			semver:   "1.34",
			expected: true,
		},
		"MayorAndMinorEndsWithDot": {
			semver:   "1.2.",
			expected: false,
		},
		"MayorAndMinorWithInvalidChar": {
			semver:   "2.2a",
			expected: false,
		},
		"MayorAndMinorAndPatch": {
			semver:   "1.1.3",
			expected: true,
		},
		"MayorAndMinorAndPatchWithMultipleDigits": {
			semver:   "1.1.34",
			expected: true,
		},
		"MayorAndMinorAmdPatchEndsWithDot": {
			semver:   "1.2.4.",
			expected: false,
		},
		"MayorAndMinorAndPatchWithInvalidChar": {
			semver:   "2.1.2a",
			expected: false,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			actual := IsValid(testCase.semver)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func Test_SliceStableComparatorFor_ValidVersions(t *testing.T) {
	semvers := []string{"2", "1.20.1", "1.3", "1.20.4", "1.19.1"}
	expected := []string{"1.3", "1.19.1", "1.20.1", "1.20.4", "2"}

	comparator, err := SliceStableComparatorFor(semvers)
	require.NoError(t, err)

	sort.SliceStable(semvers, comparator)

	assert.Equal(t, expected, semvers)
}

func Test_SliceStableComparatorFor_InvalidVersions(t *testing.T) {
	semvers := []string{"2", "1.20.a"}

	_, err := SliceStableComparatorFor(semvers)
	assert.EqualError(t, err, "invalid semantic version: 1.20.a")
}