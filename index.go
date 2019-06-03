package main

import (
	"TechnoRelyCourses/logic"
	"fmt"
	"log"
)

func main() {
	persons, err := logic.ParseCSV("csv-data/MOCK_DATA.csv")
	if err != nil {
		log.Fatal(err)
	}

	personsInRegisterRange, err := persons.GetPersonsInRegisterDateRange("7/16/2018", "7/26/2018") //мм, чч, гг
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Пользователи зарегестрированные с 7/16/2018 по 8/10/2017\n")
	for _, v := range personsInRegisterRange {
		fmt.Println(v)
	}

	fmt.Println("\n\n\n-----------------------------------------------------------------\n\n\n")
	fmt.Println("Пользователи отсортированные по Loan")
	persons.SortOfPerson("Loan")
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
}
