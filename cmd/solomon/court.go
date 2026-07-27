package main

type Court struct {
	repo CaseRepository
}

func NewCourt(repo CaseRepository) *Court {
	return &Court{
		repo: repo,
	}
}

func (c *Court) CreateCase(
	question string,
	testimonies []string,
	verdict string,
	confidence string,
) Case {
	return c.repo.Create(
		question,
		testimonies,
		verdict,
		confidence,
	)
}

func (c *Court) GetCase(id int) (Case, bool) {
	return c.repo.Get(id)
}

func (c *Court) ListCases() []Case {
	return c.repo.List()
}
