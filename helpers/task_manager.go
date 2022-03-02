package helpers

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/Bryce-Soghigian/bry-task-cli/persistence"
)

func AddTask(command_args []string, db *gorm.DB) (response string, err error) {

	persistence.CreateNewTask(command_args[1], command_args[3], db)
	return "Successfully Added Task.", nil
}
func RemoveTask(command_args []string) (response string, err error) {
	fmt.Println(command_args)
	return "Successfully Removed Task.", nil
}
func ModifyTask(command_args []string) (response string, err error) {

	return "Successfully Modified Task.", nil
}
func ListTasks(db *gorm.DB, command_args []string) (response string, err error) {

	tasks := persistence.GetAllCompletedTasks(db)
	fmt.Println(tasks)
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
	}
	return "Successfully Listed Tasks.", nil
}
