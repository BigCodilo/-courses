package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Person struct {
	FirstName string  `json:"name"`
	LastName  string  `json:"surname"`
	Email     string  `json:"email"`
	Gender    string  `json:"gender"`
	Loan      float64 `json:"loan"`
}

func main() {
	AddingTest()
	DeleteTest()
	UpdateTest()
}

func AddingTest() {
	person := Person{
		FirstName: "oleg",
		LastName:  "osyka",
		Email:     "olegosyka@gmail.com",
		Gender:    "Male",
		Loan:      254.3,
	}
	personJSON, _ := json.Marshal(person)
	resp, err := http.PostForm("http://localhost:1234/add", url.Values{
		"person": {string(personJSON)},
	})
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func DeleteTest() {
	resp, err := http.PostForm("http://localhost:1234/delete", url.Values{
		"id": {"110"},
	})
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func UpdateTest() {
	type IDPerson struct {
		ID     int    `json:"id"`
		Person Person `json:"person"`
	}
	idPerson := IDPerson{
		ID: 112,
		Person: Person{
			FirstName: "jeka",
			Email:     "olegosyka@gmail.comcomcom",
			Gender:    "Male",
			Loan:      254.3,
		},
	}
	idPersonJSON, _ := json.Marshal(idPerson)
	resp, err := http.PostForm("http://localhost:1234/update", url.Values{
		"idperson": {string(idPersonJSON)},
	})
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
