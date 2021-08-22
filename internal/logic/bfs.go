package logic

import (
	"container/list"
	"errors"

	"github.com/pedrocmart/maze-go/consts"
)

// Structure representing X and Y points in the map
type MapPoint struct {
	x, y int
}

// QueueMapPoint contains information about player's life and distance from the origin
type QueueMapPoint struct {
	p    MapPoint
	dist int32
	life int32
}

type BFSMap struct {
	maze       [][]int32
	visited    [][]bool
	rows, cols int
	start, end MapPoint
	queue      *list.List
}

func BFSInit(arr [][]int32) BFSMap {
	rows := len(arr)
	cols := len(arr[0])
	start := findPosition(arr, rows, cols, consts.StartPosition)
	exit, err := FindExit(arr, rows, cols, consts.OpenTile)

	if err != nil {
		panic(err.Error())
	}

	bfs := BFSMap{
		maze:    arr,
		visited: make([][]bool, rows),
		rows:    rows,
		cols:    cols,
		start:   start,
		end:     exit,
		queue:   list.New(),
	}

	for i := 0; i < rows; i++ {
		bfs.visited[i] = make([]bool, cols)
	}

	bfs.queue.PushBack(QueueMapPoint{p: start, dist: 0, life: consts.PlayerLife})
	bfs.visited[start.y][start.x] = true
	return bfs
}

// findPosition gets the coordinates of the position in the array and returns a MapPoint
func findPosition(array [][]int32, rows int, columns int, position int32) MapPoint {
	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			if array[y][x] == position {
				return MapPoint{x, y}
			}
		}
	}
	return MapPoint{-1, -1}
}

//FindExit checks if the Maze has an exit on the edges
func FindExit(a [][]int32, rows int, cols int, exit int32) (MapPoint, error) {
	result := MapPoint{-1, -1}
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if y == 0 && a[y][x] == exit {
				result = MapPoint{x, y}
				return MapPoint{x, y}, nil
			} else if y == rows-1 && a[y][x] == exit {
				result = MapPoint{x, y}
				return MapPoint{x, y}, nil
			} else if x == 0 && a[y][x] == exit {
				result = MapPoint{x, y}
				return MapPoint{x, y}, nil
			} else if x == cols-1 && a[y][x] == exit {
				result = MapPoint{x, y}
				return MapPoint{x, y}, nil
			}
		}
	}
	//if dont find an exit with OpenTile, try to find it with a PitTrap
	if result.x == -1 && result.y == -1 && exit == consts.OpenTile {
		return FindExit(a, rows, cols, consts.PitTrap)
	}
	//if dont find an exit with PitTrap, try to find it with a ArrowTrap
	if result.x == -1 && result.y == -1 && exit == consts.PitTrap {
		return FindExit(a, rows, cols, consts.ArrowTrap)
	}
	if result.x == -1 && result.y == -1 && exit == consts.ArrowTrap {
		return result, errors.New("there's no exit")
	}
	return result, nil
}

// traverse through non-visited points, enqueues, and marks them as visited
func (b *BFSMap) traverse(qp QueueMapPoint) {
	dist := qp.dist + 1
	life := qp.life

	//go up
	c := MapPoint{x: qp.p.x, y: qp.p.y - 1}
	if c.y >= 0 && b.maze[c.y][c.x] != 1 && !b.visited[c.y][c.x] {
		if b.maze[c.y][c.x] == consts.PitTrap {
			life -= 1
		}
		if b.maze[c.y][c.x] == consts.ArrowTrap {
			life -= 2
		}
		b.queue.PushBack(QueueMapPoint{p: c, dist: dist, life: life})
		b.visited[c.y][c.x] = true
	}
	//go down
	c = MapPoint{x: qp.p.x, y: qp.p.y + 1}
	if c.y < b.rows && b.maze[c.y][c.x] != 1 && !b.visited[c.y][c.x] {
		if b.maze[c.y][c.x] == consts.PitTrap {
			life -= 1
		}
		if b.maze[c.y][c.x] == consts.ArrowTrap {
			life -= 2
		}
		b.queue.PushBack(QueueMapPoint{p: c, dist: dist, life: life})
		b.visited[c.y][c.x] = true
	}

	//go left
	c = MapPoint{x: qp.p.x - 1, y: qp.p.y}
	if c.x >= 0 && b.maze[c.y][c.x] != 1 && !b.visited[c.y][c.x] {
		if b.maze[c.y][c.x] == consts.PitTrap {
			life -= 1
		}
		if b.maze[c.y][c.x] == consts.ArrowTrap {
			life -= 2
		}
		b.queue.PushBack(QueueMapPoint{p: c, dist: dist, life: life})
		b.visited[c.y][c.x] = true
	}
	//go right
	c = MapPoint{x: qp.p.x + 1, y: qp.p.y}
	if c.x < b.cols && b.maze[c.y][c.x] != 1 && !b.visited[c.y][c.x] {
		if b.maze[c.y][c.x] == consts.PitTrap {
			life -= 1
		}
		if b.maze[c.y][c.x] == consts.ArrowTrap {
			life -= 2
		}
		b.queue.PushBack(QueueMapPoint{p: c, dist: dist, life: life})
		b.visited[c.y][c.x] = true
	}
}

// GetSurvivablePath does a search through the maze and returns its distance and player's life
func (b *BFSMap) GetSurvivablePath() (int32, int32) {
	for b.queue.Len() != 0 {
		var qp QueueMapPoint

		//returns the first element
		qpo := b.queue.Front()

		if qpo != nil {
			qp = qpo.Value.(QueueMapPoint)
			b.queue.Remove(qpo)
		}

		//player died :(
		if qp.life <= 0 {
			return qp.dist, qp.life
		}

		//player found the exit
		if qp.p.x == b.end.x && qp.p.y == b.end.y {
			return qp.dist, qp.life
		}

		b.traverse(qp)
	}
	return -1, -1
}
