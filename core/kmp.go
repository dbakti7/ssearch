package core

import (
	"bytes"
)

/*
More details at: https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
*/

func GetTable(w string) []int {
	strLen := len(w)
	res := make([]int, strLen+1)
	pos := 1
	ind := 0
	res[0] = -1

	for {
		if pos >= strLen {
			break
		}
		if w[pos] == w[ind] {
			res[pos] = res[ind]
			pos++
			ind++
		} else {
			res[pos] = ind
			ind = res[ind]
			for {
				if ind < 0 || w[pos] == w[ind] {
					break
				}
				ind = res[ind]
			}
			pos++
			ind++
		}
	}
	res[pos] = ind
	return res
}

func findFirst(s, pattern string, table []int) (int, error) {
	sPos := 0
	pos := 0
	sLen := len(s)
	patternLen := len(pattern)

	res := -1
	for {
		if sPos >= sLen {
			break
		}
		if pattern[pos] == s[sPos] {
			pos++
			sPos++
			if pos == patternLen {
				res = sPos - pos
				return res, nil
			}
		} else {
			pos = table[pos]
			if pos < 0 {
				sPos++
				pos++
			}
		}
	}

	return res, nil
}

func findAll(s, pattern string, table []int) ([]int, error) {
	sPos := 0
	pos := 0
	sLen := len(s)
	patternLen := len(pattern)

	res := make([]int, 0)
	for {
		if sPos >= sLen {
			break
		}
		if pattern[pos] == s[sPos] {
			pos++
			sPos++
			if pos == patternLen {
				res = append(res, sPos-pos)
				pos = table[pos] // table[patternLen] can't be -1
			}
		} else {
			pos = table[pos]
			if pos < 0 {
				sPos++
				pos++
			}
		}
	}

	return res, nil
}

func replace(s, oldPattern, newPattern string, table []int) (string, error) {
	sPos := 0
	pos := 0
	sLen := len(s)
	patternLen := len(oldPattern)
	var buffer bytes.Buffer

	res := make([]int, 0)
	for {
		if sPos >= sLen {
			break
		}
		if oldPattern[pos] == s[sPos] {
			pos++
			sPos++
			if pos == patternLen {
				res = append(res, sPos-pos)
				pos = table[pos] // table[patternLen] can't be -1
			}
		} else {
			pos = table[pos]
			if pos < 0 {
				sPos++
				pos++
			}
		}
	}

	prev := 0
	for _, i := range res {
		if i < prev {
			continue
		}
		for ind := prev; ind < i; ind++ {
			buffer.WriteByte(s[ind])
		}
		buffer.WriteString(newPattern)
		prev = i + patternLen
	}

	for ind := prev; ind < sLen; ind++ {
		buffer.WriteByte(s[ind])
	}
	return buffer.String(), nil
}

func FindFirst(s, pattern string) (int, error) {
	table := GetTable(pattern)
	return findFirst(s, pattern, table)
}

func FindAll(s, pattern string) ([]int, error) {
	table := GetTable(pattern)
	return findAll(s, pattern, table)
}

func Replace(s, oldPattern, newPattern string) (string, error) {
	table := GetTable(oldPattern)
	return replace(s, oldPattern, newPattern, table)
}
