package main

type CaseRepository interface {
	Create(
		question string,
		testimonies []string,
		verdict string,
		confidence string,
	) Case

	Get(id int) (Case, bool)

	List() []Case
}
