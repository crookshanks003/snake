package models

import (
	"math/rand"
	"time"
		
	"github.com/crookshanks003/snake/cons"
)

type Bread struct {
	PosX int
	PosY int
}

func NewBread() *Bread {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	PosX := r1.Intn(cons.Width - 2) + 1
	PosY := r1.Intn(cons.Height - 2) + 1

	return &Bread{
		PosX,
		PosY,
	}
}
