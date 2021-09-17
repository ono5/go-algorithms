package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
)

type User struct {
	Email string
	Name  string
}

var Database []User

func init() {
	// ユーザー3000000人分
	for i := 0; i < 3000000; i++ {
		user := User{
			Email: faker.Email(),
			Name:  faker.Name(),
		}
		Database = append(Database, user)
	}

	specificedUser := User{
		Email: "selfnote@work",
		Name:  "selfnote",
	}

	Database = append(Database, specificedUser)
}

type Worker struct {
	database []User
}

func NewWorker(database []User) *Worker {
	return &Worker{database: database}
}

func (w *Worker) Find(email string) *User {
	for i := range w.database {
		user := &w.database[i]
		if user.Email == email {
			return user
		}
	}
	return nil
}

func main() {
	email := os.Args[1]
	w := NewWorker(Database)

	start := time.Now()
	fmt.Printf("Looking for %s\n", email)

	user := w.Find(email)

	if user != nil {
		fmt.Printf("The email %s is owned by %s\n", email, user.Name)
	} else {
		fmt.Printf("The email %s was not found\n", email)
	}

	fmt.Printf("It took %d ms\n", time.Since(start).Milliseconds())
}
