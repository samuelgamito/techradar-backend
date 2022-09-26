package misc

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"techradar-backend/internal/domain"
)

func BuildWhereCondition(filter map[string]interface{}) []qb.Cmp {

	var conditions []qb.Cmp

	for k := range filter {
		conditions = append(conditions, qb.Eq(k))
	}

	return conditions
}

func BuildUpdateObject(old *domain.TechnologyDomain, fields *domain.UpdateTechnologyDomain) {

	if fields.Active != nil {
		old.Active = *fields.Active
	}

	if fields.Quadrant != nil {
		old.Quadrant = *fields.Quadrant
	}

	if fields.Description != nil {
		old.Description = *fields.Description
	}
}

func BuildHistoryObject(techObj *domain.TechnologyDomain, move *domain.MoveTechnologyDomain) {

	newHistoryObj := domain.TechnologyHistoryDomain{}

	if move.Score == nil {
		newHistoryObj.Score = techObj.Score
		newHistoryObj.PreviousScore = techObj.Score
	} else {
		newHistoryObj.PreviousScore = techObj.Score
		newHistoryObj.Score = *move.Score
		techObj.Score = *move.Score
	}

	if move.Comments == nil {
		newHistoryObj.Comments = "N/A"
	} else {
		newHistoryObj.Comments = *move.Comments
	}

	techObj.History = append(techObj.History, newHistoryObj)
}
