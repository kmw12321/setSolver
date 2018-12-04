  // Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Perhaps the most basic file reading task is
    // slurping a file's entire contents into memory.
    dat, err := ioutil.ReadFile("./setInput.txt")
    check(err)
    var cards = strings.Split(string(dat),"\r\n")
    // commas := regexp.MustCompile(",")
    // commaLocations := commas.FindAllStringIndex(string(dat), -1)
    // numCards = len(commaLocations) + 1
    setCount := 0
    for i := 0; i < len(cards); i++ {
      for j := i; j < len(cards); j++ {
        if (j != i){
          for k := j; k < len(cards); k++ {
            set := 1
            if (k != j && k != i){
              var card1 = strings.Split(cards[i],",")
              var card2 = strings.Split(cards[j],",")
              var card3 = strings.Split(cards[k],",")
              for c := 0; c < 4; c++ {
                if ((card1[c] == card2[c] && card1[c] == card3[c]) || (card1[c] != card2[c] && card1[c] != card3[c]) && card2[c] != card3[c]){
                } else {
                  set = 0
                }
              }
              if (set == 1){
                setCount = setCount + 1
                fmt.Println("set #" + strconv.Itoa(setCount))
                fmt.Println(cards[i])
                fmt.Println(cards[j])
                fmt.Println(cards[k])
                fmt.Println()
              }
            }
          }
        }
      }
    }
}
