package nback

import (
	"fmt"
	"math/rand"
	"time"
)

func add_hits(amount_of_stims int, ammount_of_hits int, level int, collection_length int) []int {
	hit_count := 0
	t := time.Now().UnixNano()
	fmt.Println("Time:", t)
	rand.Seed(t)

	stims := make([]int, amount_of_stims)
	for i := 0; i < amount_of_stims; i++ {
		stims[i] = -1
	}

	for {
		index := randNum(amount_of_stims - level)
		position := randNum(collection_length)
		if stims[index] != -1 || (index <= amount_of_stims-level && stims[index+level] != -1) || (index >= amount_of_stims+level && stims[index-level] != -1) {
			continue
		}
		stims[index] = position
		stims[index+level] = position
		hit_count++
		if ammount_of_hits == hit_count {
			break
		}
	}
	return stims
}
