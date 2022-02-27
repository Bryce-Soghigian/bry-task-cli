package main

import (
	"bufio"
	"os"
	"github.com/Bryce-Soghigian/bry-task-cli/helpers"
	"fmt"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	response, err := helpers.CommandParser(command)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(response)
}
