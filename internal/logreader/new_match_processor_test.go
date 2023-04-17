package logreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMatchProcessorSuccess(t *testing.T) {
	processor := NewMatchProcessor{}

	rawEvt, err := processor.Process(" 20:37 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv... ")
	_, ok := rawEvt.(NewMatchEvent)

	assert.NoError(t, err, "unexpected error on processor")
	assert.True(t, ok, "expected result to be a NewMatchEvent type")
}

func TestNewMatchProcessorErr(t *testing.T) {
	processor := NewMatchProcessor{}

	_, err := processor.Process(" 20:37 Foo: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv... ")

	assert.Error(t, err, "unexpected error on processor")
}
