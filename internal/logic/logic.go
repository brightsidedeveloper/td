package logic

type Cannon struct {
	Key       string `json:"key"`
	LastFired int64  `json:"lastFired"`
}

type Cell struct {
	Key     string  `json:"key"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	IsPath  bool    `json:"isPath"`
	Enemies []Enemy `json:"enemies"`
	Cannon  Cannon  `json:"cannon"`
}

type Enemy struct {
	Key string `json:"key"`
}

type Logic struct {
	Grid         [][]Cell `json:"grid"`
	PlayerHealth int      `json:"playerHealth"`
	IsRunning    bool     `json:"isRunning"`
}