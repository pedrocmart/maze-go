package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidMap_ExitWithOpenTile(t *testing.T) {

	array := [][]int32{
		{1, 1, 1, 1, 0, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 3, 1, 1},
		{1, 0, 0, 0, 1, 0, 2, 1},
		{1, 1, 1, 0, 1, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1},
		{1, 0, 0, 4, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	require.NoError(t, err)
}

func TestValidMap_ExitWithPitTrap(t *testing.T) {

	array := [][]int32{
		{1, 1, 1, 1, 2, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 3, 1, 1},
		{1, 0, 0, 0, 1, 0, 2, 1},
		{1, 1, 1, 0, 1, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1},
		{1, 0, 0, 4, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	require.NoError(t, err)
}

func TestValidMap_ExitWithArrowTrap(t *testing.T) {

	array := [][]int32{
		{1, 1, 1, 1, 3, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 3, 1, 1},
		{1, 0, 0, 0, 1, 0, 2, 1},
		{1, 1, 1, 0, 1, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1},
		{1, 0, 0, 4, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	require.NoError(t, err)
}

func TestInvalidMap_NoExit(t *testing.T) {

	array := [][]int32{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 3, 1, 1},
		{1, 0, 0, 0, 1, 0, 2, 1},
		{1, 1, 1, 0, 1, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1},
		{1, 0, 0, 4, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	assert.Error(t, err, "there's no exit")
}

func TestInvalidMap_Rectangle(t *testing.T) {

	array := [][]int32{
		{1, 0},
		{1, 4},
	}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	assert.Error(t, err, "the map must be rectangular")
}

func TestInvalidMap_OutOfRange(t *testing.T) {

	array := [][]int32{
		{1, 5},
		{1, 4},
	}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	assert.Error(t, err, "out of range [0-4]")
}

func TestInvalidMap_WithMoreThanOneStartPoint(t *testing.T) {

	array := [][]int32{
		{1, 0},
		{1, 4},
		{1, 4},
	}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	assert.Error(t, err, "you cant have more than one start point")
}

func TestInvalidMap_WithNoStartPoint(t *testing.T) {

	array := [][]int32{
		{1, 0},
		{1, 0},
		{1, 0},
	}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	assert.Error(t, err, "you must have a start point")
}

func TestInvalidMap_WithNoMap(t *testing.T) {
	array := [][]int32{}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	assert.Error(t, err, "Maps cannot be empty")
}

func TestInvalidMap_MapHaveDifferentColumnNumber(t *testing.T) {

	array := [][]int32{
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 4, 0},
	}
	maps, err := json.Marshal(array)
	require.NoError(t, err)

	model := NewLevel()
	model.Maps = maps
	err = model.Validate()
	assert.Error(t, err, "you need to have the same number of columns for all rows")
}
