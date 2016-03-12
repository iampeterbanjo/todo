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

package task

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTask(t *testing.T) {
	Convey("Given a new task", t, func() {
		title := "Write tests"
		task, err := NewTask(title)

		Convey("It should return a task", func() {
			So(err, ShouldBeNil)
			So(task, ShouldHaveSameTypeAs, &Task{})
		})

		Convey("It should have that title", func() {
			So(task.Title, ShouldEqual, title)
		})

		Convey("It should not be done", func() {
			So(task.Done, ShouldBeFalse)
		})

		Convey("Cloning a task", func() {
			c := CloneTask(task)

			Convey("Should copy the same task", func() {
				So(c, ShouldResemble, task)
			})
		})
	})

	SkipConvey("Given a task with an empty title", t, func() {
		_, err := NewTask("")

		SkipConvey("It should error", func() {
			So(err, ShouldEqual, "missing title")
		})
	})
}

func TestManager(t *testing.T) {
	SkipConvey("Given a new task", t, func() {
		task, _ := newTaskOrFatal("learn Go")

		m := NewManager()

		SkipConvey("A saved task", func() {
			m.Save(task)
			all := m.All()

			SkipConvey("Should match created task", func() {
				So(len(all), ShouldEqual, 1)
				So(all[0], ShouldEqual, *task)
			})

			SkipConvey("Completing the task", func() {
				task.Done = true
				m.Save(task)

				SkipConvey("Should mark the saved task as complete", func() {
					So(m.All()[0].Done, ShouldBeTrue)
				})
			})
		})

		SkipConvey("Multiple saves of the task", func() {
			m.Save(task)
			m.Save(task)

			SkipConvey("Should be ok", func() {
				all := m.All()
				So(len(all), ShouldEqual, 1)
				So(*all[0], ShouldEqual, *task)
			})
		})

		SkipConvey("Finding a task", func() {
			f := m.Find(task.ID)

			SkipConvey("Should return the task", func() {
				So(f, ShouldEqual, *task)
			})
		})
	})

	SkipConvey("Given two new tasks", t, func() {
		learnGo, _ := newTaskOrFatal("learn Go")
		learnTDD, _ := newTaskOrFatal("learn TDD")

		m := NewManager()

		SkipConvey("Saving both of them", func() {
			m.Save(learnGo)
			m.Save(learnTDD)
			all := m.All()

			SkipConvey("They both should be saved", func() {
				So(len(all), ShouldEqual, 2)
			})

			SkipConvey("Saved tasks should match created tasks", func() {
				So(*all[0], ShouldEqual, *learnGo)
				So(*all[1], ShouldEqual, *learnTDD)
			})
		})
	})
}

func newTaskOrFatal(title string) (*Task, error) {
	task, err := NewTask(title)
	return task, err
}
