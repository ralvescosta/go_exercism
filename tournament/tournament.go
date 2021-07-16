package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

func Tally(reader io.Reader, writer io.Writer) error {
	rows := read(reader)
	hashMap := make(map[string]*TeamScore)
	var board Scoreboard

	for i, row := range rows {
		if len(strings.Trim(row, " ")) == 0 || row[0] == '#' {
			continue
		}
		match, err := parseMatch(row)
		if err != nil {
			return err
		}
		if _, ok := hashMap[match.HomeTeam]; !ok {
			hashMap[match.HomeTeam] = &TeamScore{match.HomeTeam, 0, 0, 0, 0, 0}
			board = append(board, hashMap[match.HomeTeam])
		}
		if _, ok := hashMap[match.VisitingTeam]; !ok {
			hashMap[match.VisitingTeam] = &TeamScore{match.VisitingTeam, 0, 0, 0, 0, 0}
			board = append(board, hashMap[match.VisitingTeam])
		}
		team1, team2 := hashMap[match.HomeTeam], hashMap[match.VisitingTeam]
		team1.MatchesPlayed += 1
		team2.MatchesPlayed += 1
		switch match.Result {
		case "win":
			team1.Wins += 1
			team1.Points += 3
			team2.Losses += 1
		case "loss":
			team1.Losses += 1
			team2.Wins += 1
			team2.Points += 3
		case "draw":
			team1.Draws += 1
			team2.Draws += 1
			team1.Points += 1
			team2.Points += 1
		default:
			return fmt.Errorf("invalid match result \"%s\" in line %d", match.Result, i)
		}
	}

	sort.Sort(board)
	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, team := range board {
		fmt.Fprintf(writer, "%-31s| %2d | %2d | %2d | %2d | %2d\n",
			team.Name,
			team.MatchesPlayed,
			team.Wins,
			team.Draws,
			team.Losses,
			team.Points)
	}

	return nil
}

type TeamScore struct {
	Name          string
	MatchesPlayed int
	Wins          int
	Losses        int
	Draws         int
	Points        int
}

type Match struct {
	HomeTeam     string
	VisitingTeam string
	Result       string
}

type Scoreboard []*TeamScore

func (a Scoreboard) Len() int      { return len(a) }
func (a Scoreboard) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Scoreboard) Less(i, j int) bool {
	if a[i].Points < a[j].Points {
		return false
	} else if a[i].Points > a[j].Points {
		return true
	} else {
		return a[i].Name < a[j].Name
	}
}

func read(reader io.Reader) []string {
	var buffer bytes.Buffer
	buffer.ReadFrom(reader)
	input := buffer.String()
	input = strings.Trim(input, "\n")
	return strings.Split(input, "\n")
}

func parseMatch(match string) (Match, error) {
	record := strings.Split(match, ";")
	if len(record) != 3 || record[0] == record[1] {
		return Match{}, errors.New("wrong number of parts on line")
	}

	return Match{
		HomeTeam:     record[0],
		VisitingTeam: record[1],
		Result:       record[2],
	}, nil
}
