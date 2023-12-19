package models

type question struct {
	Name string
	Type string
}

func NewQuestion(name string, typ string) question {
	question := question{
		Name: name,
		Type: typ,
	}

	return question
}
