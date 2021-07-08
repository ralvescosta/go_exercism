package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	RobotName string
}

var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString() string {
	return string(letters[rand.Intn(26)])
}

func (r *Robot) Name() (string, error) {
	if r.RobotName == "" {
		r1 := rand.Intn(9)
		r2 := rand.Intn(9)
		r3 := rand.Intn(9)
		s1 := randomString()
		s2 := randomString()
		r.RobotName = fmt.Sprintf("%s%s%d%d%d", s1, s2, r1, r2, r3)
	}

	return r.RobotName, nil
}

func (r *Robot) Reset() {
	r.RobotName = ""
}
