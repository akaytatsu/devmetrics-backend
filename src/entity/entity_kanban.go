package entity

import "time"

type ColumnCalc struct {
	ColumName string `json:"column_name" bson:"column_name"`
	Count     int    `json:"count" bson:"count"`
	Hours     int    `json:"hours" bson:"hours"`
	Minutes   int    `json:"minutes" bson:"minutes"`
	Seconds   int    `json:"seconds" bson:"seconds"`
}

type EntityKanban struct {
	ID              string       `json:"_id" bson:"_id"`
	Driver          string       `json:"driver" bson:"driver"`
	ExternalID      string       `json:"external_id" bson:"external_id"`
	Title           string       `json:"title" bson:"title"`
	Status          string       `json:"status" bson:"status"`
	AssignUserName  string       `json:"assign_user_name" bson:"assign_user_name"`
	AssignUserEmail string       `json:"assign_user_email" bson:"assign_user_email"`
	UpdatedAt       time.Time    `json:"updated_at" bson:"updated_at"`
	CreatedAt       time.Time    `json:"created_at" bson:"created_at"`
	ColumnCalc      []ColumnCalc `json:"column_calc" bson:"column_calc"`
	FullData        string       `json:"full_data" bson:"full_data"`
}

// type ColumnsCalculated struct {
// 	Column string `json:"column"`
// 	Hours  int    `json:"hours"`
// }

// type ListCalcColumns struct {
// 	Calcs []ColumnsCalculated `json:"calcs"`
// }

func (l *EntityKanban) GetCalcColumnIndex(column string) (int, bool) {
	for i, calc := range l.ColumnCalc {
		if calc.ColumName == column {
			return i, true
		}
	}

	return 0, false
}
func (l *EntityKanban) GetCalcColumn(column string) (int, bool) {
	for _, calc := range l.ColumnCalc {
		if calc.ColumName == column {
			return calc.Hours, true
		}
	}

	return 0, false
}

func (l *EntityKanban) AddCalcColumn(column ColumnCalc) {
	if index, ok := l.GetCalcColumnIndex(column.ColumName); ok {
		l.ColumnCalc[index].Hours += column.Hours
		l.ColumnCalc[index].Minutes += column.Minutes
		l.ColumnCalc[index].Seconds += column.Seconds
	} else {
		l.ColumnCalc = append(l.ColumnCalc, column)
	}
}
