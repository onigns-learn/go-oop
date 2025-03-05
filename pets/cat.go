package pets

import (
	"fmt"
	"time"
)

type Cat struct {
	Animal
	Color string
	Breed string
}

func (c Cat) GiveAttention(activity string) string {
	return fmt.Sprintf("%s is ignoring your attempt to %s", c.Name, activity)
}

func NewCat(name, color, breed string) Pet {
	return Cat{
		Color:  color,
		Breed:  breed,
		Animal: Animal{Name: name, lastAte: time.Now()},
	}
}
