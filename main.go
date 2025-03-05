package main

import (
	"fmt"
	"time"

	"github.com/onigns-learn/go-oop/pets"
)

func main() {

	sleepTime := time.Now().Add((time.Duration(-5) * time.Hour))

	var myPets []pets.Pet

	myPets = append(myPets, pets.NewDog(
		"Drogo",
		"Black and Grey",
		"Artlassan",
		sleepTime,
	))

	myPets = append(myPets, pets.NewCat(
		"Felixx",
		"Black and Grey",
		"Artlassan",
	))

	for _, pet := range myPets {
		if pet.IsHungry() {
			fmt.Println(pet.Feed("steak"))
		} else {
			fmt.Printf("%s isn't hungry", pet.GetName())
			time.Sleep(5 * time.Second)
			fmt.Println(pet.Feed("Kibble"))
		}
		fmt.Println(pet.GiveAttention("playing fetch"))
	}
}
