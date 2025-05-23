package main

import (
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (player *Player) calculateRating() {
	if player.Misses == 0 {
		player.Rating = float64(player.Goals) + float64(player.Assists)/2.0
	} else {
		player.Rating = (float64(player.Goals) + float64(player.Assists)/2.0) / float64(player.Misses)
	}
}
func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{Name: name, Goals: goals, Misses: misses, Assists: assists, Rating: 0.0}
	p.calculateRating()
	return p
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if b.Goals != a.Goals {
			return b.Goals - a.Goals
		} else {
			if a.Name > b.Name {
				return 1
			}
			if a.Name < b.Name {
				return -1
			}
			return 0
		}
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating > b.Rating {
			return -1
		} else if a.Rating < b.Rating {
			return 1
		}
		if a.Name > b.Name {
			return 1
		}
		if a.Name < b.Name {
			return -1
		}
		return 0
	})
	return players
}
func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		x := float64(a.Goals)/float64(a.Misses) - float64(b.Goals)/float64(b.Misses)
		if x < 0 {
			return 1
		}
		if x > 0 {
			return -1
		}
		return 0
	})
	return players
}

//func main() {
//	players := []Player{
//		NewPlayer("Player1", 10, 5, 3),
//		NewPlayer("Player2", 15, 7, 2),
//		NewPlayer("Player3", 8, 2, 5),
//	}
//	x := ratingSort(players)
//	fmt.Println(x)
//}
