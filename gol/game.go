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

func NewGame(width, height int) *Game {
	return &Game{
		title:        "Game of life",
		screenWidth:  width,
		screenHeight: height,
		gs:           *NewGameState(width * height),
	}
}

func (g *Game) Update() error {
	g.gs.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.title)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Run() {

	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle(g.title)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
