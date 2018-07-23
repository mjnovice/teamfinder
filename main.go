package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	URLFormat = "https://vintagemonster.onefootball.com/api/teams/en/%d.json"
)

func GetData(teamid uint64) (*team.TeamDataResponse, error) {
	url = fmt.Sprintf(URLFormat, teamid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp)
	if err != nil {
		return nil, err
	}
	teamDataObj := &TeamDataResponse{}
	err := json.Unmarshal(teamData, teamDataObj)
	if err != nil {
		return nil, err
	}
	return teamDataObj, nil
}

func main() {
	GetData(12)
}
