package interactionDB

import (
	"TechnoRelyCourses/logic"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DataBase struct {
	Connection *sql.DB
}

func (db *DataBase) Open() error {
	connectionString := "user=postgres password=root dbname=TRely sslmode=disable"
	var err error
	db.Connection, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) Add(person logic.Person) error {
	_, err := db.Connection.Exec("insert into Persons (firstname, lastname, email, gender, genderiota, registerdate, loan)"+
		"values ($1, $2, $3, $4, $5, $6, $7)",
		person.FirstName,
		person.LastName,
		person.Email,
		person.Gender,
		person.GenderIota,
		person.RegisterDate,
		person.Loan,
	)
	if err != nil {
		return err
	}
	fmt.Println(person, " --> Was added")
	return nil
}

func (db *DataBase) Delete(id int) error {
	_, err := db.Connection.Exec("delete from Persons where id = $1", id)
	//_, err := db.connection.Exec("delete from Persons")
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) Update(id int, email string) error {
	_, err := db.Connection.Exec("update Persons set email = $1 where id = $2", email, id)
	if err != nil {
		return err
	}
	return nil
}

// func (db *DataBase) GetAllPersons() {
// 	rows, err := db.Connection.Query("select * from Persons")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	persons := []logic.Person{}

// 	for rows.Next() {
// 		p := logic.Person{}
// 		err := rows.Scan(&p.FirstName, &p.LastName, &p.ID, &p.RegisterDate, &p.Email, &p.Gender, &p.GenderIota, &p.Loan)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		persons = append(persons, p)
// 	}
// }

func (db *DataBase) Close() error {
	err := db.Connection.Close()
	return err
}
