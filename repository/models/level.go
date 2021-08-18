package models

import (
	"encoding/json"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/pedrocmart/maze-go/api/models"
)

type Level struct {
	Id int64 `json:"id"` // database record ID

	Maps      json.RawMessage `json:"resRows"`
	CreatedAt time.Time       `json:"createdAt"`
}

func NewLevel() *Level {
	return &Level{}
}

func (p *Level) FromSwaggerModel(in models.Level) (*Level, error) {
	level := &Level{
		Id: p.Id,
	}

	if in.Maps != nil {
		maps, err := json.Marshal(in.Maps)
		if err != nil {
			return nil, err
		}

		level.Maps = maps
	}

	return level, nil
}

func (p *Level) ToSwagger() (*models.Level, error) {
	createdAt := strfmt.DateTime(p.CreatedAt)

	level := &models.Level{
		ID:        p.Id,
		CreatedAt: &createdAt,
	}

	var data [][]int64
	if err := json.Unmarshal(p.Maps, &data); err != nil {
		return nil, err
	}
	level.Maps = data

	return level, nil
}

// func (p *Level) Validate() error {
// 	//Validate max length fields
// 	if len(p.PlayerID) > 100 {
// 		return adapters.ReturnInvalidMaxLengthError("PlayerID", 100)
// 	}

// 	//Validate max length fields
// 	if len(p.SessionID) > 100 {
// 		return adapters.ReturnInvalidMaxLengthError("SessionID", 100)
// 	}

// 	//Validate max length fields
// 	if len(*p.TournamentID) > 100 {
// 		return adapters.ReturnInvalidMaxLengthError("TournamentID", 100)
// 	}

// 	return nil
// }
