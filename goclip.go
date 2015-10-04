package main

import "os"
import "io/ioutil"
import "fmt"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func contains(s []string, e string) bool {
  for _, a := range s {
    if a == e {
        return true
    }
  }
  return false
}

func main() {
  argsWithoutProg := os.Args[1:]
  homedir := os.Getenv("HOME")

  clippath := fmt.Sprint(homedir, "/.goclip")

  if contains(argsWithoutProg, "-i") {
    bytes, _ := ioutil.ReadAll(os.Stdin)
    err := ioutil.WriteFile(clippath, bytes, 0644)
    check(err)
  } else if contains(argsWithoutProg, "-o") {
    dat, err := ioutil.ReadFile(clippath)
    check(err)
    fmt.Print(string(dat))
  }
}
