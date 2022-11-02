package nback

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddHits(t *testing.T) {
	amount_of_stims := 33
	amount_of_hits := 8
	level := 3
	result := add_hits(amount_of_stims, amount_of_hits, level, 9)
	fmt.Println("Result", result)
	if len(result) != amount_of_stims {
		t.Fatalf(`Amount of stims not expected, want: %d, got: %d`, amount_of_stims, len(result))
	}
	hit_count := 0
	hits := map[string]int{}
	for i := 0; i < len(result); i++ {
		if _, ok := hits[strconv.Itoa(i)]; ok {
			continue
		}
		if result[i] != -1 {
			hits[strconv.Itoa(i+level)] = i
			hit_count++
		}
	}

	if hit_count != amount_of_hits {
		t.Fatalf(`Amount of hits not expected, want: %d, got: %d`, amount_of_hits, hit_count)
	}
	result_2 := add_hits(amount_of_stims, amount_of_hits, level, 9)
	fmt.Println("Result 2", result_2)
	for i := 0; i < len(result); i++ {
		same := true
		if result[i] != result_2[i] {
			same = false
		}
		if i == (len(result)-1) && same {
			t.Fatalf(`Results not random`)
		}
	}
}

func TestAddMisses(t *testing.T) {
	amount_of_stims := 33
	amount_of_hits := 8
	level := 3
	collection_length := 9
	stims := add_hits(amount_of_stims, amount_of_hits, level, collection_length)
	stims = add_misses(amount_of_stims, level, collection_length, stims)
	fmt.Println("Result", stims)
	miss_count := 0
	for i := 0; i < amount_of_stims; i++ {
		if i+level < amount_of_stims && stims[i+level] == stims[i] {
			continue
		}
		if i-level >= 0 && stims[i-level] == stims[i] {
			continue
		}
		miss_count++
	}
	if amount_of_stims-amount_of_hits*2 != miss_count {
		t.Fatalf(`Amount of misses not expected, want: %d, got: %d`, amount_of_stims-amount_of_hits, miss_count)
	}
}
