package pets

import (
	"fmt"
	"strings"
)

type Dog struct {
	Name  string
	Color string
	Breed string
}

func (d Dog) Feed(food string) string {
	return fmt.Sprintf("%s is eating %s", d.Name, food)
}

func (d Dog) GiveAttention(activity string) string {
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
