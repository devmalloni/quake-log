package logreader

import "strings"

type (
	KillProcessor struct {
	}
)

// 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
func (p KillProcessor) Process(line string) (any, error) {
	var event KillEvent

	line = strings.Trim(line, " ")
	line = strings.Split(line, ":")[3]
	lns := strings.Split(line, "by")
	event.By = strings.Trim(lns[1], " ")
	lns = strings.Split(lns[0], "killed")
	event.Killed = strings.Trim(lns[1], " ")
	event.Killer = strings.Trim(lns[0], " ")

	return event, nil
}

func (p KillProcessor) SelectorKey() string {
	return "Kill:"
}
