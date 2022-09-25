package repository

import (
	"backend/internal/config"
	"backend/internal/domain"
	"backend/internal/misc"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type TechnologyRepository struct {
	dbSession *gocqlx.Session
}

func NewTechnologyRepository(dbSession *gocqlx.Session) *TechnologyRepository {
	return &TechnologyRepository{
		dbSession: dbSession,
	}
}

func (t *TechnologyRepository) GetTechnologyByTeamAndTitle(team string, title string) ([]domain.TechnologyDomain, error) {
	var filters = make(map[string]interface{})
	var result []domain.TechnologyDomain

	filters["team"] = team
	filters["friendly_title"] = title

	err := t.getById(filters, &result)

	return result, err
}

func (t *TechnologyRepository) GetTechnologyByTeamAndQuadrants(team string, quadrant int) ([]domain.TechnologyDomain, error) {
	var filters = make(map[string]interface{})
	var result []domain.TechnologyDomain

	filters["team"] = team
	filters["quadrant"] = quadrant

	err := t.findTechnology(filters, &result)

	return result, err
}

func (t *TechnologyRepository) GetTechnologyByTeam(team string) ([]domain.TechnologyDomain, error) {
	var filters = make(map[string]interface{})
	var result []domain.TechnologyDomain

	filters["team"] = team

	err := t.findTechnology(filters, &result)

	return result, err
}

func (t *TechnologyRepository) CreateTechnology(tech *domain.TechnologyDomain) error {
	q := t.dbSession.Query(config.TechnologyTable.Insert()).BindStruct(tech)

	return q.ExecRelease()
}

func (t *TechnologyRepository) getById(filters map[string]interface{}, output interface{}) error {
	q := t.dbSession.Query(config.TechnologyTable.Get()).BindMap(filters)

	return q.SelectRelease(output)
}

func (t *TechnologyRepository) findTechnology(filters map[string]interface{}, output interface{}) error {
	stmt, names := qb.Select(config.TechnologyTableName).AllowFiltering().Where(misc.BuildWhereCondition(filters)...).ToCql()
	q := t.dbSession.Query(stmt, names).BindMap(filters)

	return q.SelectRelease(output)
}
