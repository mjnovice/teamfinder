package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/mjnovice/teamfinder/finder"
	"github.com/mjnovice/teamfinder/team"
)

func findPlayersParallely(batchSize, idBoundHigh uint64, teams map[string]bool) []team.Player {
	playersChan := make(chan team.Player, 1000000)
	wg := &sync.WaitGroup{}
	var i uint64
	for i = 1; i <= idBoundHigh; i += batchSize {
		wg.Add(1)
		go func(i uint64, wg *sync.WaitGroup) {
			defer wg.Done()
			players := finder.FindPlayersWithinBounds(i, i+batchSize-1, teams)
			for _, player := range players {
				playersChan <- player
			}
		}(i, wg)
	}
	wg.Wait()
	close(playersChan)
	players := []team.Player{}
	for player := range playersChan {
		players = append(players, player)
	}
	return players
}

func printPlayers(players []team.Player) {
	n := len(players)
	for i := 0; i < n; {
		playerName := players[i].Name
		out := fmt.Sprintf("%d. %s; %s; ", i, players[i].Name, players[i].Age)
		for j := i; j < n; j++ {
			if players[j].Name == playerName {
				out += players[j].PlayingFor + ", "
				i++
			} else {
				break
			}
		}
		out = strings.TrimSuffix(out, ", ")
		fmt.Println(out)
	}
}

func main() {
	teamNames := map[string]bool{
		"Germany":          true,
		"England":          true,
		"France":           true,
		"Spain":            true,
		"Manchester Utd":   true,
		"Arsenal":          true,
		"Chelsea":          true,
		"Barcelona":        true,
		"Real Madrid":      true,
		"FC Bayern Munich": true,
	}

	var w io.Writer

	//dummy output buffer to suppress debug logs.
	_, w, _ = os.Pipe()

	//w = os.Stdout
	log.SetOutput(w)

	limit := finder.FindUpperBound()
	log.Println("Limit foud! ", limit)

	players := findPlayersParallely(300, limit, teamNames)
	log.Println("Total players found", len(players))

	sort.Sort(team.PlayerSorter(players))
	printPlayers(players)
}
