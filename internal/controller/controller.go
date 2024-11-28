package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"td/internal/logic"
	"td/internal/respond"
)


type GameController struct {
	Logic logic.GameLogic
	Responder respond.Responder
}

func NewGameController(l logic.GameLogic, r respond.Responder) GameController {
	return GameController{
		Logic: l,
		Responder: r,
	}
}

func (gc GameController) State(w http.ResponseWriter, r *http.Request) {
	game := gc.Logic.State()
	gc.Responder.JSON(w, http.StatusOK, game)
}

type AddTowerRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (gc GameController) AddTower(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		gc.Responder.Error(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	// Validate the JSON contains only `x` and `y` fields
	var body map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		gc.Responder.Error(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Ensure only `x` and `y` fields exist
	if len(body) != 2 {
		gc.Responder.Error(w, http.StatusBadRequest, "Request must contain only 'x' and 'y'")
		return
	}
	if _, ok := body["x"]; !ok {
		gc.Responder.Error(w, http.StatusBadRequest, "'x' field is required")
		return
	}
	if _, ok := body["y"]; !ok {
		gc.Responder.Error(w, http.StatusBadRequest, "'y' field is required")
		return
	}

	// Decode into AddTowerRequest struct
	var req AddTowerRequest
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		gc.Responder.Error(w, http.StatusBadRequest, "Failed to parse request")
		return
	}

	// Add the tower using the provided coordinates
	gc.Logic.AddTower(req.X, req.Y)

	// Respond with success
	gc.Responder.JSON(w, http.StatusOK, map[string]string{"message": "Tower added successfully"})
}




func (gc GameController) StartGame(w http.ResponseWriter, r *http.Request) {
	gc.Logic.StartGame()
	gc.Responder.JSON(w, http.StatusOK, nil)
}


func (gc GameController) Reset(w http.ResponseWriter, r *http.Request) {
	gc.Logic.Reset()
	gc.Responder.JSON(w, http.StatusOK, nil)
}
