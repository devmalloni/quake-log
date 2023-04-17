package logreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKillProcessorSuccess(t *testing.T) {
	processor := KillProcessor{}

	rawEvt, err := processor.Process(" 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT")
	evt, ok := rawEvt.(KillEvent)

	assert.NoError(t, err, "unexpected error on processor")
	assert.True(t, ok, "expected result to be a KillEvent type")
	assert.Equal(t, "MOD_TRIGGER_HURT", evt.By, "unexpected by")
	assert.Equal(t, "Isgalamido", evt.Killed, "unexpected killed")
	assert.Equal(t, "<world>", evt.Killer, "unexpected killer")
}
