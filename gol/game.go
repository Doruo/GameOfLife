package gol

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	title                     string
	screenWidth, screenHeight int
	gs                        GameState
}

func NewGame(size int) *Game {
	return &Game{
		title:        "Game of life",
		screenWidth:  size,
		screenHeight: size,
		gs:           *NewGameState(size),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.title)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) Run() {

	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle(g.title)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
