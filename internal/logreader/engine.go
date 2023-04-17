package logreader

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/rs/zerolog/log"
)

var (
	ErrNoToken = errors.New("no token found on the given line")
)

type (
	// Processor defines an interface to handle each log line
	LineProcessor interface {
		// must return the key of this processor.
		// e.g. Kill, InitGame, ClientConnect
		SelectorKey() string
		// Process the given line and return the event associated to it
		Process(line string) (any, error)
	}

	// Event returned by every log line that represents a kill
	// like:
	//
	// 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
	KillEvent struct {
		Killer string
		Killed string
		By     string
	}

	// Event returned by every log line that represents a new match
	// has started. Like:
	//
	// 20:37 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv...
	NewMatchEvent struct{}

	// Engine process the stream of events to generate
	// the match report.
	Engine struct {
		processors map[string]LineProcessor
	}
)

func NewEngine() *Engine {
	return &Engine{
		processors: make(map[string]LineProcessor),
	}
}

func NewDefaultEngine() *Engine {
	engine := NewEngine()
	engine.AddProcessor(KillProcessor{})
	engine.AddProcessor(NewMatchProcessor{})

	return engine
}

func (p *Engine) AddProcessor(lp LineProcessor) {
	if p.processors == nil {
		p.processors = make(map[string]LineProcessor)
	}
	p.processors[lp.SelectorKey()] = lp
}

func (p *Engine) Process(r *bufio.Reader) (MatchReport, error) {
	var currentMatch *Match
	report := make(MatchReport)
	for {
		// read current line from bufio.Reader.
		// Note that readLine doesn't returns the whole line
		// if it not fits in the buffer. We're assuming here that
		// every info we want fits on the standard buffer.
		currentLine, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		kind, err := p.getEventKind(string(currentLine))
		if err != nil {
			if errors.Is(err, ErrNoToken) {
				log.Warn().Err(err).Msg("line skipped due to event kind not found")
				continue
			} else {
				return nil, err
			}
		}

		processor, ok := p.processors[kind]
		if !ok {
			log.Warn().Str("processorKind", kind).Msg("line skipped due to processor not found")
			continue
		}

		evt, err := processor.Process(string(currentLine))
		if err != nil {
			return nil, err
		}

		switch parsedEvent := evt.(type) {
		case NewMatchEvent:
			currentMatch = report.AddMatch()
		case KillEvent:
			currentMatch.AddKill(parsedEvent.Killer, parsedEvent.Killed, parsedEvent.By)
		}
	}
	return report, nil
}

func (p *Engine) getEventKind(l string) (string, error) {
	l = strings.Trim(l, " ")
	tokens := strings.Split(l, " ")

	if len(tokens) > 2 {
		return tokens[1], nil
	}

	return "", fmt.Errorf("no token found on line %s: %w", l, ErrNoToken)
}
