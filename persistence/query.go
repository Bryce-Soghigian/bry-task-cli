package persistence

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func CreateNewTask(_title string, _description string, db *gorm.DB) {
	/*
		Create a new task in the database.
	*/
	db, err := gorm.Open(postgres.Open(DbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	currTime := time.Now().Unix()
	newTask := Task{
		status:         "CREATED",
		category:       "Primary",
		title:          _title,
		task_type:      "rent",
		description:    _description,
		creation_date:  currTime,
		completed_date: 0,
	}
	db.Create(&newTask)
	fmt.Println(newTask)

}

func GetAllCompletedTasks(db *gorm.DB) []Task {
	/*
		Get all tasks of completed status.
	*/
	var task []Task
	db.Find(&task).Where("status = ?", "CREATED")
	return task
}
