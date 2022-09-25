package dto

import (
	"github.com/go-playground/validator/v10"
	"techradar-backend/internal/domain"
	"techradar-backend/internal/misc"
	"time"
)

type (
	TechnologyHistoryDTO struct {
		Score         int       `json:"score"`
		PreviousScore int       `json:"previous_score"`
		AddedDate     time.Time `json:"added_date"`
		Comments      string    `json:"comments"`
	}
	TechnologyDTO struct {
		ID          string                 `json:"id"`
		Team        string                 `json:"Team"`
		Title       string                 `json:"title"`
		Description string                 `json:"description"`
		Moved       int                    `json:"moved"`
		Score       int                    `json:"score"`
		Quadrant    int                    `json:"quadrant"`
		Active      bool                   `json:"active"`
		CreatedAt   time.Time              `json:"created_at"`
		UpdatedAt   time.Time              `json:"updated_at"`
		History     []TechnologyHistoryDTO `json:"history"`
	}
	CreateTechnologyDTO struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Team        string `json:"team" validate:"required"`
		Score       *int   `json:"score" validate:"min=0,max=3"`
		Quadrant    *int   `json:"quadrant" validate:"required,min=0"`
	}
	MoveTechnologyDTO struct {
		Score    int    `json:"score"`
		Comments string `json:"comments"`
	}
	UpdateTechnologyDTO struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Quadrant    int    `json:"quadrant"`
		Active      bool   `json:"active"`
	}
)

func (m *ErrorResponse) ErrorMessage() *ErrorResponse {
	return m
}

func (cr *CreateTechnologyDTO) ToDomain() *domain.TechnologyDomain {
	friendlyName := misc.GetFriendlyName(cr.Title)
	now := time.Now()
	tech := domain.TechnologyDomain{
		Title:         cr.Title,
		FriendlyTitle: friendlyName,
		Description:   cr.Description,
		Team:          cr.Team,
		Quadrant:      *cr.Quadrant,
		Score:         *cr.Score,
		Moved:         0,
		Active:        true,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	return &tech
}

func (cr *CreateTechnologyDTO) IsValid() *ErrorResponse {
	v := validator.New()
	err := v.Struct(cr)

	if err == nil {
		return nil
	}

	errBody := ErrorResponse{
		StatusCode: 400,
		Body: ErrorBodyDTO{
			Messages: make([]string, len(err.(validator.ValidationErrors))),
		},
	}

	for i, e := range err.(validator.ValidationErrors) {
		errBody.Body.Messages[i] = buildErrorMessage(e)
	}

	return &errBody
}
