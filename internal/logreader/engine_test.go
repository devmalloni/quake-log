package logreader

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEventKindSuccess(t *testing.T) {
	engine := Engine{}

	str, err := engine.getEventKind(" 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT  ")

	assert.NoError(t, err)
	assert.Equal(t, "Kill:", str)
}

func TestEngineSuccess(t *testing.T) {
	killProcessor := KillProcessor{}
	initProcessor := NewMatchProcessor{}
	engine := Engine{}
	// we could mock processors here to test
	// only engine. I'll not do that just to
	// save my time, but in more complex structures
	// with more dependencies, mock things like database
	// calls is a better way to ensure that what you are testing is correct.
	engine.AddProcessor(killProcessor)
	engine.AddProcessor(initProcessor)

	input := bufio.NewReader(bytes.NewReader([]byte(`
	0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0
	20:37 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\bot_minplayers\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0
	20:54 Kill: 1022 2 22: <world> killed Adam Sandler by MOD_TRIGGER_HURT
	21:07 Kill: 1022 2 22: Keanu Reeves killed Adam Sandler by MOD_TRIGGER_HURT
	21:42 Kill: 1022 2 22: Adam Sandler killed Keanu Reeves by MOD_TRIGGER_HURT
	0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0
	1:26 Kill: 1022 4 22: Emma Watson killed Robert Deniro by MOD_TRIGGER_HURT
	1:32 Kill: 1022 4 22: <world> killed Robert Deniro by MOD_SLIME
	`)))

	report, err := engine.Process(input)

	assert.NoError(t, err)
	assert.NotNil(t, report)
	assert.Equal(t, len(report), 3, "wrong match count")
	assert.Equal(t, len(report["game_2"].Kills), 2)
	assert.Equal(t, len(report["game_2"].KillByMeans), 1)
	assert.Equal(t, len(report["game_3"].Kills), 2)
	assert.Equal(t, len(report["game_3"].KillByMeans), 2)
}
