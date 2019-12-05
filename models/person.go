package models

import (
    "fmt"
    db "my-gin/database"
)

type Person struct {
    Id          int     `json:"id" form:"id"`
    FirstName   string  `json:"first_name" form:"first_name"`
    LastName    string  `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
    // Normal Query
    // sqlStatement := `INSERT INTO person(first_name, last_name) VALUES ($1, $2) RETURNING id`
    // err = db.SqlDB.QueryRow(stmt, p.FirstName, p.LastName).Scan(&id)

    // Prepared Statement
    stmt, _ := db.SqlDB.Prepare(`INSERT INTO person(first_name, last_name) VALUES ($1, $2) RETURNING id`)
    defer stmt.Close()
    err = stmt.QueryRow(p.FirstName, p.LastName).Scan(&id)
    
    if err != nil {
        panic(err)
    }
    return
}

func (p *Person) ModPerson() (num int64, err error) {
    result, err := db.SqlDB.Exec("UPDATE person SET first_name=$1, last_name=$2 WHERE id=$3", p.FirstName, p.LastName, p.Id
    if err != nil {
        fmt.Println("update data failed:", err.Error())
        return
    }

    num, err = result.RowsAffected()
    if err != nil {
        fmt.Println("fetch row affected failed:", err.Error())
        return
    }
    fmt.Println("update records number", num)
    return
}

func (p *Person) DelPerson() (num int64, err error) {
    result, err := db.SqlDB.Exec("DELETE FROM person WHERE id = ?", p.Id)
    if err != nil {
        fmt.Println("delete data failed:", err.Error())
        return
    }

    num, err = result.RowsAffected()
    if err != nil {
        fmt.Println("fetch row affedcted failed", err.Error())
        return
    }
    fmt.Println("delete record number", num)
    return
}

func (p *Person) GetPersons() (persons []Person, err error) {
    persons = make([]Person, 0)
    rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person")
    defer rows.Close()

    if err != nil {
        return
    }

    for rows.Next() {
        var person Person
        rows.Scan(&person.Id, &person.FirstName, &person.LastName)
        persons = append(persons, person)
    }
    if err = rows.Err(); err != nil {
        return
    }

    return
}

func (p *Person) GetPerson() (err error) {
    db.SqlDB.QueryRow("SELECT id, first_name, last_name FROM person WHERE id=$1", p.Id).Scan(
        &p.Id,
        &p.FirstName,
        &p.LastName,
    )
    return
}