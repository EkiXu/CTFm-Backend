package request

type EditChallengeStruct struct {
	Name        string `json:"name"`
	Flag        string `json:"flag"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	IsHidden    bool   `json:"is_hidden"`
}

type CheckFlagStruct struct {
	Flag string `json:"flag"`
}
