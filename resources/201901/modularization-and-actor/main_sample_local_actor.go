package main

import (
	"sampleactor/actor"
	"sampleactor/samplestruct"
)

func main() {
	hero := samplestruct.NewHero(1, 1, "Hero")
	//training := samplestruct.NewTraining("Training")

	heroAID := actor.StartActor(hero)
	//trainingAID := actor.StartActor(training)
	trainingAID := actor.StartWebActor(nil, "Training", "http://localhost:9999")

	//GetLevel() int
	results, _ := actor.Call(heroAID, (*samplestruct.Hero).GetLevel)

	// Do(level int) int
	//results, _ = actor.Call(trainingAID, (*samplestruct.Training).Do, results[0])
	results, _ = actor.Call(trainingAID, "Do", results[0])

	//SetLevel(level int)
	_, _ = actor.Call(heroAID, (*samplestruct.Hero).SetLevel, results[0])
}
