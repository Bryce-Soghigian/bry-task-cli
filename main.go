package main

import (
	"bufio"
	"fmt"
	"github.com/Bryce-Soghigian/bry-task-cli/helpers"
	"github.com/Bryce-Soghigian/bry-task-cli/persistence"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	db, err := gorm.Open(postgres.Open(persistence.DbURL), &gorm.Config{})
	persistence.MigrateTables(db)
	response, err := helpers.CommandParser(command, db)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(response)
}
