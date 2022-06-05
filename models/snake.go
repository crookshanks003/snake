package models

import (
	"math"
)

type Node struct {
	PosX int
	PosY int
	Next *Node
	Prev *Node
	Dir  int //0: right, 1: down, 2:left, 3:up
}

type Snake struct {
	Head Node
	Tail Node
}

func (node *Node) updatePos() {
	switch node.Dir {
	case 0:
		node.PosX++
	case 1:
		node.PosY++
	case 2:
		node.PosX--
	case 3:
		node.PosY--
	}
}

func NewSnake() *Snake {
	tail := Node{
		PosX: 1,
		PosY: 1,
		Dir:  0,
	}
	snake := Snake{
		Tail: tail,
	}
	head := Node{
		PosX: 8,
		PosY: 1,
		Dir:  0,
		Next: &snake.Tail,
	}
	snake.Head = head
	snake.Tail.Prev = &snake.Head
	return &snake
}

func (s *Snake) UpdatePos() {
	s.Tail.updatePos()
	s.Head.updatePos()
	if s.Tail.Prev.PosX == s.Tail.PosX && s.Tail.Prev.PosY == s.Tail.PosY {
		tail := s.Tail.Prev
		s.Tail = Node{
			PosX: tail.PosX,
			PosY: tail.PosY,
			Dir:  tail.Dir,
			Next: nil,
			Prev: tail.Prev,
		}
		tail.Prev.Next = &s.Tail
	}

}

func (s *Snake) ChangeDir(dir int) {
	if math.Abs(float64(dir)-float64(s.Head.Dir)) == 2 {
		return
	}

	head := s.Head
	head.Dir = dir

	head.Next.Prev = &head

	s.Head = Node{
		PosX: head.PosX,
		PosY: head.PosY,
		Dir:  dir,
		Next: &head,
	}
	head.Prev = &s.Head
}
