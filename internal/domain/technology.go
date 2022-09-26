package domain

import "time"

type (
	TechnologyHistoryDomain struct {
		Score         int       `cql:"score"`
		PreviousScore int       `cql:"previous_score"`
		AddedDate     time.Time `cql:"added_date"`
		Comments      string    `cql:"comments"`
	}
	TechnologyDomain struct {
		Team          string                    `cql:"team"`
		Title         string                    `cql:"title"`
		FriendlyTitle string                    `cql:"friendly_title"`
		Description   string                    `cql:"description"`
		Moved         int                       `cql:"moved"`
		Score         int                       `cql:"score"`
		Quadrant      int                       `cql:"quadrant"`
		Active        bool                      `cql:"active"`
		CreatedAt     time.Time                 `cql:"created_at"`
		UpdatedAt     time.Time                 `cql:"updated_at"`
		History       []TechnologyHistoryDomain `cql:"history"`
	}
	UpdateTechnologyDomain struct {
		Description *string `cql:"description"`
		Quadrant    *int    `cql:"quadrant" validate:"min=0"`
		Active      *bool   `cql:"active"`
	}
	MoveTechnologyDomain struct {
		Score    *int    `cql:"score"`
		Comments *string `cql:"comments"`
	}
)
