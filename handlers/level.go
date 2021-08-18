package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pedrocmart/maze-go/api/models"
	"github.com/pedrocmart/maze-go/api/restapi/operations"
	"github.com/pedrocmart/maze-go/consts"
	"github.com/pedrocmart/maze-go/repository"
	repoModels "github.com/pedrocmart/maze-go/repository/models"
)

type LevelHandler struct {
	repository repository.Level
}

func NewLevelHandler(levelRepository repository.Level) *LevelHandler {
	return &LevelHandler{
		repository: levelRepository,
	}
}

func (r *LevelHandler) PostLevel(params operations.PostLevelParams) middleware.Responder {
	if params.Payload == nil {
		return operations.NewPostLevelBadRequest().WithPayload(&models.BaseResponse{
			Success: consts.HandlerFailed,
			Message: "Empty body"})
	}

	model := repoModels.NewLevel()
	model, err := model.FromSwaggerModel(params.Payload.Level)
	if err != nil {
		return operations.NewPostLevelInternalServerError().WithPayload(&models.BaseResponse{
			Success: consts.HandlerFailed,
			Message: err.Error()})
	}

	cLevel, cerr := r.repository.Create(model)
	if cerr != nil {
		return operations.NewPostLevelInternalServerError().WithPayload(&models.BaseResponse{
			Success: consts.HandlerFailed,
			Message: cerr.Error()})
	}

	// Return level Created request.
	return operations.NewPostLevelCreated().WithPayload(&models.LevelResponse{
		Message: consts.HandlerMessageSuccess,
		Status:  consts.HandlerStatusCreated,
		Success: true,
		LevelData: struct{ models.Level }{
			models.Level{
				ID: cLevel.Id,
			},
		},
	})
}

// GetLevelID returns the list of levels ordered by created_at descending.
// http GET ":5000/v1/level/1
func (r *LevelHandler) GetLevelID(params operations.GetLevelParams) middleware.Responder {
	levelId := params.LevelID

	// retrieve all the levels for the player
	levels, cerr := r.repository.FindByLevelId(levelId)
	if cerr != nil {
		return operations.NewGetLevelInternalServerError().WithPayload(&models.BaseResponse{
			Success: consts.HandlerFailed,
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
		return operations.NewGetLevelInternalServerError().WithPayload(&models.BaseResponse{
			Success: consts.HandlerFailed,
			Message: cerr.Error()})
	}

	return operations.NewGetLevelOK().WithPayload(payload)
}

func (r *LevelHandler) GetLevel(params operations.GetLevelParams) middleware.Responder {
	levels, cerr := r.repository.FindAll()
	if cerr != nil {
		return operations.NewGetLevelInternalServerError().WithPayload(&models.BaseResponse{
			Success: consts.HandlerFailed,
			Message: cerr.Error()})
	}

	payload := models.LevelAllResponse{}
	var toSwaggerErr error

	for _, level := range levels {
		var model *models.Level
		model, toSwaggerErr = level.ToSwagger()
		if toSwaggerErr != nil {
			break
		}
		payload = append(payload, model)
	}

	if toSwaggerErr != nil {
		return operations.NewGetLevelInternalServerError().WithPayload(
			&models.BaseResponse{
				Success: consts.HandlerFailed,
				Message: toSwaggerErr.Error()})
	}

	return operations.NewGetLevelOK().WithPayload(payload)
}

func RegisterLevelHandlers(api *operations.MazeGoAPI, levelRepository repository.Level) {
	levelHandler := NewLevelHandler(levelRepository)

	// PostLevel saves the given Level data
	api.PostLevelHandler = operations.PostLevelHandlerFunc(levelHandler.PostLevel)
	api.GetLevelHandler = operations.GetLevelHandlerFunc(levelHandler.GetLevel)
}
