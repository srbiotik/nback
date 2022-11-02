package nback

import (
	"github.com/srbiotik/nback/models"
)

func Generate(d models.ModelDeclaration) []int {
	amount_of_stims := d.Level + d.Level*d.Rounds
	ammount_of_hits := int(float64(amount_of_stims)*d.Reccurence) - int(float64(amount_of_stims)*d.Reccurence)%2

	stims := add_hits(amount_of_stims, ammount_of_hits, d.Level, d.StimCollectionLength)
	stims = add_misses(amount_of_stims, d.Level, d.StimCollectionLength, stims)

	return stims
}
