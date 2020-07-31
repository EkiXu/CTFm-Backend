package request

type AddChallengeStruct struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	IsHidden    bool   `json:"is_hidden"`
}
