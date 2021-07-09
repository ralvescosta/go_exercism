package tournament

import (
	"errors"
	"io"
	"strings"
)

func Tally(reader io.Reader, writer io.Writer) error {
	scoreboard := ScoreBoard{}
	splitReader, err := readTournament(reader)
	if err != nil {
		return err
	}

	for _, readerSlice := range splitReader {
		match, err := getMatch(readerSlice)
		if err != nil {
			continue
		}
		scoreboard.UpdateBoard(*match)
	}

	writer.Write([]byte(scoreboard.SprintF()))
	return nil
}

func readTournament(reader io.Reader) ([]string, error) {
	b := make([]byte, 1024)
	if _, err := reader.Read(b); err != nil {
		return []string{""}, err
	}
	readerToString := string(b)
	return strings.Split(readerToString, "\n"), nil
}

func getMatch(readerSlice string) (*Match, error) {
	if readerSlice == "" {
		return nil, errors.New("")
	}
	splitSlice := strings.Split(readerSlice, ";")
	if len(splitSlice) < 2 {
		return nil, errors.New("")
	}

	matchResult := splitSlice[2]
	m := Match{
		Home:     strings.TrimSpace(splitSlice[0]),
		Visiting: strings.TrimSpace(splitSlice[1]),
		Winier:   "",
		Loser:    "",
	}

	switch matchResult {
	case "win":
		m.Winier = m.Home
		m.Loser = m.Visiting
	case "loss":
		m.Winier = m.Visiting
		m.Loser = m.Home
	}

	return &m, nil
}

type Match struct {
	Home     string
	Visiting string
	Winier   string
	Loser    string
}

type Score struct {
	MP int
	W  int
	D  int
	L  int
	P  int
}

type ScoreBoard struct {
	Board map[string]Score
}

func (sb *ScoreBoard) initializeScoreboard(m Match) {
	if m.Winier == "" {
		sb.Board = map[string]Score{
			m.Home: {
				MP: 1,
				W:  0,
				D:  1,
				L:  0,
				P:  1,
			},
			m.Visiting: {
				MP: 1,
				W:  0,
				D:  1,
				L:  0,
				P:  1,
			},
		}
	} else {
		sb.Board = map[string]Score{
			m.Winier: {
				MP: 1,
				W:  1,
				D:  0,
				L:  0,
				P:  3,
			},
			m.Loser: {
				MP: 1,
				W:  0,
				D:  0,
				L:  1,
				P:  0,
			},
		}
	}
}

func (sb *ScoreBoard) UpdateBoard(match Match) {
	if sb.Board == nil {
		sb.initializeScoreboard(match)
		return
	}
}

func (ScoreBoard) SprintF() string {
	return ""
}
