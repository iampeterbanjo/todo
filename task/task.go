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

// Package task provides a Task and TaskManager implemented using TDD techniques.
// The tests were developed before the code was written.
package task

// Task has a title
type Task struct {
	Title string
	Done  bool
}

// NewTask creates Tasks given a title
func NewTask(title string) (*Task, error) {
	return &Task{title, false}, nil
}
