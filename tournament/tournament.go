package tournament

import (
	"errors"
	"io"
	"sort"
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
	Team string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

type ScoreBoard struct {
	Board []Score
}

func (sb *ScoreBoard) initializeScoreboard(m Match) {
	if m.Winier == "" {
		sb.Board = []Score{
			{
				Team: m.Home,
				MP:   1,
				W:    0,
				D:    1,
				L:    0,
				P:    1,
			},
			{
				Team: m.Visiting,
				MP:   1,
				W:    0,
				D:    1,
				L:    0,
				P:    1,
			},
		}
	} else {
		sb.Board = []Score{
			{
				Team: m.Home,
				MP:   1,
				W:    1,
				D:    0,
				L:    0,
				P:    3,
			},
			{
				Team: m.Visiting,
				MP:   1,
				W:    0,
				D:    0,
				L:    1,
				P:    0,
			},
		}
	}
}

func (sb *ScoreBoard) UpdateBoard(match Match) {
	if sb.Board == nil {
		sb.initializeScoreboard(match)
		return
	}

	if match.Winier != "" {
		indexWinier := sort.Search(len(sb.Board), func(i int) bool {
			return sb.Board[i].Team == match.Winier
		})
		indexLoser := sort.Search(len(sb.Board), func(i int) bool {
			return sb.Board[i].Team == match.Loser
		})

		if indexWinier < len(sb.Board) {
			sb.Board[indexWinier].W += 1
			sb.Board[indexWinier].P += 3
			sb.Board[indexWinier].MP += 1

			if indexWinier > 1 && sb.Board[indexWinier].P >= sb.Board[indexWinier-1].P {
				old := sb.Board[indexWinier-1]
				sb.Board[indexWinier-1] = sb.Board[indexWinier]
				sb.Board[indexWinier] = old
			}
		} else {
			sb.Board = append(sb.Board, Score{
				Team: match.Winier,
				MP:   1,
				W:    1,
				D:    0,
				L:    0,
				P:    3,
			})
		}

		if indexLoser < len(sb.Board) {
			sb.Board[indexLoser].L += 1
			sb.Board[indexLoser].MP += 1
		} else {
			sb.Board = append(sb.Board, Score{
				Team: match.Loser,
				MP:   1,
				W:    0,
				D:    0,
				L:    1,
				P:    0,
			})
		}
	} else {
		homeIndex := sort.Search(len(sb.Board), func(i int) bool {
			return sb.Board[i].Team == match.Home
		})
		visitingIndex := sort.Search(len(sb.Board), func(i int) bool {
			return sb.Board[i].Team == match.Visiting
		})

		if homeIndex < len(sb.Board) {
			sb.Board[homeIndex].W += 1
			sb.Board[homeIndex].P += 3
			sb.Board[homeIndex].MP += 1

			if homeIndex > 1 && sb.Board[homeIndex].P >= sb.Board[homeIndex-1].P {
				old := sb.Board[homeIndex-1]
				sb.Board[homeIndex-1] = sb.Board[homeIndex]
				sb.Board[homeIndex] = old
			}
		} else {
			sb.Board = append(sb.Board, Score{
				Team: match.Home,
				MP:   1,
				W:    0,
				D:    1,
				L:    0,
				P:    1,
			})
		}

		if visitingIndex < len(sb.Board) {
			sb.Board[visitingIndex].W += 1
			sb.Board[visitingIndex].P += 3
			sb.Board[visitingIndex].MP += 1

			if visitingIndex > 1 && sb.Board[visitingIndex].P >= sb.Board[visitingIndex-1].P {
				old := sb.Board[visitingIndex-1]
				sb.Board[visitingIndex-1] = sb.Board[visitingIndex]
				sb.Board[visitingIndex] = old
			}
		} else {
			sb.Board = append(sb.Board, Score{
				Team: match.Visiting,
				MP:   1,
				W:    0,
				D:    1,
				L:    0,
				P:    1,
			})
		}
	}
}

func (ScoreBoard) SprintF() string {
	return ""
}
