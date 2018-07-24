# team
--
    import "github.com/mjnovice/teamfinder/team"


## Usage

#### type Color

```go
type Color struct {
	ShirtColorHome string `json:"shirtColorHome"`
	ShirtColorAway string `json:"shirtColorAway"`
	CrestMainColor string `json:"crestMainColor"`
	MainColor      string `json:"mainColor"`
}
```


#### type Competition

```go
type Competition struct {
	CompetitionID   int    `json:"competitionId"`
	CompetitionName string `json:"competitionName"`
}
```


#### type Data

```go
type Data struct {
	Team Team `json:"team"`
}
```


#### type LogoUrl

```go
type LogoUrl struct {
	Size string `json:"size"`
	URL  string `json:"url"`
}
```


#### type Official

```go
type Official struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
	Position  string `json:"position"`
}
```


#### type Player

```go
type Player struct {
	ID           string `json:"id"`
	Country      string `json:"country"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	Number       int    `json:"number"`
	BirthDate    string `json:"birthDate"`
	Age          string `json:"age"`
	Height       int    `json:"height"`
	Weight       int    `json:"weight"`
	ThumbnailSrc string `json:"thumbnailSrc"`
	PlayingFor   string `json:"-"`
}
```


#### type PlayerSorter

```go
type PlayerSorter []Player
```


#### func (PlayerSorter) Len

```go
func (a PlayerSorter) Len() int
```

#### func (PlayerSorter) Less

```go
func (a PlayerSorter) Less(i, j int) bool
```

#### func (PlayerSorter) Swap

```go
func (a PlayerSorter) Swap(i, j int)
```

#### type Team

```go
type Team struct {
	ID           int           `json:"id"`
	OptaID       int           `json:"optaId"`
	Name         string        `json:"name"`
	LogoUrls     []LogoUrl     `json:"logoUrls"`
	IsNational   bool          `json:"isNational"`
	Competitions []Competition `json:"competitions"`
	Players      []Player      `json:"players"`
	Officials    []Official    `json:"officials"`
	Colors       Color         `json:"colors"`
}
```


#### type TeamDataResponse

```go
type TeamDataResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}
```
