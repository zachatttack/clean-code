package passwd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
  "strings"
)

type user struct {
	name      string
	password  string
	timestamp string
}

var users = [100]user {}
var userCount = 0

func SetPassword() {
	return
}

func TestPassword() bool {
	return false
}

func GetUsers() []string {
  for i := range users {
    name[i] = users[i].name 
  }
	return name
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
	defer file.Close()

	scanner := bufio.NewScanner(file)

  var userString [100]string

	for scanner.Scan() {
    userString[userCount] = scanner.Text()
    splitString := strings.Split(userString[userCount], ",")

    users[userCount].name = splitString[0] 
    users[userCount].password = splitString[1] 
    users[userCount].timestamp = splitString[2] 

    userCount++
		// fmt.Printf("%s\n", time.Now())
	}

  fmt.Printf("%s  \n", users[0].name)
  fmt.Printf("%s  \n", users[1].name)

  fmt.Printf("%s  \n", users[0].password)
  fmt.Printf("%s  \n", users[1].password)

  fmt.Printf("%s  \n", users[0].timestamp)
  fmt.Printf("%s  \n", users[1].timestamp)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
