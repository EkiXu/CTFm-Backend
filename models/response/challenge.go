package response

type ChallengeResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Solved   int    `json:"solved"`
	Attempts int    `json:"attempts"`
	Points   int    `json:"points"`
	IsHidden bool   `json:"is_hidden"`
}

type ChallengeContentResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Points      int    `json:"points"`
}

type ChallengeDetailResponse struct {
	ChallengeContentResponse
	IsHidden bool   `json:"is_hidden"`
	Flag     string `json:"flag"`
}

type ChallengesListResponse struct {
	Challenges []ChallengeResponse `json:"challenges"`
}

type ChallengedAddedResponse struct {
	ID uint `json:"id"`
}
