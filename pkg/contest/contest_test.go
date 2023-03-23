package contest_test

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/asahnoln/go-winner-picker/pkg/contest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomPicker(t *testing.T) {
	c := &contestStub{
		cs: []contest.Contestant{
			{"Not Good"},
			{"Iam Random"},
			{"Not Good"},
		},
	}

	r := rand.New(rand.NewSource(2))
	w, err := contest.PickFrom(c, r)
	require.NoError(t, err)

	assert.Equal(t, "Iam Random", w.Name)
}

func TestPickerError(t *testing.T) {
	c := &contestStub{
		err: errors.New("source error"),
	}

	r := rand.New(rand.NewSource(1))
	_, err := contest.PickFrom(c, r)
	require.Error(t, err)
}

func TestNoContestants(t *testing.T) {
	c := &contestStub{}

	r := rand.New(rand.NewSource(1))
	_, err := contest.PickFrom(c, r)
	require.ErrorIs(t, err, contest.ErrNoContestants)
}

type contestStub struct {
	cs  []contest.Contestant
	err error
}

func (c *contestStub) Contestants() ([]contest.Contestant, error) {
	if c.err != nil {
		return []contest.Contestant{}, c.err
	}

	return c.cs, nil
}
