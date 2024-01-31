package usecase_kanban

import "app/entity"

type IRepositoryKanban interface {
	GetRemoteIssues() (issues []entity.EntityKanban, err error)
	GetByID(id string) (issue *entity.EntityKanban, err error)
	GetByExternalID(externalID string) (issue *entity.EntityKanban, err error)
	Create(issue *entity.EntityKanban) (err error)
	Update(issue *entity.EntityKanban) (err error)
	Delete(issue *entity.EntityKanban) (err error)
	CalcColumnsTime(issue *entity.EntityKanban) (err error)
}

type IUseCaseKanban interface {
	UpdateIssues() (err error)
}
