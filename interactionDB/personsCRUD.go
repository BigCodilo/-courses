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

func (db *DataBase) Update(id int, person logic.Person) error {
	personFromDB := db.GetPerson(id)
	if len(person.FirstName) == 0 {
		person.FirstName = personFromDB.FirstName
	}
	if len(person.Email) == 0 {
		person.Email = personFromDB.Email
	}
	if len(person.Gender) == 0 {
		person.Gender = personFromDB.Gender
	}
	if person.Loan == 0.0 {
		person.Loan = personFromDB.Loan
	}
	_, err := db.Connection.Exec("update Persons set firstname = $1, email = $2, gender = $3, loan = $4 where id = $5",
		person.FirstName,
		person.Email,
		person.Gender,
		person.Loan,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) GetAllPersons() logic.Persons {
	rows, err := db.Connection.Query("select * from Persons")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	persons := logic.Persons{}
	for rows.Next() {
		p := logic.Person{}
		err := rows.Scan(
			&p.FirstName,
			&p.LastName,
			&p.ID,
			&p.RegisterDate,
			&p.Email,
			&p.Gender,
			&p.GenderIota,
			&p.Loan,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		persons = append(persons, p)
	}
	return persons
}

func (db *DataBase) GetPerson(id int) logic.Person {
	rows, err := db.Connection.Query("select * from Persons where id = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	person := logic.Person{}
	for rows.Next() {
		err := rows.Scan(
			&person.FirstName,
			&person.LastName,
			&person.ID,
			&person.RegisterDate,
			&person.Email,
			&person.Gender,
			&person.GenderIota,
			&person.Loan,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return person
}

func (db *DataBase) Close() error {
	return db.Connection.Close()
}
