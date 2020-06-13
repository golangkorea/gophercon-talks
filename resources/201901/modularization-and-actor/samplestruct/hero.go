package samplestruct

import "fmt"

type Hero struct {
	id       int
	level    int
	nodeName string
}

func NewHero(id int, level int, nodeName string) *Hero {
	return &Hero{
		id:       id,
		level:    level,
		nodeName: nodeName,
	}
}

func (h *Hero) GetLevel() int {
	return h.level
}

func (h *Hero) SetLevel(level int) {
	h.level = level
	fmt.Printf("Hero.SetLevel: Level %d \n", level)
}

func (h *Hero) GetNodeName() string {
	return h.nodeName
}
