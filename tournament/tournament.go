package tournament

import (
	"io"
	"strings"
)

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

func Tally(reader io.Reader, writer io.Writer) error {
	scoreboard := ScoreBoard{}

	b := make([]byte, 1024)
	if _, err := reader.Read(b); err != nil {
		return err
	}
	readerToString := string(b)
	splitReader := strings.Split(readerToString, "\n")

	for _, readerSlice := range splitReader {
		if readerSlice == "" {
			continue
		}
		splitSlice := strings.Split(readerSlice, ";")
		if len(splitSlice) < 2 {
			continue
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

		scoreboard.UpdateBoard(m)
	}

	writer.Write([]byte(scoreboard.SprintF()))
	return nil
}
