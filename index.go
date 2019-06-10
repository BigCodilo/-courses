package main

import (
	"TechnoRelyCourses/interactionDB"
	"TechnoRelyCourses/logic"
	"fmt"
	"log"
	"net/http"
)

var DB interactionDB.DataBase

func main() {
	DB = interactionDB.DataBase{}
	err := DB.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func(db interactionDB.DataBase) {
		err := db.Close()
		if err != nil {
			fmt.Println("Something wrong with Database")
		}
	}(DB)

	persons, err := logic.ParseCSV("csv-data/MOCK_DATA.csv")
	if err != nil {
		log.Fatal(err)
	}

	//----- DATABASE ------//

	// for _, v := range persons {
	// 	db.Add(v)
	// }
	//db.GetAllPersons()

	personsInRegisterRange, err := persons.GetPersonsInRegisterDateRange("07/28/2018", "09/26/2018") //мм, чч, гг
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Пользователи зарегестрированные с 7/28/2018 по 9/26/2017\n")
	for _, v := range personsInRegisterRange {
		fmt.Println(v)
	}

	fmt.Println("\n\n\n-----------------------------------------------------------------\n\n\n")
	fmt.Println("Пользователи отсортированные по FirstName")
	persons.SortOfPerson("FirstName")
	for _, v := range persons {
		fmt.Println(v)
	}

	fmt.Println("\n\n\n-----------------------------------------------------------------\n\n\n")
	fmt.Println("Количество женщин и мужчин\n")
	p1 := persons.GetPersentOFGender("Male")
	p2 := persons.GetPersentOFGender("Female")
	fmt.Println(p1, " ===>", p2)

	fmt.Println("\n\n\n-----------------------------------------------------------------\n\n\n")
	fmt.Println("Пользователи по диапазону займа\n")
	personsInLoanRange := persons.GetPersentOfLoanRange(300000, 600000)
	for _, v := range personsInLoanRange {
		fmt.Println(v)
	}
	StartServer()
}

func StartServer() {
	http.HandleFunc("/persons", GetPersonHandler)
	http.HandleFunc("/add", AddPersonHandler)
	http.HandleFunc("/delete", DeletePersonHandler)
	http.HandleFunc("/update", UpdatePersonHandler)
	http.ListenAndServe(":1234", nil)
}
