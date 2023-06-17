package snake

import (
	"math/rand"
	"time"
)

func NewSnake() *Snake {
	return &Snake{
		Body: []Point{
			{X: ScreenWidth / TileSize / 2, Y: ScreenHeight / TileSize / 2},
		},
		Direction: Point{X: 1, Y: 0},
	}
}

func (s *Snake) Move() {
	newHead := Point{
		X: s.Body[0].X + s.Direction.X,
		Y: s.Body[0].Y + s.Direction.Y,
	}
	s.Body = append([]Point{newHead}, s.Body...)

	if s.GrowCounter > 0 {
		s.GrowCounter--
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}

func NewFood() *Food {
	rand.Seed(time.Now().UnixNano())
	return &Food{
		Position: Point{
			X: rand.Intn(ScreenWidth / TileSize),
			Y: rand.Intn(ScreenHeight / TileSize),
		},
	}
}
