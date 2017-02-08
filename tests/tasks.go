package tests

import (
	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/yale-cpsc-213/social-todo-selenium-tests/tests/selectors"
)

// Task ...
//
type Task struct {
	name          string
	description   string
	collaborator1 string
	collaborator2 string
	collaborator3 string
	flaw          string
}

func randomTask() Task {
	return Task{
		name:        randomdata.Adjective() + " " + randomdata.Noun(),
		description: randomdata.Paragraph(),
	}
}
func (t Task) createFormData() map[string]string {
	data := map[string]string{
		selectors.TaskFormName:          t.name,
		selectors.TaskFormDescription:   t.description,
		selectors.TaskFormCollaborator1: t.collaborator1,
		selectors.TaskFormCollaborator3: t.collaborator2,
	}
	return data
}
