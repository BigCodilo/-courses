package main

import (
	"TechnoRelyCourses/models"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	persons, err := ParseCSV("csv-data/MOCK_DATA.csv")
	if err != nil {
		log.Fatal(err)
	}
	personInRegRange, _ := GetPersonsInRegisterDateRange(persons, "7/16/2018", "12/28/2018")
	fmt.Println(personInRegRange)
	p1 := GetPersentOFGender(persons, "Male")
	p2 := GetPersentOFGender(persons, "Female")
	fmt.Println(p1, " ===>", p2)
	SetIotaGender(persons)
	fmt.Println(persons)
}

func ParseCSV(path string) ([]models.Person, error) {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	persons := []models.Person{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		month, day, year, _ := ParseStringToDate(line[5])
		persons = append(persons, models.Person{
			ID:        line[0],
			FirstName: line[1],
			LastName:  line[2],
			Email:     line[3],
			Gender:    line[4],
			RegisterDate: models.Date{
				Month: month,
				Day:   day,
				Year:  year,
			},
			Loan: line[6],
		})
	}
	return persons, nil
}

func SetIotaGender(persons []models.Person) {
	for i := 0; i < len(persons); i++ {
		if persons[i].Gender == "Female" {
			persons[i].GenderIota = models.Female
		}
		if persons[i].Gender == "Male" {
			persons[i].GenderIota = models.Male
		}
	}
}

func GetPersonsInRegisterDateRange(persons []models.Person, fromDate, toDate string) ([]models.Person, error) {
	fromMonth, fromDay, fromYear, err := ParseStringToDate(fromDate)
	if err != nil {
		return nil, err
	}
	toMonth, toDay, toYear, err := ParseStringToDate(toDate)
	if err != nil {
		return nil, err
	}
	personsInRegisterRange := []models.Person{}
	for _, value := range persons {
		if value.RegisterDate.Year > fromYear &&
			value.RegisterDate.Year < toYear {
			personsInRegisterRange = append(personsInRegisterRange, value)
		}
		if value.RegisterDate.Year == fromYear ||
			value.RegisterDate.Year == toYear {
			if value.RegisterDate.Month > fromMonth &&
				value.RegisterDate.Month < toMonth {
				personsInRegisterRange = append(personsInRegisterRange, value)
			}
			if value.RegisterDate.Month == fromMonth ||
				value.RegisterDate.Month == toMonth {
				if value.RegisterDate.Day > fromDay &&
					value.RegisterDate.Day < toDay {
					personsInRegisterRange = append(personsInRegisterRange, value)
				}
			}
		}
	}
	return personsInRegisterRange, nil
}

func GetPersentOFGender(persons []models.Person, gender string) float64 {
	var count int
	sumOfInputed := len(persons)
	for _, v := range persons {
		if v.Gender != "Male" && v.Gender != "Female" {
			sumOfInputed--
		}
		if v.Gender == gender {
			count++
		}
	}
	return float64(count) / 100 * float64(sumOfInputed)
}

func ParseStringToDate(date string) (int, int, int, error) {
	dateSlice := strings.Split(date, "/")
	month, err := strconv.Atoi(dateSlice[0])
	if err != nil {
		return -1, -1, -1, err
	}
	day, err := strconv.Atoi(dateSlice[1])
	if err != nil {
		return -1, -1, -1, err
	}
	year, err := strconv.Atoi(dateSlice[2])
	if err != nil {
		return -1, -1, -1, err
	}
	return month, day, year, nil
}
