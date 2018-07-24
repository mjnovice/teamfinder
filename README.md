## How to Run ? 
```go run main.go```

# finder
--
    import "github.com/mjnovice/teamfinder/finder"


## Usage

```go
const (
	URLFormat = "https://vintagemonster.onefootball.com/api/teams/en/%d.json"
)
```

#### func  FindPlayersWithinBounds

```go
func FindPlayersWithinBounds(idBoundLow, idBoundHigh uint64, teams map[string]bool) []team.Player
```
FindPlayersWithinBounds returns all the players given two team ids.

#### func  FindUpperBound

```go
func FindUpperBound() uint64
```
FindUpperBound fetches the valid upper bound for the team_id in the URL for the
api call.

#### func  GetData

```go
func GetData(teamid uint64) (*team.TeamDataResponse, error)
```
GetData fetches data from the onefootball API endpoint by making an HTTP call
and giving back the response received.
