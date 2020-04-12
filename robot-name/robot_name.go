package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var used = map[string]bool{}

type Robot struct {
	name string
}

func generateName() string {
	rand.Seed(time.Now().UnixNano())

	firstLeter := 'A' + byte(rand.Intn(26))
	secondLeter := 'A' + byte(rand.Intn(26))
	firstN := rand.Intn(10)
	secondN := rand.Intn(10)
	thirdN := rand.Intn(10)

	return fmt.Sprintf("%s%s%d%d%d", string(firstLeter), string(secondLeter), firstN, secondN, thirdN)
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.Reset()
	} else {
		return r.name, nil
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = generateName()
	for used[r.name] {
		r.name = generateName()
	}

	used[r.name] = true
}
