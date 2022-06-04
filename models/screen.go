package models

import (
	"math"
	"strings"
)

type Screen struct {
	w     int
	h     int
	Snake *Snake
	data  [][]rune
}

func NewScreen(h, w int) Screen {

	data := make([][]rune, h)
	for i := 0; i < h; i++ {
		data[i] = make([]rune, w)
	}

	screen := Screen{
		h:     h,
		w:     w,
		data:  data,
		Snake: NewSnake(),
	}
	screen.RenderSpace()
	screen.RenderBorder()
	screen.RenderSnake()

	return screen
}

func (s *Screen) RenderSpace() {
	for row := 1; row < s.h-1; row++ {
		for col := 1; col < s.w-1; col++ {
			s.data[row][col] = ' '
		}
	}
}

func (s *Screen) RenderBorder() {
	for col := 1; col < s.w-1; col++ {
		s.data[0][col] = horizontal
		s.data[s.h-1][col] = horizontal
	}

	for row := 1; row < s.h-1; row++ {
		s.data[row][0] = vertical
		s.data[row][s.w-1] = vertical
	}

	s.data[0][0] = topLeft
	s.data[0][s.w-1] = topRight
	s.data[s.h-1][0] = bottomLeft
	s.data[s.h-1][s.w-1] = bottomRight
}

func (s *Screen) RenderSnake() {
	node := s.Snake.Head

	for node.Next != nil {
		next := node.Next
		var start, end int

		if next.PosY == node.PosY {
			start = int(math.Min(float64(node.PosX), float64(next.PosX)))
			end = int(math.Max(float64(node.PosX), float64(next.PosX)))
			for i := start; i <= end; i++ {
				s.data[node.PosY][i] = snakeRune
			}
		} else {
			start = int(math.Min(float64(node.PosY), float64(next.PosY)))
			end = int(math.Max(float64(node.PosY), float64(next.PosY)))
			for i := start; i <= end; i++ {
				s.data[i][node.PosX] = snakeRune
			}
		}
		node = *node.Next
	}

	s.data[s.Snake.Head.PosY][s.Snake.Head.PosX] = headRune
}

func (s *Screen) UpdateSnakePos() {
	s.Snake.UpdatePos()
	s.RenderSpace()
	s.RenderSnake()
}

func render() {

}

func (s *Screen) Render() string {
	rowStrings := []string{}
	for _, row := range s.data {
		rowStrings = append(rowStrings, string(row))
	}
	return strings.Join(rowStrings, "\n")
}
