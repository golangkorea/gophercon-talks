package samplestruct

import "fmt"

type Training struct {
	nodeName string
}

func NewTraining(nodeName string) *Training {
	return &Training{
		nodeName: nodeName,
	}
}

func (t *Training) TrainingHero(hero *Hero) {
	level := hero.GetLevel()
	level++
	hero.SetLevel(level)
}

func (t *Training) Do(level int) int {
	newLevel := level + 1
	fmt.Printf("Training.Do: oldLevel %d, nowLevel %d \n", level, newLevel)
	return newLevel
}

func (t *Training) GetNodeName() string {
	return t.nodeName
}
