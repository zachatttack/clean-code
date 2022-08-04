package main

import (
	passwd "github.com/zachatttack/clean-code/passwd"
	"fmt"
)

func main() {

  var users = []string {}
  users = passwd.GetUsers()
  fmt.Printf("%s\n",users)

  time := passwd.GetLastPasswordChange("user1")
  fmt.Printf("%s\n",time)

  fmt.Printf("%t\n", passwd.TestPassword("user1", "password1"))
  fmt.Printf("%t\n", passwd.TestPassword("user2", "password1"))

  passwd.AddUser("testadd", "addpass")
  passwd.SetPassword("user1", "newpassword")

}
