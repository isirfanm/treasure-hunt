package main

import (
	"fmt"
	"os"
)

// Mapz type
type Mapz [][]rune

// tiles const
const (
	Obstacle = '#'
	Path     = '.'
	Player   = 'X'
	Treasure = '$'
)

// Loc position type
type Loc struct {
	Row int
	Col int
}

// FindPlayer find the player position
func FindPlayer(m Mapz) *Loc {
	for r, row := range m {
		for c, t := range row {
			if t == Player {
				// found player
				return &Loc{Row: r, Col: c}
			}
		}
	}

	// player not found
	return nil
}

// Direction type
type Direction func(Loc) Loc

// Direction enum
var (
	Up = func(l Loc) Loc {
		l.Row -= 1
		return l
	}
	Down = func(l Loc) Loc {
		l.Row += 1
		return l
	}
	Right = func(l Loc) Loc {
		l.Col += 1
		return l
	}
	Left = func(l Loc) Loc {
		l.Col -= 1
		return l
	}
)

// Hunter entity
type Hunter struct {
	Mapz         Mapz
	StartLoc     Loc
	TreasureLocs []Loc
}

// Search all the treasures, hunter! Go!
func (h *Hunter) Search(d []Direction) {
	nRow := len(h.Mapz)
	nCol := len(h.Mapz[0])

	curPos := []Loc{h.StartLoc}
	// follow each direction
	for _, df := range d {
		nextPos := []Loc{}

		// move each current position
		for _, p := range curPos {
			// move according to the direction
			for {
				// move it
				p = df(p)

				if p.Col < 0 || p.Col == nCol {
					// out of boundary
					break
				}

				if p.Row < 0 || p.Row == nRow {
					// out of boundary
					break
				}

				// get tile
				t := h.Mapz[p.Row][p.Col]
				if t == Obstacle {
					// hit obstacle
					break
				}

				// clear path, add to next position
				nextPos = append(nextPos, p)
			}

		}

		// set next pos to cur pos
		curPos = nextPos
	}

	// last current positions is the possible treasure positions
	h.TreasureLocs = curPos

	// mark treasure position to mapz
	for _, l := range h.TreasureLocs {
		h.Mapz[l.Row][l.Col] = Treasure
	}
}

func ShowLocations(ts []Loc) {
	for _, l := range ts {
		fmt.Printf("(%d, %d)\n", l.Col, l.Row)
	}
}

func ShowMapz(m Mapz) {
	for _, row := range m {
		fmt.Println(string(row))
	}
}

func main() {
	// the map
	mapz := Mapz{
		[]rune("########"),
		[]rune("#......#"),
		[]rune("#.###..#"),
		[]rune("#...#.##"),
		[]rune("#X#....#"),
		[]rune("########"),
	}
	ShowMapz(mapz)

	// find player location
	playerLoc := FindPlayer(mapz)
	if playerLoc == nil {
		fmt.Println("Player location is not found. Terminating...")
		os.Exit(1)
	}

	// make direction list
	directions := []Direction{
		Up,
		Right,
		Down,
	}

	// make Hunter
	hunter := Hunter{
		Mapz:     mapz,
		StartLoc: *playerLoc,
	}

	// hunter use the direction list to search the treasure
	hunter.Search(directions)

	// get treasure locations
	treasures := hunter.TreasureLocs
	ShowLocations(treasures)

	// get final mapz
	finalMapz := hunter.Mapz
	ShowMapz(finalMapz)
}
