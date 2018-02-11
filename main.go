package main

import (
  "strings"
  "log"
  "fmt"
  "github.com/sethvargo/go-diceware/diceware"
)

func main() {
  // Generate 4 words using the diceware algorithm.

  // TODO: read some parameters from args:
  //  * number of words to generate
  //  * set a delimiter other than space e.g. correct-horse-battery-stapler
  
  list, err := diceware.Generate(4)
  if err != nil  {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", strings.Join(list, " "))
}

