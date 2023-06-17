package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"log"

	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jtuz/snake-game/snake"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	snake         *snake.Snake
	food          *snake.Food
	score         int
	gameOver      bool
	ticks         int
	updateCounter int
	speed         int
}

func NewGame() *Game {
	return &Game{
		snake:    snake.NewSnake(),
		food:     snake.NewFood(),
		gameOver: false,
		ticks:    0,
		speed:    10,
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.restart()
		}

		return nil
	}
	g.updateCounter++
	if g.updateCounter < g.speed {
		return nil
	}
	g.updateCounter = 0
	g.snake.Move()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.snake.Direction.X == 0 {
		g.snake.Direction = snake.Point{X: -1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && g.snake.Direction.X == 0 {
		g.snake.Direction = snake.Point{X: 1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) && g.snake.Direction.Y == 0 {
		g.snake.Direction = snake.Point{X: 0, Y: -1}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && g.snake.Direction.Y == 0 {
		g.snake.Direction = snake.Point{X: 0, Y: 1}
	}

	head := g.snake.Body[0]
	if head.X < 0 || head.Y < 0 || head.X >= snake.ScreenWidth/snake.TileSize ||
		head.Y >= snake.ScreenHeight/snake.TileSize {
		g.gameOver = true
		g.speed = 10
	}

	for _, part := range g.snake.Body[1:] {
		if head.X == part.X && head.Y == part.Y {
			g.gameOver = true
			g.speed = 10
		}
	}

	if head.X == g.food.Position.X && head.Y == g.food.Position.Y {
		g.score++
		g.snake.GrowCounter += 1
		g.food = snake.NewFood()
		g.score++
		g.food = snake.NewFood()
		g.snake.GrowCounter += 1

		if g.speed > 2 {
			g.speed--
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw Snake
	for _, p := range g.snake.Body {
		ebitenutil.DrawRect(
			screen,
			float64(p.X*snake.TileSize),
			float64(p.Y*snake.TileSize),
			snake.TileSize,
			snake.TileSize,
			color.RGBA{0, 255, 0, 255},
		)
	}
	// Draw food
	ebitenutil.DrawRect(
		screen,
		float64(g.food.Position.X*snake.TileSize),
		float64(g.food.Position.Y*snake.TileSize),
		snake.TileSize,
		snake.TileSize,
		color.RGBA{255, 0, 255, 0},
	)

	face := basicfont.Face7x13

	if g.gameOver {
		text.Draw(
			screen,
			"Game Over",
			face,
			snake.ScreenWidth/2-40,
			snake.ScreenHeight/2,
			color.White,
		)
		text.Draw(
			screen,
			"Press 'R' to restart",
			face,
			snake.ScreenWidth/2-60,
			snake.ScreenHeight/2+16,
			color.White,
		)
	}

	scoreDraw := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreDraw, face, 5, snake.ScreenHeight-5, color.White)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) restart() {
	g.snake = snake.NewSnake()
	g.score = 0
	g.gameOver = false
	g.food = snake.NewFood()
	g.speed = 10
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(snake.ScreenWidth*2, snake.ScreenHeight*2)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
