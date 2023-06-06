package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	randomdata "github.com/Pallinder/go-randomdata"
)


func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("No argument provided.")
		return
	}
	
	// excluding the program name itself
	cusRange := os.Args[1]
	fmt.Println("Argument %d: %s\n", cusRange)

	//convert arg to int
	newRange, err := strconv.Atoi(cusRange)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//file name
	fileName := "../csvCustom/" + strconv.Itoa(newRange) + "-routes.csv"

	flow := []string{"ACME-BAP"}

	// users := []string{"acme", "agnes", "alan", "anais", "arnaud", "aude", "benjamin", "benoit", "benoit2", "catherine", "christophe", "dev", "didier", "eva", "fabrice", "fgf", "francois", "frauke", "frederic", "frederique", "fritzi", "guillaume", "heinz", "itesoft", "jade", "james", "jason", "jean-luc", "john", "kevin", "laurent", "marie", "mike", "monika", "nick", "patrice", "paul", "pierrick", "rachid", "regis", "sarah", "service", "sonia", "stan", "thierry", "tom", "yves"}
	users := []string{"robert", "marcel", "michel"}
	headers := []string{
		"flow", "label", "companyCode", "siteCode", "usersLevel1",
		"netAmount2", "usersLevel2", "netAmount3", "usersLevel3", "netAmount4", "usersLevel4"}

	records := [][]string{}
	records = append(records, headers)

	file, err := os.Create(fileName)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(file)
	csvwriter.Comma = ';'
	defer csvwriter.Flush()

	for i := range N(newRange) {
		fmt.Println(i)
		label := randomdata.SillyName()
		site := "ACME_" + randomdata.City()
		siteCountry := "ACME_" + randomdata.Country(randomdata.TwoCharCountry) + strconv.Itoa(i)

		lengthUsers := len(users)

		data := []string{
			flow[0],
			label, siteCountry, site,
			users[randomdata.Number(0, lengthUsers)],
			strconv.Itoa(randomdata.Number(10, 10000)), users[randomdata.Number(0, lengthUsers)],
			strconv.Itoa(randomdata.Number(10, 10000)), users[randomdata.Number(0, lengthUsers)],
			strconv.Itoa(randomdata.Number(10, 10000)), users[randomdata.Number(0, lengthUsers)]}

		records = append(records, data)
	}

	err = csvwriter.WriteAll(records) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}
