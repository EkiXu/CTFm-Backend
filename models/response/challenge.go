package response

import "ctfm_backend/models"

type ChallengesResponse struct {
	Challenges []models.Challenge `json:"challenges"`
}
