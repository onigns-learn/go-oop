package pets

import (
	"fmt"
	"strings"
	"time"
)

type Dog struct {
	Animal
	Color     string
	Breed     string
	lastSlept time.Time
}

func (d Dog) needsSleep() bool {
	return time.Since(d.lastSlept) > 4*time.Hour
}

func (d Dog) sleep() {
	d.lastSlept = time.Now()
}

func (d Dog) FeedDog(food string) string {
	return fmt.Sprintf("%s is eating %s", d.Name, food)
}

func (d Dog) GiveAttention(activity string) string {
	if d.needsSleep() {
		d.sleep()
		return "You dog is asleep"
	}

	response := ""

	switch strings.ToUpper(activity) {
	case "PETS":
		response = "wags tail"
	case "PLAYING FETCH":
		response = "return the ball and ump waiting for you to throw again"
	default:
		response = "barks"
	}

	return fmt.Sprintf("%s loves attention, %s will cause him to %s", d.Name, activity, response)
}

func NewDog(name, color, breed string, lastSlept time.Time) Pet {
	return Dog{
		Color:     color,
		Breed:     breed,
		lastSlept: lastSlept,
		Animal:    Animal{Name: name, lastAte: time.Now()},
	}
}
