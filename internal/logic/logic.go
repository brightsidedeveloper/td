package logic

import (
	"fmt"
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
	Round		int 	`json:"round"`
}

type GameMethods interface {
	State() Game
	StartGame()
	StopGame()
	AddTower(x int, y int)
	Reset()
}

type GameLogic struct {
	game *Game
	path [][]int
	spawnCh chan Enemy
}

func (g *GameLogic) State() Game {
	return *g.game
}

func (g *GameLogic) startGameLoop() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		if !g.game.IsRunning {
			fmt.Println("Game stopped")
			break
		}
		g.moveEnemies()
		g.updateTowers()
	}
}


func (g *GameLogic) moveEnemies() {
	// Iterate through all cells in reverse path order to avoid overwriting moves
	for idx := len(g.path) - 1; idx >= 0; idx-- {
		x, y := g.path[idx][0], g.path[idx][1]
		cell := &g.game.Grid[x][y] // Pointer to the current cell

		if len(cell.Enemies) > 0 {
			var remainingEnemies []Enemy // To keep enemies that can't move

			for _, enemy := range cell.Enemies {
				// If this is the last step in the path, remove the enemy
				if idx == len(g.path)-1 {
					fmt.Printf("Enemy %v reached the end at (%d, %d)\n", enemy.Key, x, y)
					g.game.PlayerHealth -= 1 // Reduce player health
					if g.game.PlayerHealth <= 0 {
						g.StopGame()
						fmt.Println("Game Over!")
					}
					continue
				}

				// Move enemy to the next position in the path
				nextX, nextY := g.path[idx+1][0], g.path[idx+1][1]
				g.game.Grid[nextX][nextY].Enemies = append(g.game.Grid[nextX][nextY].Enemies, enemy)
				fmt.Printf("Enemy %v moved from (%d, %d) to (%d, %d)\n", enemy.Key, x, y, nextX, nextY)
			}

			// Update current cell's enemies to those that couldn't move
			cell.Enemies = remainingEnemies
		}
	}
}




func (g *GameLogic) updateTowers() {
	for i, row := range g.game.Grid {
		for j, cell := range row {
			if cell.Tower.IsActive {
				if time.Now().UnixNano()-cell.Tower.LastFired > 1_000_000_000 { // 1 second
					fmt.Printf("Tower at (%d, %d) fires!\n", i, j)
					g.game.Grid[i][j].Tower.LastFired = time.Now().UnixNano()
				}
			}
		}
	}
}

func (g *GameLogic) NextRound() {
	g.game.Round++
}

func (g *GameLogic) spawnEnemies() {
	count := g.game.Round * 2
	go func() {
		for i := 0; i < count; i++ {
			enemy := Enemy{Key: uuid.New()}
			g.spawnCh <- enemy
			time.Sleep(2 * time.Second) // Spawn 1 enemy per second
		}
	}()
}

func (g *GameLogic) AddTower(x int, y int) {
	for i, row := range g.game.Grid {
		for j, cell := range row {
			if cell.X == x && cell.Y == y {
				g.game.Grid[i][j].Tower.IsActive = true
				g.game.Grid[i][j].Tower.LastFired = time.Now().UnixNano()
			}
		}
	}
}

func (g *GameLogic) addEnemyToGrid() {
	for enemy := range g.spawnCh {
		start := g.path[0]
		g.game.Grid[start[0]][start[1]].Enemies = append(
			g.game.Grid[start[0]][start[1]].Enemies, enemy)
		fmt.Printf("Enemy %v added to (%d, %d)\n", enemy.Key, start[0], start[1])
	}
}

func (g *GameLogic) StartGame() {
	g.game.IsRunning = true
	go g.startGameLoop()
	go g.spawnEnemies()
	go g.addEnemyToGrid()
}

func (g *GameLogic) StopGame() {
	g.game.IsRunning = false
	close(g.spawnCh)
}


func (g *GameLogic) Reset() {
	*g.game = *NewGame()
}

func containsPath(path [][]int, x, y int) bool {
	for _, p := range path {
		if p[0] == x && p[1] == y {
			return true
		}
	}
	return false
}

var pathMatrix = [][]int{
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
    {5, 6},
	{5, 7},
}

func NewGame() *Game {
	grid := make([][]Cell, 8)
	for i := 0; i < 8; i++ {
		row := make([]Cell, 8)
		for j := 0; j < 8; j++ {
			row[j] = Cell{
				Key:    uuid.New(),
				X:      i,
				Y:      j,
				IsPath: containsPath(pathMatrix, i, j),
			}
		}
		grid[i] = row
	}

	return &Game{
		Grid:         grid,
		PlayerHealth: 100,
		IsRunning:    false,
		Round:        1,
	}
}

func NewGameLogic(game *Game) GameLogic {
	return GameLogic{
		game: game,
		path: pathMatrix,
		spawnCh: make(chan Enemy, 10),
	}
}
