package main

import (
	"fmt"
	"os"
	"strings"
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
	ch       chan *User
	database []User
	name     string
}

func NewWorker(ch chan *User, database []User, name string) *Worker {
	return &Worker{ch: ch, database: database, name: name}
}

func (w *Worker) Find(email string) {
	for i := range w.database {
		user := &w.database[i]
		if strings.Contains(user.Email, email) {
			fmt.Println(">>", w.name)
			w.ch <- user
		}
	}
}

func main() {
	email := os.Args[1]

	ch := make(chan *User)

	start := time.Now()
	fmt.Printf("Looking for %s\n", email)

	go NewWorker(ch, Database[:1000000], "#1").Find(email)
	go NewWorker(ch, Database[1000000:2000000], "#2").Find(email)
	go NewWorker(ch, Database[2000000:], "#3").Find(email)

	select {
	case user := <-ch:
		fmt.Printf("The email %s is owned by %s\n", email, user.Name)
	case <-time.After(15 * time.Second):
		fmt.Printf("The email %s was not found\n", email)
	}

	fmt.Printf("It took %d ms\n", time.Since(start).Milliseconds())
}
