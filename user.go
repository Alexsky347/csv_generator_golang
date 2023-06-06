package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	randomdata "github.com/Pallinder/go-randomdata"
)

var cusRangeUser = 2000
var fileNameUser = "../csvCustom/" + strconv.Itoa(cusRangeUser) + "-users.csv"

func generateUser() {


	headers := []string{
		"userName", "userPassword", "userPasswordSalt", "userSupervisor",
		"userLanguage", "userLastName", "userFirstName", "userIdentityPicture", "userEmail", "userState", "userEmailNotification", "userSubstituteNames", "TITLE"}

	records := [][]string{}
	records = append(records, headers)

	file, err := os.Create(fileNameUser)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(file)
	csvwriter.Comma = ';'
	defer csvwriter.Flush()

	for i := range N(cusRangeUser) {
		fmt.Println(i)
		typeC := randomdata.Number(0, cusRangeUser*10)
		firstName := strings.ToLower(randomdata.FirstName(typeC))
		lastName := strings.ToLower(randomdata.LastName())
		login := strings.ToLower(firstName[0:3] + lastName[0:3])


		data := []string{
			login,
			"",
			"uudfjidfgji",
			"",
			"fr",
			lastName,
			firstName,
			"",
			login + "@itesoft-test.com",
			"present",
			"none",
			"",
			"Admin ITESOFT"}

		records = append(records, data)
	}

	err = csvwriter.WriteAll(records) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}
