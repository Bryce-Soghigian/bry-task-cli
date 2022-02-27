package helpers

func AddTask(command_args []string) (response string, err error) 
	CreateNewTask(command_args[0], command_args[1])
	return "Successfully Added Task.", nil
}
func RemoveTask(command_args []string) (response string, err error) {
	return "Successfully Removed Task.", nil
}
func ModifyTask(command_args []string) (response string, err error) {

	return "Successfully Modified Task.", nil
}
func ListTasks(command_args []string) (response string, err error) {
	return "Successfully Listed Tasks.", nil
}
