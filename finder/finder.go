package finder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"log"

	"github.com/mjnovice/teamfinder/team"
)

const (
	URLFormat = "https://vintagemonster.onefootball.com/api/teams/en/%d.json"
)

/*
GetData fetches data from the onefootball API endpoint by making an HTTP call
and giving back the response received.
*/
func GetData(teamid uint64) (*team.TeamDataResponse, error) {
	url := fmt.Sprintf(URLFormat, teamid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	teamData := &team.TeamDataResponse{}
	err = json.Unmarshal(data, teamData)
	if err != nil {
		return nil, err
	}
	return teamData, nil
}

/*
FindUpperBound fetches the valid upper bound for the team_id
in the URL for the api call.
*/
func FindUpperBound() uint64 {
	var l, r, mid uint64
	l = 1
	r = uint64(18446744073709551615)
	for {
		if l == r {
			return l - 1
		}
		mid = l + (r-l)/2
		teamData, err := GetData(mid)
		if err != nil || teamData.Code != 0 {
			r = mid
		} else {
			l = mid + 1
		}
		log.Println("Finding upper bound in the range ", l, r)
	}
	return 0
}

/*
	FindPlayersWithinBounds returns all the players given two team ids.
*/
func FindPlayersWithinBounds(idBoundLow, idBoundHigh uint64, teams map[string]bool) []team.Player {
	players := []team.Player{}
	var i uint64
	for i = idBoundLow; i <= idBoundHigh; i++ {
		log.Println("processing players for team id: ", i)
		teamData, err := GetData(i)
		if err != nil {
			return nil
		}
		teamName := teamData.Data.Team.Name
		if _, ok := teams[teamName]; ok {
			for _, player := range teamData.Data.Team.Players {
				player.PlayingFor = teamName
				players = append(players, player)
			}
		}
	}
	return players
}
