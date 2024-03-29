package main

import (
  "fmt"
  "os"
  "log"
  "github.com/mattn/go-sqlite3"
)

func main() {

}

func init() String {
  homedir, err := os.UserHomeDir()
  if err != nil {
    log.Fatal(err)
  }
  
  String path = homedir + "/.passman"
  err := os.Mkdir(&path, 0700)
    
  }
}
