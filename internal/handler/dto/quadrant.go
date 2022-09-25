package dto

import "backend/internal/domain"

type Quadrant struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (q *Quadrant) ToDomain() domain.Quadrant {

	return domain.Quadrant{
		Id:    q.Id,
		Title: q.Title,
	}
}
