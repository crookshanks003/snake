package models

type Node struct {
	PosX int
	PosY int
	Next *Node
}

type Snake struct {
	Head    Node
	Tail    Node
	HeadDir int //0: right, 1: down, 2:left, 3:up
	TailDir int //0: right, 1: down, 2:left, 3:up
}

func (node *Node) updatePos(dir int) {
	switch dir {
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
		Next: nil,
	}
	snake := Snake{
		HeadDir: 0,
		TailDir: 0,
		Tail:    tail,
	}
	head := Node{
		PosX: 8,
		PosY: 1,
		Next: &snake.Tail,
	}
	snake.Head = head
	return &snake
}

func (s *Snake) UpdatePos() {
	s.Tail.updatePos(s.TailDir)
	s.Head.updatePos(s.HeadDir)
}

func (s *Snake) UpdateDir(dir int) {
	head := s.Head
	// s.UpdateHeadPos()
	s.Head = Node{
		PosX: head.PosX,
		PosY: head.PosY,
		Next: &head,
	}
	s.HeadDir = dir
}
