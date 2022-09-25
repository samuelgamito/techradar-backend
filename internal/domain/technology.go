package domain

import "time"

type (
	TechnologyHistoryDomain struct {
		Score         int       `json:"score" cql:"score"`
		PreviousScore int       `json:"previous_score" cql:"previous_score"`
		AddedDate     time.Time `json:"added_date" cql:"added_date"`
		Comments      string    `json:"comments" cql:"comments"`
	}
	TechnologyDomain struct {
		Team          string                    `json:"team" cql:"team" `
		Title         string                    `json:"title" cql:"title"`
		FriendlyTitle string                    `json:"friendly_title" cql:"friendly_title"`
		Description   string                    `json:"description" cql:"description"`
		Moved         int                       `json:"moved" cql:"moved"`
		Score         int                       `json:"score" cql:"score"`
		Quadrant      int                       `json:"quadrant" cql:"quadrant"`
		Active        bool                      `json:"active" cql:"active"`
		CreatedAt     time.Time                 `json:"created_at" cql:"created_at"`
		UpdatedAt     time.Time                 `json:"updated_at" cql:"updated_at"`
		History       []TechnologyHistoryDomain `json:"history" cql:"history"`
	}
)
