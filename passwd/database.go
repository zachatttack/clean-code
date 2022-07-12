package passwd

import (
	"fmt"
	"log"
	"os"
	"time"
)

type user struct {
	name      string
	password  string
	timestamp time.Time
}

func SetPassword() {
	return
}

func TestPassword() bool {
	return false
}

func GetUsers() []string {
	users := []string{}

	return users
}

func GetLastPasswordChange() time.Time {
	date := time.Now()

	return date
}

func DeleteUser() {
	return
}

func ReadFile() {
	file, err := os.Open("database.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q\n", data[:count])
}
