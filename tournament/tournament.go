package tournament

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamScore struct {
	Name          string
	MatchesPlayed int
	Wins          int
	Losses        int
	Draws         int
	Points        int
}

type SortTeam []*TeamScore

func (a SortTeam) Len() int      { return len(a) }
func (a SortTeam) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortTeam) Less(i, j int) bool {
	if a[i].Points < a[j].Points {
		return false
	} else if a[i].Points > a[j].Points {
		return true
	} else {
		return a[i].Name < a[j].Name
	}
}

func Tally(reader io.Reader, writer io.Writer) error {
	var buffer bytes.Buffer
	buffer.ReadFrom(reader)
	input := buffer.String()
	input = strings.Trim(input, "\n")
	rows := strings.Split(input, "\n")
	teamMap := make(map[string]*TeamScore)
	var teamList SortTeam

	for i, row := range rows {
		// Skip blank lines and comments
		if len(strings.Trim(row, " ")) == 0 || row[0] == '#' {
			continue
		}
		record := strings.Split(row, ";")
		if len(record) != 3 || record[0] == record[1] {
			return fmt.Errorf("wrong number of parts on line %d", i)
		}
		name1, name2 := record[0], record[1]
		if _, ok := teamMap[name1]; !ok {
			teamMap[name1] = &TeamScore{name1, 0, 0, 0, 0, 0}
			teamList = append(teamList, teamMap[name1])
		}
		if _, ok := teamMap[name2]; !ok {
			teamMap[name2] = &TeamScore{name2, 0, 0, 0, 0, 0}
			teamList = append(teamList, teamMap[name2])
		}
		team1, team2 := teamMap[name1], teamMap[name2]
		team1.MatchesPlayed += 1
		team2.MatchesPlayed += 1
		switch result := record[2]; result {
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
			return fmt.Errorf("invalid match result \"%s\" in line %d", result, i)
		}
	}

	sort.Sort(teamList)
	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, team := range teamList {
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
