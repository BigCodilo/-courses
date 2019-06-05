package interactionDB

import (
	"TechnoRelyCourses/logic"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DataBase struct {
	Connection *sql.DB
}

func (db *DataBase) Open() {
	connectionString := "user=postgres password=root dbname=TRely sslmode=disable"
	var err error
	db.Connection, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) Add(person logic.Person) {
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
		log.Fatal(err)
	}
	fmt.Println(person, " was added")
}

func (db *DataBase) Delete(id int) {
	_, err := db.Connection.Exec("delete from Persons where id = $1", id)
	//_, err := db.connection.Exec("delete from Persons")
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) Update(id int, email string) {
	_, err := db.Connection.Exec("update Persons set email = $1 where id = $2", email, id)
	if err != nil {
		panic(err)
	}
}

func (db *DataBase) GetAllPersons() {
	rows, err := db.Connection.Query("select * from Persons")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	persons := []logic.Person{}

	for rows.Next() {
		p := logic.Person{}
		err := rows.Scan(&p.FirstName, &p.LastName, &p.ID, &p.RegisterDate, &p.Email, &p.Gender, &p.GenderIota, &p.Loan)
		if err != nil {
			fmt.Println(err)
			continue
		}
		persons = append(persons, p)
	}
	for _, p := range persons {
		fmt.Println(p)
	}
}

func (db *DataBase) Close() error {
	err := db.Connection.Close()
	return err
}
