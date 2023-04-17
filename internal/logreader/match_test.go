package logreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddKillWorld(t *testing.T) {
	m := NewMatch()

	m.AddKill("<world>", "Killed", "By")

	assert.Equal(t, 1, len(m.Kills))
	assert.Equal(t, 1, m.TotalKills)
	assert.Equal(t, struct{}{}, m.playersMap["<world>"])
	assert.Equal(t, 1, len(m.Players))
}
