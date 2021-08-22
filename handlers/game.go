package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pedrocmart/maze-go/api/models"
	"github.com/pedrocmart/maze-go/api/restapi/operations"
	"github.com/pedrocmart/maze-go/consts"
	"github.com/pedrocmart/maze-go/internal/logic"
	"github.com/pedrocmart/maze-go/repository"
)

type GameHandler struct {
	repository repository.Level
}

func NewGameHandler(levelRepository repository.Level) *GameHandler {
	return &GameHandler{
		repository: levelRepository,
	}
}

// GetLevelID returns the list of levels ordered by created_at descending.
// http GET ":5000/v1/game/1
func (r *GameHandler) GetGameID(params operations.GetGameLevelIDParams) middleware.Responder {
	levelId := params.LevelID

	// retrieve all the levels for the player
	levels, cerr := r.repository.FindByLevelId(levelId)
	if cerr != nil {
		return operations.NewGetGameLevelIDInternalServerError().WithPayload(&models.BaseResponse{
			Success: consts.HandlerFailed,
			Status:  consts.HandlerStatusCodeInternalServerError,
			Message: cerr.Error()})
	}

	payload := models.LevelAllResponse{}
	var errors []error

	for _, level := range levels {
		model, errModel := level.ToSwagger()
		if errModel != nil {
			errors = append(errors, errModel)
			break
		}
		payload = append(payload, model)
	}

	if len(errors) > 0 {
		return operations.NewGetGameLevelIDInternalServerError().WithPayload(&models.BaseResponse{
			Success: consts.HandlerFailed,
			Status:  consts.HandlerStatusCodeInternalServerError,
			Message: cerr.Error()})
	}

	bfs := logic.BFSInit(payload[0].Maps)
	shortestPath, life := bfs.GetSurvivablePath()
	if life <= 0 {
		return operations.NewGetGameLevelIDInternalServerError().WithPayload(&models.BaseResponse{
			Success: consts.HandlerSuccess,
			Status:  consts.HandlerStatusCodeOK,
			Message: fmt.Sprintf("You died. Life %d", life)})
	}
	if shortestPath <= 0 {
		return operations.NewGetGameLevelIDInternalServerError().WithPayload(&models.BaseResponse{
			Success: consts.HandlerSuccess,
			Status:  consts.HandlerStatusCodeOK,
			Message: "It was not possible to find an exit"})
	}

	return operations.NewGetGameLevelIDOK().WithPayload(&models.BaseResponse{
		Success: consts.HandlerSuccess,
		Status:  consts.HandlerStatusCodeOK,
		Message: fmt.Sprintf("The survivable path is %d and your life is %d", shortestPath, life)})
}

func RegisterGameHandlers(api *operations.MazeGoAPI, levelRepository repository.Level) {
	gameHandler := NewGameHandler(levelRepository)

	api.GetGameLevelIDHandler = operations.GetGameLevelIDHandlerFunc(gameHandler.GetGameID)
}
