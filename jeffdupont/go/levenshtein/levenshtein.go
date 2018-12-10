package levenshtein

import "math"

// Distance find the levenshtein value between two strings
func Distance(s, t string) int {
	if len(s) == 0 {
		return len(t)
	}
	if len(t) == 0 {
		return len(s)
	}

	m := generate(s, t)
	for x := 1; x <= len(s); x++ {
		for y := 1; y <= len(t); y++ {
			cost := 0
			if s[x-1] != t[y-1] {
				cost = 1
			}
			m[xy{x, y}] = int(math.Min(float64(m[xy{x - 1, y}]+1), math.Min(float64(m[xy{x, y - 1}]+1), float64(m[xy{x - 1, y - 1}]+cost))))
		}
	}

	return m[xy{len(s), len(t)}]
}

type xy struct{ x, y int }
type matrix map[xy]int

func generate(s, t string) matrix {
	m := matrix{}
	for x := 0; x <= len(s); x++ {
		for y := 0; y <= len(t); y++ {
			m[xy{x, y}] = 0
			if x == 0 {
				m[xy{x, y}] = y
			}
			if y == 0 {
				m[xy{x, y}] = x
			}
		}
	}
	return m
}
