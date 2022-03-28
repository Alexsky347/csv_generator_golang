package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	randomdata "github.com/Pallinder/go-randomdata"
)


func main() {

	headers := []string{
		"userName", "userFirstName", "userLastName", "userEmail", "TITLE",
		"userLanguage", "userState", "userSupervisor", "userEmailNotification", "userIdentityPicture"}

	records := [][]string{}
	records = append(records, headers)

	file, err := os.Create("employee.csv")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(file)
	csvwriter.Comma = ';'
	defer csvwriter.Flush()


	for i := range [1000]int{} {
        fmt.Println(i)
		custType := randomdata.Male
		lastName := randomdata.LastName()
		firstName := randomdata.FirstName(custType)
		email := lastName + "." + firstName + "@gmail.com"
		username := strings.ToLower(firstName[0:1]) + strings.ToLower(lastName[0:2])

		data :=[]string{
		 	username, firstName, lastName, email, "Admin",
			"fr", "userState", "userSupervisor", "", ""}
		
		records = append(records, data)
	}
	
	err = csvwriter.WriteAll(records) // calls Flush internally

    if err != nil {
        log.Fatal(err)
    }
}

