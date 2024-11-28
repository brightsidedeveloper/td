package controller

import (
	"bsdserv/logic"
	"bsdserv/respond"
	"net/http"
)


func NewTestController(l logic.TestLogic, r respond.Responder) TestController {
	return TestController{
		Logic: l,
		Responder: r,
	}
}

type TestController struct {
	Logic logic.TestLogic
	Responder respond.Responder
}

func (tc TestController) Readiness(w http.ResponseWriter, r *http.Request) {
	tc.Responder.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (tc TestController) Error(w http.ResponseWriter, r *http.Request) {
	tc.Responder.Error(w, http.StatusInternalServerError, "An error occurred")
}