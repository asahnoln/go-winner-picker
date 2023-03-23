package types

import "github.com/asahnoln/go-winner-picker/pkg/contest"

type Instagram struct {
}

func NewInstagram() *Instagram {
	return nil
}

func (i *Instagram) Contestants() ([]contest.Contestant, error) {
	return []contest.Contestant{}, nil
}
