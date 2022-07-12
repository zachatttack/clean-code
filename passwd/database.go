package passwd

import (
	"fmt"
	"log"
	"os"
)

// void SetPassword(username, newPassword);
// bool TestPassword(username, passwordToTry);
// string[] GetUsers(void);
// Date GetLastPasswordChange(username);
// void DeleteUser(username, passwordToTry);

func ReadFile(){
  file, err := os.Open("database.txt") 
  if err != nil {
    log.Fatal(err)
  }

  data :=make([]byte,100)
  count, err := file.Read(data)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%q\n", data[:count])
}
