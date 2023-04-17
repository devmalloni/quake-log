package logreader

import "fmt"

const (
	WorldPlayer = "<world>"
)

type (
	Match struct {
		TotalKills  int            `json:"total_kills"`
		Players     []string       `json:"players"`
		Kills       map[string]int `json:"kills"`
		KillByMeans map[string]int `json:"kill_by_means"`

		playersMap map[string]struct{}
	}

	MatchReport map[string]*Match
)

func NewMatch() *Match {
	return &Match{
		Kills:       make(map[string]int),
		Players:     make([]string, 0),
		KillByMeans: make(map[string]int),
		playersMap:  make(map[string]struct{}),
	}
}

func (p *Match) AddKill(killer, killed, by string) {
	p.TotalKills++
	p.KillByMeans[by]++
	if killer == WorldPlayer {
		// Should we handle negative score cases?
		p.Kills[killed]--
	} else {
		p.Kills[killer]++
	}

	p.AddPlayer(killer, killed)
}

func (p *Match) AddPlayer(players ...string) {
	for _, player := range players {
		if player == WorldPlayer {
			continue
		}
		// we use p.playersMap as support structure to
		// make search for players O(1) for insertion.
		//
		// It is here just to show some possible optimization, because
		// adds code complexity if we decided to re-read this report.
		// For small files with problably not too many players,
		// a linear search on p.Players should not be a problem at all.
		if _, ok := p.playersMap[player]; ok {
			continue
		}
		p.playersMap[player] = struct{}{}
		p.Players = append(p.Players, player)
	}
}

func (p MatchReport) AddMatch() *Match {
	currentIndex := len(p) + 1
	id := fmt.Sprintf("game_%d", currentIndex)
	m := NewMatch()

	p[id] = m

	return m
}
