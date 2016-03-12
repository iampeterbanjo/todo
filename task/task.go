// Copyright 2014 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

// Package task provides a Task and Manager implemented using TDD techniques.
// The tests were developed before the code was written.
package task

import (
	"errors"
	"math/rand"
)

// Task has a title
type Task struct {
	Title string
	Done  bool
	ID    float64
}

// Manager has methods to work with tasks
type Manager struct {
	Tasks []*Task
}

// NewTask creates Tasks given a title
func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, errors.New("missing title")
	}

	return &Task{title, false, rand.Float64()}, nil
}

// NewManager creates a Manager
func NewManager() *Manager {
	return &Manager{}
}

// Save task to Manager.Tasks
func (m *Manager) Save(task *Task) {
	all := m.All()
	found := false

	for i, f := range all {
		if f.ID == task.ID {
			m.Tasks[i] = CloneTask(task)
			found = true
			return
		}
	}

	if !found {
		m.Tasks = append(m.Tasks, CloneTask(task))
		return
	}
}

// CloneTask to create a copy
func CloneTask(t *Task) *Task {
	c := *t
	return &c
}

// All tasks
func (m Manager) All() []*Task {
	return m.Tasks
}

// Find a Task in Tasks
func (m Manager) Find(id float64) *Task {
	f := new(Task)

	for _, t := range m.All() {
		if t.ID == id {
			f = t
		}
	}

	return f
}
