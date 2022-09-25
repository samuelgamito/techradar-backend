package repository

import (
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"sort"
	"techradar-backend/internal/config"
	"techradar-backend/internal/domain"
)

type InfoRepository struct {
	dbSession *gocqlx.Session
}

func NewInfoRepository(dbSession *gocqlx.Session) *InfoRepository {
	return &InfoRepository{dbSession: dbSession}
}

func (i *InfoRepository) GetDistinctTeam() ([]string, error) {

	var output []string
	stmt, names := qb.Select(config.TechnologyTableName).Distinct("team").ToCql()
	q := i.dbSession.Query(stmt, names)
	err := q.SelectRelease(&output)

	return output, err
}

func (i *InfoRepository) GetQuadrantsInfo() ([]domain.Quadrant, error) {
	var output []domain.Quadrant
	stmt, names := qb.Select(config.QuadrantTableName).ToCql()
	q := i.dbSession.Query(stmt, names)

	err := q.SelectRelease(&output)

	sort.SliceStable(output, func(i, j int) bool {
		return output[i].Title < output[j].Title
	})

	return output, err
}

func (i *InfoRepository) GetQuadrantInfoById(id int) (*domain.Quadrant, error) {
	var outputArr []domain.Quadrant
	q := i.dbSession.Query(config.QuadrantTable.Get()).BindMap(qb.M{"id": id})

	err := q.SelectRelease(&outputArr)

	if outputArr != nil {
		return &outputArr[0], nil
	}
	return nil, err
}

func (i *InfoRepository) UpsertQuadrantsInfo(quadrant domain.Quadrant) error {

	q := i.dbSession.Query(config.QuadrantTable.Insert()).BindStruct(quadrant)
	return q.ExecRelease()
}
