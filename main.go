package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/gilmae/deputy/orgchart"
)

type UsersAndRoles struct {
	Roles []orgchart.Role
	Users []orgchart.User
}

func main() {
	// For the sake of keeping scope under control, roles and users data is loaded
	// from a json file rather than from some variety of database.
	dataFilePath := flag.String("data", "./sample.json", "Data file containing roles and users")
	supervisor := flag.Int("supervisor", 1, "Returns subordinates of the supervising user with the given id")
	flag.Parse()

	dataFile, err := os.Open(*dataFilePath)
	if err != nil {
		panic(err)
	}

	defer dataFile.Close()

	byteValue, _ := ioutil.ReadAll(dataFile)

	var data UsersAndRoles

	json.Unmarshal(byteValue, &data)

	// Set up organisation and populate users and roles
	o := orgchart.NewOrganisation()
	o.SetRoles(data.Roles)
	o.SetUsers(data.Users)

	// Get subordinates and print
	subordinates, err := o.GetSubordinates(*supervisor)
	if err != nil {
		panic(err)
	}
	fmt.Println(subordinates)
}