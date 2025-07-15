package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}
type LabGuard struct {
	pos position
	dir rune
}

const (
	UP       = '^'
	DOWN     = 'v'
	LEFT     = '<'
	RIGHT    = '>'
	OUTOFMAP = '$'
	OBSTACLE = '#'
)

func main() {
	//bytes_raw, err := (os.ReadFile("./example.txt"))
	bytes_raw, err := (os.ReadFile("./input.txt"))
	if err != nil {
		fmt.Println("ERROR: could not read file because of: ", err)
		os.Exit(1)

	}
	map_data := strings.Split(string(bytes_raw[:len(bytes_raw)-1]), "\n")
	var g LabGuard
	nwp := new([][]rune)
	for _, line := range map_data {
		tmp := []rune(line)
		*nwp = append(*nwp, tmp)
	}
	GetGuardPos(nwp, &g)
	Part1(nwp, &g)

}

func Part1(nwp *[][]rune, g *LabGuard) {
	fmt.Println("Part1:")
	fmt.Println("----------------------------------")
	fmt.Printf("the Guard visited %d distinct position\n", GuardPath(nwp, g))
	fmt.Println("----------------------------------")
}

func GuardPath(nwp *[][]rune, g *LabGuard) int {
	sum := 0
	for {
		Obj, ObjPos := GetWhatsInFrontofGuard(nwp, g)
		if Obj == OUTOFMAP {
			(*nwp)[g.pos.y][g.pos.x] = 'X'
			break
		}
		MoveGuard(nwp, g, Obj, ObjPos)

	}
	for _, line := range *nwp {
		for _, c := range line {
			if c == 'X' {
				sum++
			}
		}

	}
	return sum
}

func MoveGuard(nwp *[][]rune, g *LabGuard, Obj rune, posOfObj position) {
	pointer := *nwp
	if Obj == '.' || Obj == 'X' {
		pointer[g.pos.y][g.pos.x] = 'X'
		g.pos.x = posOfObj.x
		g.pos.y = posOfObj.y
		pointer[g.pos.y][g.pos.x] = g.dir
	}
	if Obj == '#' {
		pointer[g.pos.y][g.pos.x] = 'X'
		switch g.dir {
		case UP:
			g.pos.x = g.pos.x + 1
			g.dir = RIGHT
			pointer[g.pos.y][g.pos.x] = g.dir
		case DOWN:
			g.pos.x = g.pos.x - 1
			g.dir = LEFT
			pointer[g.pos.y][g.pos.x] = g.dir
		case LEFT:
			g.pos.y = g.pos.y - 1
			g.dir = UP
			pointer[g.pos.y][g.pos.x] = g.dir
		case RIGHT:
			g.pos.y = g.pos.y + 1
			g.dir = DOWN
			pointer[g.pos.y][g.pos.x] = g.dir

		}
	}

}
func GetGuardPos(data_map *[][]rune, lbguard *LabGuard) {
Outerloop:
	for i, line := range *data_map {
		for j, char := range line {
			switch char {
			case UP:
				lbguard.dir = UP
				lbguard.pos.x = j
				lbguard.pos.y = i
				break Outerloop
			case DOWN:
				lbguard.dir = DOWN
				lbguard.pos.x = j
				lbguard.pos.y = i
				break Outerloop
			case LEFT:
				lbguard.dir = LEFT
				lbguard.pos.x = j
				lbguard.pos.y = i
				break Outerloop
			case RIGHT:
				lbguard.dir = RIGHT
				lbguard.pos.x = j
				lbguard.pos.y = i
				break Outerloop

			}
		}
	}
}

func GetWhatsInFrontofGuard(map_data *[][]rune, lbguard *LabGuard) (rune, position) {
	var object rune
	var posOfObject position
	ly := len((*map_data))
	lx := len((*map_data)[0])
	if (lbguard.pos.x+1 >= lx && lbguard.dir == RIGHT) || (lbguard.pos.x-1 < 0 && lbguard.dir == LEFT) {
		object = OUTOFMAP
		posOfObject.x = -1
		posOfObject.y = -1
		return object, posOfObject
	}
	if (lbguard.pos.y+1 >= ly && lbguard.dir == DOWN) || (lbguard.pos.y-1 < 0 && lbguard.dir == UP) {
		object = OUTOFMAP
		posOfObject.x = -1
		posOfObject.y = -1
		return object, posOfObject
	}
	switch lbguard.dir {
	case UP:
		posOfObject.x = lbguard.pos.x
		posOfObject.y = lbguard.pos.y - 1
		object = (*map_data)[posOfObject.y][posOfObject.x]
	case DOWN:
		posOfObject.x = lbguard.pos.x
		posOfObject.y = lbguard.pos.y + 1
		object = (*map_data)[posOfObject.y][posOfObject.x]
	case LEFT:
		posOfObject.x = lbguard.pos.x - 1
		posOfObject.y = lbguard.pos.y
		object = (*map_data)[posOfObject.y][posOfObject.x]
	case RIGHT:
		posOfObject.x = lbguard.pos.x + 1
		posOfObject.y = lbguard.pos.y
		object = (*map_data)[posOfObject.y][posOfObject.x]
	}
	return object, posOfObject
}
