package main

import (
	"fmt"
	"time"

	"github.com/onigns-learn/go-oop/pets"
)

func main() {
	pet := pets.NewDog(
		"Drogo",
		"Black and Grey",
		"Artlassan",
		time.Now().Add((time.Duration(-5) * time.Hour)),
	)

	fmt.Println(pet.Feed("steak"))
	fmt.Println(pet.GiveAttention("playing fetch"))
}
