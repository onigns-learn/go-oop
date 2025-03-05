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

	if pet.IsHungry() {
		fmt.Println(pet.Feed("steak"))
	} else {
		fmt.Printf("%s isn't hungry", pet.Name)
		time.Sleep(5 * time.Second)
		fmt.Println(pet.Feed("Kibble"))
	}

	fmt.Println(pet.GiveAttention("playing fetch"))
}
