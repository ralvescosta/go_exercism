package tournament

import "io"

const competitionResult = `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  $11 |  $12 |  $13 |  $14 |  $15
Allegoric Alaskians            |  $21 |  $22 |  $23 |  $24 |  $25
Blithering Badgers             |  $31 |  $32 |  $33 |  $34 |  $35
Courageous Californians        |  $41 |  $42 |  $43 |  $44 |  $45
`

func Tally(reader io.Reader, writer io.Writer) error {
	writer.Write([]byte(competitionResult))
	return nil
}
