package usecase_kanban

type UsecaseKanban struct {
	repo IRepositoryKanban
}

func NewService(repository IRepositoryKanban) *UsecaseKanban {
	return &UsecaseKanban{repo: repository}
}

func (u *UsecaseKanban) UpdateIssues() (err error) {

	issues, err := u.repo.GetRemoteIssues()

	if err != nil {
		return err
	}

	for _, issue := range issues {

		issueDB, err := u.repo.GetByExternalID(issue.ExternalID)

		if err != nil {
			return err
		}

		if issueDB != nil {
			issue.ID = issueDB.ID
			err = u.repo.Update(&issue)
		} else {
			err = u.repo.Create(&issue)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
