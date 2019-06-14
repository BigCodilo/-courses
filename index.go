package main

import (
	"TechnoRelyCourses/interactionDB"
	"TechnoRelyCourses/logger"
	"TechnoRelyCourses/logic"
	"fmt"
	"net/http"
)

var DB interactionDB.DataBase

func main() {
	logger.SetLoggers()

	DB = interactionDB.DataBase{}
	err := DB.Open()
	logger.Info.Println("connection with database opened")
	if err != nil {
		logger.Error.Println("problem with open connection with database", err)
	}

	defer func(db interactionDB.DataBase) {
		err := db.Close()
		if err != nil {
			logger.Error.Println("problem with closing database", err)
			return
		}
		logger.Info.Println("connection with database closed")
	}(DB)

	persons, err := logic.ParseCSV("csv-data/MOCK_DATA.csv")
	if err != nil {
		logger.Error.Println("problem with parcing CSV", err)
	}
	logger.Info.Println("trying parsed CSV file")

	//----- DATABASE ------//

	// for _, v := range persons {
	// 	db.Add(v)
	// }
	//db.GetAllPersons()

	personsInRegisterRange, err := persons.GetInRegisterRange("07/28/2018", "09/26/2018") //мм, чч, гг
	if err != nil {
		logger.Error.Println("problem with search persons ia register range")
	}
	logger.Info.Println("getted person with range from 07/28/2018 to 09/26/2018")

	fmt.Println("Пользователи зарегестрированные с 7/28/2018 по 9/26/2018\n")
	for _, v := range personsInRegisterRange {
		fmt.Println(v)
	}

	fmt.Println("\n\n\n-----------------------------------------------------------------\n\n\n")
	fmt.Println("Пользователи отсортированные по FirstName")
	logger.Info.Println("persons sorted by FirstName")
	err = persons.SortOfPerson("FirstName")
	if err != nil {
		logger.Error.Println("problem with sorting. ", err)
	}
	for _, v := range persons {
		fmt.Println(v)
	}

	fmt.Println("\n\n\n-----------------------------------------------------------------\n\n\n")
	fmt.Println("Количество женщин и мужчин\n")
	logger.Info.Println("getted persent of Male or Female")
	p1 := persons.GetPersentOFGender("Male")
	p2 := persons.GetPersentOFGender("Female")
	fmt.Println(p1, " ===>", p2)

	fmt.Println("\n\n\n-----------------------------------------------------------------\n\n\n")
	fmt.Println("Пользователи по диапазону займа\n")
	logger.Info.Println("getted persons it LOAN range 300000:600000")
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
	logger.Info.Println("server started")
	http.ListenAndServe(":1234", nil)
}
