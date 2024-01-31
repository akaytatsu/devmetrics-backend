package repository_test

import (
	"app/entity"
	"app/infrastructure/repository"
	"app/pkg/testing_utils"
	usecase_kanban "app/usecase/kanban"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcColumnsTime(t *testing.T) {
	t.Run("should return the time in seconds between two dates", func(t *testing.T) {

		content, _ := testing_utils.GetMockedFile("calc_columns_time_1.json")

		issue := entity.EntityKanban{
			ExternalID:     "TEST-1",
			Title:          "Test",
			AssignUserName: "Test",
			ColumnCalc:     []entity.ColumnCalc{},
			FullData:       string(content),
		}

		var r usecase_kanban.IRepositoryKanban = &repository.RepositoryKanbanJira{}

		r.CalcColumnsTime(&issue)

		println(len(issue.ColumnCalc))

		for _, calc := range issue.ColumnCalc {
			println(calc.ColumName)
			println(calc.Hours)
		}

		assert.Len(t, issue.ColumnCalc, 6)

	})
}

func TestGetMovimentsGrouped(t *testing.T) {
	r := &repository.RepositoryKanbanJira{}

	// Create a mock issue
	issue := &entity.EntityKanban{
		FullData: `{
            "changelog": {
                "histories": [
                    {
                        "created": "2022-12-25T00:00:00.000+0000",
                        "items": [
                            {
                                "field": "status",
                                "toString": "Done"
                            }
                        ]
                    }
                ]
            }
        }`,
	}

	moviments, err := r.GetMovimentsGrouped(issue)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(moviments) != 1 {
		t.Errorf("Expected 1 moviment, got %d", len(moviments))
	}

	if moviments[0].Status != "Done" {
		t.Errorf("Expected status to be 'Done', got '%s'", moviments[0].Status)
	}
}
