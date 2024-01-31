package repository

import (
	"app/entity"
	infra_jira "app/infrastructure/jira"
	"app/pkg/utils"
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/andygrunwald/go-jira"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const LAYOUT_DATA = "2006-01-02T15:04:05.999-0700"

type Moviment struct {
	Date   string `json:"date"`
	Status string `json:"status"`
}

// type ColumnsCalculated struct {
// 	Column string `json:"column"`
// 	Hours  int    `json:"hours"`
// }

// type ListCalcColumns struct {
// 	Calcs []ColumnsCalculated `json:"calcs"`
// }

// func (l *ListCalcColumns) GetCalcColumnIndex(column string) (int, bool) {
// 	for i, calc := range l.Calcs {
// 		if calc.Column == column {
// 			return i, true
// 		}
// 	}

// 	return 0, false
// }
// func (l *ListCalcColumns) GetCalcColumn(column string) (int, bool) {
// 	for _, calc := range l.Calcs {
// 		if calc.Column == column {
// 			return calc.Hours, true
// 		}
// 	}

// 	return 0, false
// }

// func (l *ListCalcColumns) AddCalcColumn(column string, hours int) {
// 	if index, ok := l.GetCalcColumnIndex(column); ok {
// 		l.Calcs[index].Hours += hours
// 	} else {
// 		l.Calcs = append(l.Calcs, ColumnsCalculated{
// 			Column: column,
// 			Hours:  hours,
// 		})
// 	}
// }

type RepositoryKanbanJira struct {
	JiraClient *jira.Client
	DB         *mongo.Database
}

func NewKanbanJira(jiraClient *jira.Client, mongoClient *mongo.Database) *RepositoryKanbanJira {
	return &RepositoryKanbanJira{JiraClient: jiraClient, DB: mongoClient}
}

func (r *RepositoryKanbanJira) GetRemoteIssues() (entities []entity.EntityKanban, err error) {

	options := jira.SearchOptions{
		Expand: "changelog",
	}

	issues, _, err := r.JiraClient.Issue.Search("statusCategoryChangedDate >= -7d order by created desc", &options)

	if err != nil {
		println(err)

		return nil, err
	}

	for _, issue := range issues {

		entityAux := entity.EntityKanban{
			ID:         issue.ID,
			ExternalID: issue.Key,
			Title:      issue.Fields.Summary,
			Status:     issue.Fields.Status.Name,
			Driver:     "jira",
		}

		if issue.Fields.Assignee != nil {
			entityAux.AssignUserName = issue.Fields.Assignee.DisplayName
			entityAux.AssignUserEmail = issue.Fields.Assignee.EmailAddress
		}
		jsonData, _ := json.Marshal(issue)

		entityAux.FullData = string(jsonData)

		entities = append(entities, entityAux)

	}

	return entities, nil
}

func (r *RepositoryKanbanJira) GetByID(id string) (issue *entity.EntityKanban, err error) {

	filter := bson.D{{Key: "id", Value: id}}

	r.DB.Collection("kanban").FindOne(context.Background(), filter).Decode(&issue)

	return issue, err
}

func (r *RepositoryKanbanJira) GetByExternalID(externalID string) (issue *entity.EntityKanban, err error) {

	filter := bson.D{{Key: "external_id", Value: externalID}}

	r.DB.Collection("kanban").FindOne(context.Background(), filter).Decode(&issue)

	return issue, err
}

func (r *RepositoryKanbanJira) Create(issue *entity.EntityKanban) (err error) {
	newID, _ := primitive.NewObjectID().MarshalText()

	issue.ID = string(newID)
	issue.ColumnCalc = []entity.ColumnCalc{}

	r.CalcColumnsTime(issue)

	_, err = r.DB.Collection("kanban").InsertOne(context.Background(), issue)

	return err

}

func (r *RepositoryKanbanJira) Update(issue *entity.EntityKanban) (err error) {

	_, err = r.GetByExternalID(issue.ExternalID)

	if err != nil {
		return err
	}

	if issue.ColumnCalc == nil {
		issue.ColumnCalc = []entity.ColumnCalc{}
	}

	r.CalcColumnsTime(issue)

	filter := bson.D{{Key: "external_id", Value: issue.ExternalID}}
	update := bson.D{{Key: "$set", Value: issue}}

	_, err = r.DB.Collection("kanban").UpdateOne(context.Background(), filter, update)

	return err
}

func (r *RepositoryKanbanJira) Delete(issue *entity.EntityKanban) (err error) {

	_, err = r.GetByExternalID(issue.ExternalID)

	if err != nil {
		return err
	}

	filter := bson.D{{Key: "external_id", Value: issue.ExternalID}}

	_, err = r.DB.Collection("kanban").DeleteOne(context.Background(), filter)

	return err
}

func (r *RepositoryKanbanJira) CalcColumnsTime(issue *entity.EntityKanban) (err error) {

	issue.ColumnCalc = []entity.ColumnCalc{}

	movimentWorked, _ := r.GetMovimentsGrouped(issue)

	for i, movto := range movimentWorked {
		startDate, _ := time.Parse(LAYOUT_DATA, movto.Date)

		var endDate time.Time

		if i+1 < len(movimentWorked) {
			endDate, _ = time.Parse(LAYOUT_DATA, movimentWorked[i+1].Date)
		}

		statusKey := strings.ToLower(movto.Status)
		respTime := utils.GetHoursInTimes(startDate, endDate)

		issue.AddCalcColumn(entity.ColumnCalc{
			ColumName: statusKey,
			// Count:     counter,
			Hours:   respTime.ToHours(),
			Minutes: respTime.ToMinutes(),
			Seconds: respTime.ToSeconds(),
		})

	}

	return nil

}

func (r *RepositoryKanbanJira) GetMovimentsGrouped(issue *entity.EntityKanban) (moviments []Moviment, err error) {

	var jiraLog infra_jira.JiraChangeLog

	err = json.Unmarshal([]byte(issue.FullData), &jiraLog)

	if err != nil {
		return nil, err
	}

	for _, history := range jiraLog.Changelog.Histories {
		if len(history.Items) == 0 {
			continue
		}

		firstItem := history.Items[0]

		if firstItem.Field != "status" {
			continue
		}

		moviments = append(moviments, Moviment{
			Date:   history.Created,
			Status: firstItem.ToString,
		})
	}

	return moviments, nil
}
