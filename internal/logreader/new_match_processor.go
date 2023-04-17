package logreader

import (
	"errors"
	"strings"
)

type (
	NewMatchProcessor struct {
	}
)

// returns a new match event. We will consider that there is no error
// since we only need to check the selector key to verify if we
// have a new match
//
// 20:37 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv...
func (p NewMatchProcessor) Process(line string) (any, error) {
	// just to ensure that we are processing the right line
	idx := strings.Index(line, "InitGame")
	if idx == -1 {
		return nil, errors.New("InitGame not found on the given line")
	}
	return NewMatchEvent{}, nil
}

func (p NewMatchProcessor) SelectorKey() string {
	return "InitGame:"
}
