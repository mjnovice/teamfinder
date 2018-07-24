package team

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

type PlayerSorter []Player

func (a PlayerSorter) Len() int           { return len(a) }
func (a PlayerSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PlayerSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

type Color struct {
	ShirtColorHome string `json:"shirtColorHome"`
	ShirtColorAway string `json:"shirtColorAway"`
	CrestMainColor string `json:"crestMainColor"`
	MainColor      string `json:"mainColor"`
}

type LogoUrl struct {
	Size string `json:"size"`
	URL  string `json:"url"`
}

type Competition struct {
	CompetitionID   int    `json:"competitionId"`
	CompetitionName string `json:"competitionName"`
}

type Official struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
	Position  string `json:"position"`
}

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
type Data struct {
	Team Team `json:"team"`
}
type TeamDataResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}
