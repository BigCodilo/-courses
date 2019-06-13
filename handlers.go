package main

import (
	"TechnoRelyCourses/logger"
	"TechnoRelyCourses/logic"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func GetRequestNumber() []byte {
	reqRange := "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	reqNumber := make([]byte, 10)
	for i := range reqNumber {
		reqNumber[i] = reqRange[rand.Intn(len(reqRange))]
	}
	return reqNumber
}

func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	reqNumber := GetRequestNumber()
	name := r.URL.Query().Get("name")
	fromDate := r.URL.Query().Get("fromDate")
	toDate := r.URL.Query().Get("toDate")
	gender := r.URL.Query().Get("gender")
	persons, err := DB.GetAllPersons()
	if err != nil {
		http.Error(w, "something wrong with database", http.StatusNotFound)
		logger.Info.Println("GET for", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("GET for", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	err = nil
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
		validPersons, err = validPersons.GetInRegisterRange(fromDate, toDate)
	}
	if len(fromDate) == 0 && len(toDate) > 0 {
		validPersons, err = validPersons.GetInRegisterRange("01/01/2001", toDate)
	}
	if len(fromDate) > 0 && len(toDate) == 0 {
		validPersons, err = validPersons.GetInRegisterRange(fromDate, "12/31/3000")
	}
	if err != nil {
		http.Error(w, "uncorrect date form", http.StatusNotFound)
		logger.Info.Println("GET for", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("GET for", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	//Если ненайдено людей с удовлетворяющими требованиями то напишет об этом
	if len(validPersons) == 0 {
		http.Error(w, "no results for this query", http.StatusNotFound)
		logger.Info.Println("GET for", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("GET for", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	validPersonsJSON, err := json.Marshal(validPersons)
	if len(validPersons) == 0 {
		http.Error(w, "json error", http.StatusNotFound)
		logger.Info.Println("GET for", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("GET for", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	logger.Info.Println("GET for", r.RequestURI, "successfully. #"+string(reqNumber))
	w.Write(validPersonsJSON)
}

func AddPersonHandler(w http.ResponseWriter, r *http.Request) {
	//personJSON := r.FormValue("person")
	reqNumber := GetRequestNumber()
	person := logic.Person{}
	//err := json.Unmarshal([]byte(personJSON), &person)
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "uncorrect format", http.StatusBadRequest)
		logger.Info.Println("POST with body ", person, "to", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("POST with body ", person, "to", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	person.RegisterDate = time.Now()
	logic.SetIotaGender(person)
	err = DB.Add(person)
	if err != nil {
		http.Error(w, "Problem with database", 418)
		logger.Info.Println("POST with body ", person, "to", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("POST with body ", person, "to", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	logger.Info.Println("POST with body ", person, "to", r.RequestURI, ". Added. #"+string(reqNumber))
	w.Write([]byte("Answer from server: person added"))
}

func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	reqNumber := GetRequestNumber()
	var id int
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, "Uncorrect format -", http.StatusBadRequest)
		logger.Info.Println("POST with body ", id, " to ", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("POST with body ", id, "to", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	//id, err := strconv.Atoi(idS)
	if err != nil {
		http.Error(w, "Uncorrect format", http.StatusBadRequest)
		logger.Info.Println("POST with body ", id, " to ", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("POST with body ", id, "to", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	err = DB.Delete(id)
	if err != nil {
		http.Error(w, "Something wrong", 418)
		logger.Info.Println("POST with body ", id, " to ", r.RequestURI, " unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("POST with body ", id, "to", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	logger.Info.Println("POST with body ", id, " to ", r.RequestURI, ". Deleted. #"+string(reqNumber))
	w.Write([]byte("user deleted"))
}

func UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	reqNumber := GetRequestNumber()
	type IDPerson struct {
		ID     int          `json:"id"`
		Person logic.Person `json:"person"`
	}

	idPerson := IDPerson{}
	err := json.NewDecoder(r.Body).Decode(&idPerson)
	if err != nil {
		http.Error(w, "Something wrong", 418)
		logger.Info.Println("POST with body ", idPerson, " to ", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("POST with body ", idPerson, "to", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	err = DB.Update(idPerson.ID, idPerson.Person)
	if err != nil {
		http.Error(w, "Something wrong", 418)
		logger.Info.Println("POST with body ", idPerson, " to ", r.RequestURI, "unsuccessfully. #"+string(reqNumber))
		logger.Error.Println("POST with body ", idPerson, "to", r.RequestURI, "---", err, "---", "#"+string(reqNumber))
		return
	}
	logger.Info.Println("POST with body ", idPerson, " to ", r.RequestURI, ". Updated. #"+string(reqNumber))
	w.Write([]byte("update succeseful"))
}
