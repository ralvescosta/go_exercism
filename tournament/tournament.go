package tournament

import (
	"fmt"
	"io"
	"strings"
)

var competitionResult = `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  $00 |  $01 |  $02 |  $03 |  $04
Allegoric Alaskians            |  $10 |  $11 |  $12 |  $13 |  $14
Blithering Badgers             |  $20 |  $21 |  $22 |  $23 |  $24
Courageous Californians        |  $30 |  $31 |  $32 |  $33 |  $34
`

var timePosition = map[string]int{
	"Devastating Donkeys":     0,
	"Allegoric Alaskians":     1,
	"Blithering Badgers":      2,
	"Courageous Californians": 3,
}

func Tally(reader io.Reader, writer io.Writer) error {
	matrix := make([][]int, 4)
	b := make([]byte, 1024)
	if _, err := reader.Read(b); err != nil {
		return err
	}
	toString := string(b)
	splitInput := strings.Split(toString, "\n")

	for _, v := range splitInput {
		if v == "" {
			continue
		}
		splitSlice := strings.Split(v, ";")
		if len(splitSlice) < 2 {
			continue
		}
		splitSlice[0] = strings.TrimSpace(splitSlice[0])
		splitSlice[1] = strings.TrimSpace(splitSlice[1])

		winierTime := matrix[timePosition[splitSlice[0]]]
		if winierTime == nil {
			winierTime = make([]int, 5)
		}
		loserTime := matrix[timePosition[splitSlice[1]]]
		if loserTime == nil {
			loserTime = make([]int, 5)
		}
		winierTime[0] += 1
		loserTime[0] += 1
		switch splitSlice[2] {
		case "win":
			winierTime[1] += 1
			winierTime[4] += 3
			loserTime[3] += 1
		case "draw":
			winierTime[2] += 1
			winierTime[4] += 1
			loserTime[2] += 1
			loserTime[4] += 1
		case "loss":
			winierTime[3] += 1
			loserTime[1] += 1
			loserTime[4] += 3
		}
		matrix[timePosition[splitSlice[0]]] = winierTime
		matrix[timePosition[splitSlice[1]]] = loserTime
	}

	for i, m := range matrix {
		for j, d := range m {
			competitionResult = strings.ReplaceAll(competitionResult, fmt.Sprintf("$%d%d", i, j), fmt.Sprintf("%d", d))
		}
	}
	fmt.Println(competitionResult)
	writer.Write([]byte(competitionResult[1:]))
	return nil
}
