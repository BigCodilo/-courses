//Тут находится модель пользователя и все методы связанные с ним
package logic

import (
	"errors"
	"reflect"
	"sort"
	"time"
)

type Persons []Person

type Gender int

type Person struct {
	ID           int
	FirstName    string
	LastName     string
	Email        string
	Gender       string
	GenderIota   Gender
	RegisterDate time.Time
	Loan         float64
}

const (
	Male = iota
	Female
)

//SortOfPerson sortes a slice by one of the basic types (int, string, float64) ----> (reflection, reflection, reflection)
func (persons Persons) SortOfPerson(fieldName string) {
	refletOfStruct := reflect.ValueOf(persons[1])              //Получение типа структуры
	_, ifExist := refletOfStruct.Type().FieldByName(fieldName) //Проверка на существования поля по введенному названию
	if ifExist {                                               //Если сущестует
		sort.Slice(persons, func(i, j int) bool { //Начало сортировки
			refletOfStructInSort1 := reflect.ValueOf(persons[i])                     //Получения типа сравниваемой персоны №1
			field1 := reflect.Indirect(refletOfStructInSort1).FieldByName(fieldName) //Получения поля для сортировки по названию для персоны №1
			refletOfStructInSort2 := reflect.ValueOf(persons[j])                     //Получение типа сравниваемой персоны №2
			field2 := reflect.Indirect(refletOfStructInSort2).FieldByName(fieldName) //Получения поля для сортировки по названию для персоны №2
			switch field1.Type().String() {                                          //Ищется выбраный тип поля
			case "int": //Если это число
				return int(field1.Int()) < int(field2.Int()) //Говорится что значение поля это число и сравниваем числа
			case "string": //Если это строка
				return field1.String() < field2.String() //Говорится что значение поля это строка и сравниваем строки
			case "float64": //Если это число с плавающей точкой
				return field1.Float() < field2.Float() //Говорится что значение поля это дробное число и сравниваем дробные числа
			}
			return true
		})
	}
}

//GetPersonsInRegisterDateRange - return slice of person where each of person was register in range of inputing date
// it returns error if we haven't found any persons in inpited range
func (persons Persons) GetPersonsInRegisterDateRange(fromDate, toDate string) (Persons, error) {
	fromParseDate, _ := ParseStringToDate(fromDate)
	toParseDate, _ := ParseStringToDate(toDate)
	rangeOfInputesDate := toParseDate.Sub(*fromParseDate) //Вычитание дат с которого и по которую искать
	personsInRagisterRange := []Person{}
	for _, v := range persons {
		rangeOfPersonsDate := fromParseDate.Sub(v.RegisterDate)                          //Вычитания даты начала регистрации и даты когда зарегестрировался пользователь
		if rangeOfPersonsDate.Hours() < 0 && rangeOfInputesDate+rangeOfPersonsDate > 0 { //Херня которую и закомментировать сложно
			personsInRagisterRange = append(personsInRagisterRange, v)
		}
	}
	if len(personsInRagisterRange) == 0 {
		return nil, errors.New("No persons in this range")
	}
	return personsInRagisterRange, nil
}

//GetPersentOFGender return a persent of each gender
func (persons Persons) GetPersentOFGender(gender string) float64 {
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

func (persons Persons) GetPersentOfLoanRange(fromLoan, toLoan float64) Persons {
	personsInRange := Persons{}
	for _, v := range persons {
		if v.Loan > fromLoan && v.Loan < toLoan {
			personsInRange = append(personsInRange, v)
		}
	}
	return personsInRange
}
