package main

import (
	"TechnoRelyCourses/logic"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fromDate := r.URL.Query().Get("fromDate")
	toDate := r.URL.Query().Get("toDate")
	gender := r.URL.Query().Get("gender")
	persons := DB.GetAllPersons()
	validPersons := logic.Persons{}
	//Если в урле есть запрос по полу или мени, то в этом цикле это отфильтрует
	for _, v := range persons {
		nameFlag := true
		genderFlag := true
		if len(gender) > 0 && v.Gender != gender {
			genderFlag = false
		}
		if len(name) > 0 && v.FirstName != name {
			nameFlag = false
		}
		if nameFlag && genderFlag {
			validPersons = append(validPersons, v)
		}
	}
	//Если есть запрос по дате регитсрации  то вот здесь его обработает
	if len(fromDate) > 0 && len(toDate) > 0 {
		validPersons, _ = validPersons.GetPersonsInRegisterDateRange(fromDate, toDate)
	}
	if len(fromDate) == 0 && len(toDate) > 0 {
		validPersons, _ = validPersons.GetPersonsInRegisterDateRange("01/01/1001", toDate)
	}
	if len(fromDate) > 0 && len(toDate) == 0 {
		validPersons, _ = validPersons.GetPersonsInRegisterDateRange(fromDate, "12/31/3000")
	}
	//Если ненайдено людей с удовлетворяющими требованиями то напишет об этом
	if len(validPersons) == 0 {
		http.Error(w, "No results for this query.", http.StatusNotFound)
		return
	}
	validPersonsJSON, _ := json.Marshal(validPersons)
	w.Write(validPersonsJSON)
}

func AddPersonHandler(w http.ResponseWriter, r *http.Request) {
	personJSON := r.FormValue("person")
	person := logic.Person{}
	err := json.Unmarshal([]byte(personJSON), &person)
	if err != nil {
		http.Error(w, "Uncorrect format", http.StatusBadRequest)
		return
	}
	person.RegisterDate = time.Now()
	logic.SetIotaGender(person)
	err = DB.Add(person)
	if err != nil {
		http.Error(w, "Problem with database", 418)
		return
	}
	fmt.Println(person, "---> added")
	w.Write([]byte("Answer from server: person added"))
}

func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Uncorrect format", http.StatusBadRequest)
		return
	}
	err = DB.Delete(id)
	if err != nil {
		http.Error(w, "Something wrong", 418)
		return
	}
	fmt.Println("id: ", id, "---> deleted")
	w.Write([]byte("user deleted"))
}

func UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	type IDPerson struct {
		ID     int          `json:"id"`
		Person logic.Person `json:"person"`
	}

	idPerson := IDPerson{}
	idPersonJSONString := r.FormValue("idperson")
	err := json.Unmarshal([]byte(idPersonJSONString), &idPerson)
	if err != nil {
		http.Error(w, "Something wrong", 418)
		return
	}
	err = DB.Update(idPerson.ID, idPerson.Person)
	if err != nil {
		http.Error(w, "Something wrong", 418)
		return
	}
	fmt.Println(idPerson.Person, "---> updated")
	w.Write([]byte("update succeseful"))
}
