package nback

import (
	"math/rand"
	"time"

	"github.com/srbiotik/nback/models"
)

func Generate(declaration models.ModelDeclaration) []int {
	amount_of_stims, ammount_of_hits := amount_of_stims_and_hits(declaration)

	hits := make([]int, ammount_of_hits)
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < ammount_of_hits; i++ {
		hits = append(hits, rand.Intn(amount_of_stims-declaration.Level))
	}

	collection := make([]int, amount_of_stims)
	for i := 0; i < amount_of_stims; i++ {
		collection = append(collection, rand.Intn(declaration.StimCollectionLength))
	}

	for i := 0; i < len(hits); i++ {
		collection[hits[i]+declaration.Level] = collection[hits[i]]
	}

	return collection
}

func amount_of_stims_and_hits(d models.ModelDeclaration) (int, int) {
	amount_of_stims := d.Level + d.Level*d.Rounds
	ammount_of_hits := int(float64(amount_of_stims)*d.Reccurence) - int(float64(amount_of_stims)*d.Reccurence)%2
	return amount_of_stims, ammount_of_hits
}
