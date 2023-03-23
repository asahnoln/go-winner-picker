package types_test

import (
	"testing"

	"github.com/asahnoln/go-winner-picker/pkg/contest"
	"github.com/asahnoln/go-winner-picker/pkg/contest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstagramContest(t *testing.T) {
	i := types.NewInstagram()
	require.Implements(t, (*contest.Contest)(nil), i)

	cs, err := i.Contestants()
	require.NoError(t, err)
	assert.Len(t, cs, 4)
}
