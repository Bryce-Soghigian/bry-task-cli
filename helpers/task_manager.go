package helpers

import (
	"fmt"

	"github.com/Bryce-Soghigian/bry-task-cli/persistence"
)

func AddTask(command_args []string) (response string, err error) {
	persistence.CreateNewTask(command_args[1], command_args[3])
	return "Successfully Added Task.", nil
}
func RemoveTask(command_args []string) (response string, err error) {
	fmt.Println(command_args)
	return "Successfully Removed Task.", nil
}
func ModifyTask(command_args []string) (response string, err error) {

	return "Successfully Modified Task.", nil
}
func ListTasks(command_args []string) (response string, err error) {
	return "Successfully Listed Tasks.", nil
}
