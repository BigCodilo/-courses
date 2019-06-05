//Остальной функционал
package logic

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"time"
)

var db *sql.DB

func SetDatabaseConnector(dbconn *sql.DB) {
	db = dbconn
}

//ParseCSV - parse of CSV file with data
func ParseCSV(path string) (Persons, error) {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	persons := Persons{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		registerDate, err := ParseStringToDate(line[5])
		if err != nil {
			continue
		}
		loan, err := strconv.ParseFloat(line[6], 64)
		if err != nil {
			continue
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			continue
		}
		person := Person{
			ID:           id,
			FirstName:    line[1],
			LastName:     line[2],
			Email:        line[3],
			Gender:       line[4],
			RegisterDate: *registerDate,
			Loan:         loan,
		}
		SetIotaGender(person)
		persons = append(persons, person)
	}
	return persons, nil
}

//ParseStringToDate - return a time of register (time.Time)
func ParseStringToDate(date string) (*time.Time, error) {
	// dateSlice := strings.Split(date, "/")
	// month, err := strconv.Atoi(dateSlice[0])
	// if err != nil {
	// 	return nil, errors.New("Incorrect inputing date")
	// }
	// day, _ := strconv.Atoi(dateSlice[1])
	// if err != nil {
	// 	return nil, errors.New("Incorrect inputing date")
	// }
	// year, _ := strconv.Atoi(dateSlice[2])
	// if err != nil {
	// 	return nil, errors.New("Incorrect inputing date")
	// }
	// registerTime := time.Time{}.AddDate(year-1, month-1, day-1)
	// return &registerTime, nil
	layoutParseDate := "1/2/2006"
	registerTime, err := time.Parse(layoutParseDate, date)
	if err != nil {
		return nil, errors.New("Incorrect inputing date")
	}
	return &registerTime, nil
}

//SetIotaGender set 0 or 1 to each gender
func SetIotaGender(person Person) {
	if person.Gender == "Female" {
		person.GenderIota = Female
	}
	if person.Gender == "Male" {
		person.GenderIota = Male
	}
}
