package main

import (
	"fmt"

	passwd "github.com/zachatttack/clean-code/passwd"
)

func main() {

  var users = []string {}
  users = passwd.GetUsers()
  fmt.Printf("%s\n",users)

  return   
}
