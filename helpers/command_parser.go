package helpers

import (
	"errors"
	"gorm.io/gorm"
	"strings"
)

func validateFlag(flag string) (output bool) {
	return string(flag[0]) == "-"
}
func validateContent(content string) (output bool) {
	return len(content) == 0
}
func ValidateCommandArgs(command_args []string) (output bool) {
	for i := 0; i < len(command_args); i += 2 {
		flag := command_args[i]
		if !validateFlag(flag) {
			return false
		}
		content := command_args[i+1]
		if !validateContent(content) {
			return false
		}
	}
	return true
}

func GetCommandParts(command_string string) (command_header string, command_type string, command_args []string) {
	split_string := strings.Split(command_string, " ")
	command_header = string(split_string[0])
	command_type = string(split_string[1])
	command_args = split_string[2:]
	return command_header, command_type, command_args
}

func CommandParser(command_string string, db *gorm.DB) (response string, err error) {

	if len(command_string) < 9 {
		return "", errors.New("invalid command string. its too short somethings missing")
	}
	command_header, command_type, command_args := GetCommandParts(command_string)
	if command_header != "task" {
		return "", errors.New("invalid command header")
	}
	ValidateCommandArgs(command_args)
	if command_type == "complete" {
		return ModifyTask(command_args)
	} else if command_type == "create" {
		return AddTask(command_args, db)
	} else if command_type == "list" {
		return ListTasks(db, command_args)
	} else if command_type == "abort" {
		return RemoveTask(command_args)
	} else {
		return "", errors.New("not a valid command type")
	}
}
