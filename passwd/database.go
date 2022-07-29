package passwd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
  "time"
)

type user struct {
	name      string
	password  string
	timestamp string
}

var users = [100]user{}
var userCount = 0

func SetPassword() {
	return
}

func TestPassword(username string, password string) bool {
	readFile()
	for i := 0; i < userCount; i++ {
		if username == users[i].name {
			if password == users[i].password {
				return true
			}
		}
	}
	return false
}

func GetUsers() []string {
	readFile()
	var usernames = []string{}
	for i := 0; i < userCount; i++ {
		usernames = append(usernames, users[i].name)
	}
	return usernames
}

func GetLastPasswordChange(username string) string {
	readFile()
	ret := "nil"
	for i := 0; i < userCount; i++ {
		if username == users[i].name {
			ret = users[i].timestamp
		}
	}
	return ret
}

func DeleteUser() {
	return
}

func readFile() {
	userCount = 0
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
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func AddUser(name string, password string) {
	readFile()

	users[userCount].name = name
	users[userCount].password = password
	users[userCount].timestamp = time.Now().String()
	userCount++

	fmt.Printf("userCount: %d\n", userCount)
	writeFile()
}

func writeFile() {
	file, err := os.Create("database2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for i := 0; i < userCount; i++ {
		writer.WriteString(users[i].name + "," + users[i].password + "," + users[i].timestamp + "\n")
	}

	writer.Flush()
}
