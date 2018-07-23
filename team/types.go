package team

type TeamDataResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   struct {
		Team struct {
			ID       int    `json:"id"`
			OptaID   int    `json:"optaId"`
			Name     string `json:"name"`
			LogoUrls []struct {
				Size string `json:"size"`
				URL  string `json:"url"`
			} `json:"logoUrls"`
			IsNational   bool `json:"isNational"`
			Competitions []struct {
				CompetitionID   int    `json:"competitionId"`
				CompetitionName string `json:"competitionName"`
			} `json:"competitions"`
			Players   []interface{} `json:"players"`
			Officials []interface{} `json:"officials"`
			Colors    struct {
				ShirtColorHome string `json:"shirtColorHome"`
				ShirtColorAway string `json:"shirtColorAway"`
				CrestMainColor string `json:"crestMainColor"`
				MainColor      string `json:"mainColor"`
			} `json:"colors"`
		} `json:"team"`
	} `json:"data"`
	Message string `json:"message"`
}
