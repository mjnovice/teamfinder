sample utility to read team data from one football's api

use 'godoc cmd/github.com/mjnovice/teamfinder/finder' for documentation on the github.com/mjnovice/teamfinder/finder command

PACKAGE DOCUMENTATION

package finder
    import "github.com/mjnovice/teamfinder/finder"


CONSTANTS

const (
    URLFormat = "https://vintagemonster.onefootball.com/api/teams/en/%d.json"
)

FUNCTIONS

func FindPlayersWithinBounds(idBoundLow, idBoundHigh uint64, teams map[string]bool) []team.Player
    FindPlayersWithinBounds returns all the players given two team ids.

func FindUpperBound() uint64
    FindUpperBound fetches the valid upper bound for the team_id in the URL
    for the api call.

func GetData(teamid uint64) (*team.TeamDataResponse, error)
    GetData fetches data from the onefootball API endpoint by making an HTTP
    call and giving back the response received.



