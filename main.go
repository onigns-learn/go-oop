package main

import (
	"fmt"

	"github.com/onigns-learn/go-oop/pets"
)

func main() {
	pet := pets.Dog{
		Name: "Drogo", Color: "Black and Grey", Breed: "Artlassan",
	}

	fmt.Println(pet.Feed("steak"))
	fmt.Println(pet.GiveAttention("playing fetch"))
}
