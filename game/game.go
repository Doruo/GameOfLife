package game

type GameState struct {
	Grid   Grid
	Length int
}

func NewGameState(len int) *GameState {
	return &GameState{
		Grid:   *NewSeed(len),
		Length: len,
	}
}
