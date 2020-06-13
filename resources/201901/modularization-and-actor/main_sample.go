package main

import (
	"sampleactor/samplestruct"
)

func main() {
	hero := samplestruct.NewHero(1, 1, "Hero")
	training := samplestruct.NewTraining("Training")

	training.TrainingHero(hero)
}
