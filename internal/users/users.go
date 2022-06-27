package users

import (
	"fmt"
	"log"

	database "github.com/joud-cxu/hackernews/internal/pkg/db/migrations/mysql"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user User) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []User {
	stmt, err := database.Db.Prepare("select id, username from Users")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

func GetUserById(id string) (*User, error) {
	fmt.Println("IDDDDDDD L ", id)
	stmt, err := database.Db.Prepare("select * from Users where id = ?")
	if err != nil {
		fmt.Println("er111111111111111111111 : ", err)
		log.Fatal(err)
	}
	var user User

	err2 := stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Password)
	if err2 != nil {
		fmt.Println("er222222222222 : ", err2)
		log.Fatal(err2)
		return &user, err2
	}
	fmt.Println("USERrrrrr : /", user)
	return &user, nil
}
