package nback

import (
	"math/rand"
)

func is_hit(stims []int, index int, level int, stim int) bool {
	return stims[index] == stim && stims[index+level] == stim
}

func randNum(limit int) int {
	return rand.Intn(limit)
}
