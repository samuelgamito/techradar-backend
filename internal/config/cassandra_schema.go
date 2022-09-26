package config

import "github.com/scylladb/gocqlx/v2/table"

var technologyMetadata = table.Metadata{
	Name: "technology",
	Columns: []string{
		"title",
		"friendly_title",
		"description",
		"team",
		"moved",
		"score",
		"quadrant",
		"active",
		"history",
		"created_at",
		"updated_at",
	},
	PartKey: []string{"team", "friendly_title"},
}

var quadrantMetadata = table.Metadata{
	Name: "quadrant",
	Columns: []string{
		"id",
		"title",
	},
	PartKey: []string{"id"},
}

var TechnologyTableName = "tech_radar.technology"
var TechnologyTable = table.New(technologyMetadata)

var QuadrantTableName = "tech_radar.quadrant"
var QuadrantTable = table.New(quadrantMetadata)
