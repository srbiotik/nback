package nback

import (
	"math/rand"
	"time"
)

func add_misses(amount_of_stims int, level int, collection_length int, stims []int) []int {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < len(stims); i++ {
		if stims[i] == -1 {
			for {
				stim := randNum(collection_length)
				if (i+level < amount_of_stims && stims[i+level] == stim) {
					continue
				}
				if (i-level >= 0 && stims[i-level] == stim) {
					continue
				}
				stims[i] = stim
				break
			}
		}
	}
	return stims
}

