package misc

import "github.com/scylladb/gocqlx/v2/qb"

func BuildWhereCondition(filter map[string]interface{}) []qb.Cmp {

	var conditions []qb.Cmp

	for k := range filter {
		conditions = append(conditions, qb.Eq(k))
	}

	return conditions
}
