package test

import (
	"testing"

	"../core"
	"github.com/stretchr/testify/assert"
)

func TestKMPTable(t *testing.T) {
	w := "abcdabd"
	expected := []int{-1, 0, 0, 0, -1, 0, 2, 0}
	res := core.GetTable(w)
	assert.Equal(t, expected, res)

	w = "participate in parachute"
	res = core.GetTable(w)
	expected = []int{-1, 0, 0, 0, 0, 0, 0, -1, 0, 2, 0, 0, 0, 0, 0, -1, 0, 0, 3, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, expected, res)
}

func TestKMPFindAll(t *testing.T) {
	// single occurence
	str := "abcdabd"
	pattern := "abc"
	res, err := core.FindAll(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{0}, res)

	// multiple occurence
	pattern = "ab"
	res, err = core.FindAll(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{0, 4}, res)

	// whole string
	pattern = "abcdabd"
	res, err = core.FindAll(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{0}, res)

	// no match
	pattern = "abcdabda"
	res, err = core.FindAll(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{}, res)

	// overlapping occurences
	str = "abababab"
	pattern = "abab"
	res, err = core.FindAll(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{0, 2, 4}, res)
}

func TestKMPFindFirst(t *testing.T) {
	// single occurence
	str := "abcdabd"
	pattern := "abc"
	res, err := core.FindFirst(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res)

	// multiple occurence
	pattern = "ab"
	res, err = core.FindFirst(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res)

	// whole string
	pattern = "abcdabd"
	res, err = core.FindFirst(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res)

	// no match
	pattern = "abcdabda"
	res, err = core.FindFirst(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, -1, res)

	// overlapping occurences
	str = "abababab"
	pattern = "abab"
	res, err = core.FindFirst(str, pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res)
}

func TestKMPReplace(t *testing.T) {
	// single occurence
	str := "abcdabd"
	pattern := "abc"
	newPattern := "aaa"
	expected := "aaadabd"
	res, err := core.Replace(str, pattern, newPattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)

	// multiple occurence
	pattern = "ab"
	expected = "aaacdaaad"
	res, err = core.Replace(str, pattern, newPattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)

	// whole string
	pattern = "abcdabd"
	expected = "aaa"
	res, err = core.Replace(str, pattern, newPattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)

	// no match
	pattern = "abcdabda"
	expected = "abcdabd"
	res, err = core.Replace(str, pattern, newPattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)

	// overlapping occurences
	str = "abababab"
	pattern = "abab"
	expected = "aaaaaa"
	res, err = core.Replace(str, pattern, newPattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)
}
