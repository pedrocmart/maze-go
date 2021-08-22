package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS_SurvivablePath(t *testing.T) {
	expectedPath := int32(12)
	expectedLife := int32(1)
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
	bfs := BFSInit(array)
	shortestPath, life := bfs.GetSurvivablePath()
	assert.Equal(t, expectedPath, shortestPath)
	assert.Equal(t, expectedLife, life)
}

func TestBFS_SurvivablePath_2(t *testing.T) {
	expectedPath := int32(16)
	expectedLife := int32(4)
	array := [][]int32{
		{1, 1, 1, 1, 0, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1, 0, 1, 1},
		{1, 1, 1, 0, 1, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1},
		{1, 0, 0, 4, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	bfs := BFSInit(array)
	shortestPath, life := bfs.GetSurvivablePath()
	assert.Equal(t, expectedPath, shortestPath)
	assert.Equal(t, expectedLife, life)
}
func TestBFS_NotSurvivablePath(t *testing.T) {
	expectedPath := int32(12)
	expectedLife := int32(0)
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
	bfs := BFSInit(array)
	shortestPath, life := bfs.GetSurvivablePath()
	assert.Equal(t, expectedPath, shortestPath)
	assert.Equal(t, expectedLife, life)
}

func TestBFS_NotSurvivablePath_2(t *testing.T) {
	expectedPath := int32(12)
	expectedLife := int32(-1)
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
	bfs := BFSInit(array)
	shortestPath, life := bfs.GetSurvivablePath()
	assert.Equal(t, expectedPath, shortestPath)
	assert.Equal(t, expectedLife, life)
}
