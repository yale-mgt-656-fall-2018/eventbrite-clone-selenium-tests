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

func randomTask(setCollaborators bool) Task {
	t := Task{
		name:        randomdata.Adjective() + " " + randomdata.Noun(),
		description: randomdata.Paragraph(),
	}
	if setCollaborators {
		t.collaborator1 = randomdata.Email()
		t.collaborator2 = randomdata.Email()
		t.collaborator3 = randomdata.Email()
	}
	return t
}
func (t Task) createFormData() map[string]string {
	data := map[string]string{
		selectors.TaskFormName:          t.name,
		selectors.TaskFormDescription:   t.description,
		selectors.TaskFormCollaborator1: t.collaborator1,
		selectors.TaskFormCollaborator2: t.collaborator2,
		selectors.TaskFormCollaborator3: t.collaborator3,
	}
	return data
}

func getBadTasks() []Task {
	var t Task
	var tasks []Task

	t = randomTask(true)
	t.flaw = "no name"
	t.name = ""
	tasks = append(tasks, t)

	t = randomTask(true)
	t.name = randomString(501)
	t.flaw = "name is more than 500 letters"
	tasks = append(tasks, t)

	t = randomTask(true)
	t.description = randomString(5001)
	t.flaw = "description is more than 5000"
	tasks = append(tasks, t)

	t = randomTask(true)
	t.collaborator1 = randomString(10)
	t.flaw = "invalid email for collaborator1"
	tasks = append(tasks, t)

	t = randomTask(true)
	t.collaborator2 = randomString(10)
	t.flaw = "invalid email for collaborator2"
	tasks = append(tasks, t)

	t = randomTask(true)
	t.collaborator3 = randomString(10)
	t.flaw = "invalid email for collaborator3"
	tasks = append(tasks, t)

	return tasks
}
