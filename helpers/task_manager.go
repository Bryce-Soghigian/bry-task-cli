package helpers

import "errors"

func AddTask(command_args []string) (response string, err error) {
	return "Successfully Added Task.", errors.New("This is broken")
}
func RemoveTask(command_args []string) (response string, err error) {
	return "Successfully Removed Task.", errors.New("This is broken")
}
func ModifyTask(command_args []string) (response string, err error) {

	return "Successfully Modified Task.", errors.New("This is broken")
}
func ListTasks(command_args []string) (response string, err error) {
	return "Successfully Listed Tasks.", errors.New("This is broken")
}
