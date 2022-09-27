package misc

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"reflect"
	"techradar-backend/internal/domain"
	"testing"
	"time"
)

type (
	buildHistoryObjectTestStruct struct {
		arg1, expected domain.TechnologyDomain
		arg2           domain.MoveTechnologyDomain
	}
	buildUpdateObjectTestStruct struct {
		arg1, expected domain.TechnologyDomain
		arg2           domain.UpdateTechnologyDomain
	}
)

var falseBool = false
var quadrantOne = 1
var scoreTwo = 2
var scoreZero = 0
var description = "test update"
var comments = "new comment to update"

var (
	buildUpdateObjectTestScenarios = []buildUpdateObjectTestStruct{
		{
			arg1: domain.TechnologyDomain{
				Active:      true,
				Quadrant:    0,
				Description: "old description",
			},
			arg2: domain.UpdateTechnologyDomain{
				Active:      &falseBool,
				Quadrant:    &quadrantOne,
				Description: &description,
			},
			expected: domain.TechnologyDomain{
				Active:      false,
				Quadrant:    1,
				Description: "test update",
			},
		},
	}
	buildHistoryObjectTestScenarios = []buildHistoryObjectTestStruct{
		{
			arg1: domain.TechnologyDomain{
				Score:   1,
				Moved:   0,
				History: nil,
			},
			arg2: domain.MoveTechnologyDomain{
				Score:    &scoreTwo,
				Comments: &comments,
			},
			expected: domain.TechnologyDomain{
				Score: 2,
				Moved: 1,
				History: []domain.TechnologyHistoryDomain{
					{
						Score:         2,
						PreviousScore: 1,
						Comments:      "new comment to update",
					},
				},
			},
		},
		{
			arg1: domain.TechnologyDomain{
				Score:   1,
				Moved:   0,
				History: nil,
			},
			arg2: domain.MoveTechnologyDomain{
				Score:    &scoreTwo,
				Comments: nil,
			},
			expected: domain.TechnologyDomain{
				Score: 2,
				Moved: 1,
				History: []domain.TechnologyHistoryDomain{
					{
						Score:         2,
						PreviousScore: 1,
						Comments:      "N/A",
					},
				},
			},
		},
		{
			arg1: domain.TechnologyDomain{
				Score:   1,
				Moved:   0,
				History: nil,
			},
			arg2: domain.MoveTechnologyDomain{
				Score:    nil,
				Comments: &comments,
			},
			expected: domain.TechnologyDomain{
				Score: 1,
				Moved: 0,
				History: []domain.TechnologyHistoryDomain{
					{
						Score:         1,
						PreviousScore: 1,
						Comments:      "new comment to update",
					},
				},
			},
		},
		{
			arg1: domain.TechnologyDomain{
				Score:   1,
				Moved:   0,
				History: nil,
			},
			arg2: domain.MoveTechnologyDomain{
				Score:    &scoreZero,
				Comments: &comments,
			},
			expected: domain.TechnologyDomain{
				Score: 0,
				Moved: -1,
				History: []domain.TechnologyHistoryDomain{
					{
						Score:         0,
						PreviousScore: 1,
						Comments:      "new comment to update",
					},
				},
			},
		},
		{
			arg1: domain.TechnologyDomain{
				Score:   0,
				Moved:   -1,
				History: nil,
			},
			arg2: domain.MoveTechnologyDomain{
				Score:    &scoreZero,
				Comments: &comments,
			},
			expected: domain.TechnologyDomain{
				Score: 0,
				Moved: 0,
				History: []domain.TechnologyHistoryDomain{
					{
						Score:         0,
						PreviousScore: 0,
						Comments:      "new comment to update",
					},
				},
			},
		},
	}
)

func TestBuildUpdateObject(t *testing.T) {
	for _, testScenario := range buildUpdateObjectTestScenarios {
		if BuildUpdateObject(&testScenario.arg1, &testScenario.arg2); !reflect.DeepEqual(testScenario.arg1, testScenario.expected) {
			t.Errorf("Output %+v not equal to expected %+v", testScenario.arg1, testScenario.expected)
		}
	}
}

func TestBuildHistoryObject(t *testing.T) {
	for _, testScenario := range buildHistoryObjectTestScenarios {

		timeBeforeExecute := time.Now()

		BuildHistoryObject(&testScenario.arg1, &testScenario.arg2)

		isValid := testScenario.arg1.Score == testScenario.expected.Score &&
			testScenario.arg1.Moved == testScenario.expected.Moved &&
			len(testScenario.arg1.History) == len(testScenario.expected.History) &&
			testScenario.arg1.History[0].Score == testScenario.expected.History[0].Score &&
			testScenario.arg1.History[0].PreviousScore == testScenario.expected.History[0].PreviousScore &&
			testScenario.arg1.History[0].Comments == testScenario.expected.History[0].Comments

		if !isValid {
			t.Errorf("Output %+v not equal to expected %+v", testScenario.arg1, testScenario.expected)
		}
	}
}

func TestBuildWhereCondition(t *testing.T) {
	var f = make(map[string]interface{})

	f["test"] = "test"
	f["test_2"] = "test_2"
	var resp []qb.Cmp
	resp = BuildWhereCondition(f)

	isValid := len(resp) == 2 &&
		resp[0] == qb.Eq("test") &&
		resp[1] == qb.Eq("test_2")

	if !isValid {
		t.Errorf("Where condition not valid")
	}
}
