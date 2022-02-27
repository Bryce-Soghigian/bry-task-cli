package persistence

import (
	"fmt"
	"time"
)

func NewTask(_title string, _description string) {
	/*
		Create a new task in the database.
	*/
	curr_time := time.Now().Unix()
	new_task := &Task{
		status:         "CREATED",
		category:       "Primary",
		title:          _title,
		task_type:      "Parent",
		description:    _description,
		creation_date:  curr_time,
		completed_date: 0,
	}
	fmt.Println(new_task)

}
