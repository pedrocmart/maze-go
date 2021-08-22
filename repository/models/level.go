package models

import (
	"encoding/json"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/pedrocmart/maze-go/api/models"
	"github.com/pedrocmart/maze-go/consts"
	"github.com/pedrocmart/maze-go/internal/adapters"
	"github.com/pedrocmart/maze-go/internal/logic"
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

	var data [][]int32
	if err := json.Unmarshal(p.Maps, &data); err != nil {
		return nil, err
	}
	level.Maps = data

	return level, nil
}

func (p *Level) Validate() error {

	if p.Maps == nil || len(p.Maps) == 0 {
		return adapters.ReturnEmptyFieldError("Maps")
	}

	var maps [][]int32
	_ = json.Unmarshal(p.Maps, &maps)

	numRows := len(maps)
	if numRows == 0 {
		return adapters.ReturnEmptyFieldError("Maps")
	}
	if numRows > 100 {
		return adapters.ReturnInvalidMaxLengthError("Maps", 100)
	}
	numColumns := len(maps[0])

	oneStartPoint := false

	for _, r := range maps {
		cerr := checkColumnsHaveSameSize(r, numColumns)
		if cerr != nil {
			return cerr
		}
		//check if row is greater than 100
		lerr := checkRowMaxLength(r)
		if lerr != nil {
			return lerr
		}

		for _, c := range r {

			//check if map is inside range [0-4]
			ierr := checkIfMapIsInsideRange(c)
			if ierr != nil {
				return ierr
			}

			//check if map have only one start point
			if c == consts.StartPosition {
				if oneStartPoint {
					return adapters.ReturnError("you cant have more than one start point")
				}
				oneStartPoint = true
			}

		}
	}

	//TODO check if the map has exit
	xerr := checkIfMapHasExit(maps, numRows, numColumns, consts.OpenTile)
	if xerr != nil {
		return xerr
	}

	//check if map have start point
	if !oneStartPoint {
		return adapters.ReturnError("you must have a start point")
	}

	//check if map is a rectangle
	rerr := checkIfMapIsRectangle(numRows, numColumns)
	if rerr != nil {
		return rerr
	}

	return nil
}

/*check if columns have same size. you cant have rows with different column size
example: [[1, 4, 1], [0, 0, 0]] -> valid
		 [[1, 4, 1], [0, 0, 0, 1]] -> invalid
*/
func checkColumnsHaveSameSize(row1 []int32, numColumns int) error {
	if len(row1) != numColumns {
		return adapters.ReturnError("you need to have the same number of columns for all rows")
	}
	return nil
}

func checkRowMaxLength(row1 []int32) error {
	if len(row1) > 100 {
		return adapters.ReturnInvalidMaxLengthError("Maps", 100)
	}
	return nil
}

func checkIfMapIsInsideRange(i int32) error {
	if i < consts.OpenTile || i > consts.StartPosition {
		return adapters.ReturnError("out of range [0-4]")
	}
	return nil
}

func checkIfMapIsRectangle(numLines int, numColumns int) error {
	if numLines == numColumns {
		return adapters.ReturnError("the map must be rectangular")
	}
	return nil
}

func checkIfMapHasExit(maps [][]int32, rows int, cols int, exit int32) error {
	_, err := logic.FindExit(maps, rows, cols, exit)
	return err
}
