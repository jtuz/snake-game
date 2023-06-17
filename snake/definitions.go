package snake

const (
	ScreenWidth  = 320
	ScreenHeight = 240
	TileSize     = 5
)

type Point struct {
	X int
	Y int
}

type Snake struct {
	Body        []Point
	Direction   Point
	GrowCounter int
}

type Food struct {
	Position Point
}

type Game struct {
	snake         *Snake
	food          *Food
	score         int
	gameOver      bool
	ticks         int
	updateCounter int
	speed         int
}
