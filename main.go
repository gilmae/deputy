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
	dataFilePath := flag.String("data", "./sample.json", "Data file containing roles and users")
	flag.Parse()

	dataFile, err := os.Open(*dataFilePath)
	if err != nil {
		panic(err)
	}

	defer dataFile.Close()

	byteValue, _ := ioutil.ReadAll(dataFile)

	var data UsersAndRoles

	json.Unmarshal(byteValue, &data)

	o := orgchart.NewOrganisation()
	o.SetRoles(data.Roles)
	o.SetUsers(data.Users)

	subordinates, err := o.GetSubordinates(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(subordinates)

	subordinates, err = o.GetSubordinates(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(subordinates)

}