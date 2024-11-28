package controller

import (
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

