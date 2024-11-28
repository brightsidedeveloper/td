package logic

import (
	"time"

	"github.com/google/uuid"
)

type Tower struct {
	Key       uuid.UUID `json:"key"`
	IsActive  bool   `json:"isActive"`
	LastFired int64  `json:"lastFired"`
	X 	   int    `json:"x"`
	Y 	   int    `json:"y"`
}

type Cell struct {
	Key     uuid.UUID  `json:"key"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	IsPath  bool    `json:"isPath"`
	Enemies []Enemy `json:"enemies"`
	Tower  Tower  `json:"tower"`
}

type Enemy struct {
	Key uuid.UUID `json:"key"`
}

type Game struct {
	Grid         [][]Cell `json:"grid"`
	PlayerHealth int      `json:"playerHealth"`
	IsRunning    bool     `json:"isRunning"`
}

type GameMethods interface {
	State() Game
	StartGame()
	StopGame()
	AddTower(x int, y int)
	Reset()
}

type GameLogic struct {
	game Game
}

func (g *GameLogic) State() Game {
	return g.game
}

func (g *GameLogic) startGameLoop() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		if !g.game.IsRunning {
			break
		}

	}
}

func (g *GameLogic) updateCell() {
	for i, row := range g.game.Grid {
}

func (g *GameLogic) updateEnemy() {

}

func ()

func (g *GameLogic) StartGame() {
	g.game.IsRunning = true
	go g.startGameLoop()
}

func (g *GameLogic) StopGame() {
	g.game.IsRunning = false
}

func (g *GameLogic) AddTower(x int, y int) {
	for i, row := range g.game.Grid {
		for j, cell := range row {
			if cell.X == x && cell.Y == y {
				g.game.Grid[i][j].Tower.IsActive = true
			}
		}
	}
}

func (g *GameLogic) Reset() {
	g.game = NewGame()
}

func NewGame() Game {
	path := [][]int{
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 2},
		{3, 2},
		{3, 3},
		{3, 4},
		{3, 5},
		{4, 5},
		{5, 5},
		{5, 4},
		{5, 3},
		{5, 2},
		{5, 1},
		{5, 0},
		{6, 0},
		{7, 0},
		{7, 1},
		{7, 2},
		{7, 3},
		{7, 4},
		{7, 5},
		{7, 6},
		{7, 7},
		{6, 7},
		{5, 7},
		{4, 7},
		{3, 8},
	}
	grid := make([][]Cell, 8)
	for i := 0; i < 8; i++ {
		row := make([]Cell, 8)
		for j := 0; j < 8; j++ {
			key_1, err := uuid.NewRandom()
			if (err != nil) {
				panic(err)
			}
			key_2, err := uuid.NewRandom()
			if (err != nil) {
				panic(err)
			}

			isPath := false
			for _, p := range path {
				if p[0] == i && p[1] == j {
					isPath = true
					break
				}
			}

			row[j] = Cell{
				Key:     key_1,
				X:       i + 1,
				Y:       j + 1,
				IsPath:  isPath, 
				Enemies: []Enemy{},
				Tower:  Tower{
					Key: key_2,
					X: i + 1,
					Y: j + 1,
				},
			}
		}
		grid[i] = row
	}

	return Game{
		Grid:         grid,
		PlayerHealth: 100,
		IsRunning:    false,
	}
}

func NewGameLogic(game Game) GameLogic {
	return GameLogic{
		game: game,
	}
}
