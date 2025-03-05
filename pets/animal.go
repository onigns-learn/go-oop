package pets

import (
	"fmt"
	"time"
)

type Animal struct {
	lastAte time.Time
}

func (a Animal) Feed(food string) string {
	a.lastAte = time.Now()
	return fmt.Sprintf("The Animal is eating %s", food)
}

func (a Animal) IsHungry() bool {
	return time.Since(a.lastAte) > 5*time.Minute
}
