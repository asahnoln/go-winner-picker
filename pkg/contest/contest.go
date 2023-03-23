package contest

import (
	"errors"
	"math/rand"
)

var ErrNoContestants = errors.New("contest: got no contestants to pick from")

type Contestant struct {
	Name string
}

type Contest interface {
	Contestants() ([]Contestant, error)
}

func PickFrom(c Contest, r *rand.Rand) (Contestant, error) {
	cs, err := c.Contestants()
	if err != nil {
		return Contestant{}, err
	}

	if len(cs) == 0 {
		return Contestant{}, ErrNoContestants
	}

	return cs[r.Intn(len(cs))], nil
}
