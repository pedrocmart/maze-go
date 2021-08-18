package models

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/pedrocmart/maze-go/api/models"
)

type Level struct {
	Id int64 `json:"id"` // database record ID

	Maps      [][]int64 `json:"resRows"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewLevel() *Level {
	return &Level{}
}

func (p *Level) FromSwaggerModel(in models.Level) (*Level, error) {
	round := &Level{
		Maps: in.Maps,
	}

	// if in.Spins != nil {
	// 	spins, err := json.Marshal(in.Spins)
	// 	if err != nil {
	// 		// TODO err or strong validation
	// 	}

	// 	round.Spins = spins
	// }

	// if in.FreeGames != nil {
	// 	freeGames, err := json.Marshal(in.FreeGames)
	// 	if err != nil {
	// 		// TODO err or strong validation
	// 	}

	// 	round.FreeGames = freeGames
	// }

	return round, nil
}

func (p *Level) ToSwagger() (*models.Level, error) {
	createdAt := strfmt.DateTime(p.CreatedAt)

	// var spinData []*models.Level
	// if err := json.Unmarshal(p.Spins, &spinData); err != nil {
	// 	return nil, err
	// }

	// var freeSpinData []*models.Spin
	// if err := json.Unmarshal(p.Spins, &freeSpinData); err != nil {
	// 	return nil, err
	// }

	round := &models.Level{
		Maps:      p.Maps,
		ID:        p.Id,
		CreatedAt: &createdAt,
	}

	return round, nil
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
