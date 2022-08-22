package main

import (
   "fmt"
   "os"
   "io"
   "golang.org/x/text/language"
)

func main() {
   os.Exit(checkBCP47(os.Stdout))
}

func checkBCP47(out io.Writer) int {
   argsWithoutProg := os.Args[1:]
   if len(argsWithoutProg) != 1 || (len(argsWithoutProg) == 1 && argsWithoutProg[0] == "") {
      fmt.Fprintln(out,"")
   } else {
      tag, err := language.Parse(argsWithoutProg[0])
		if err != nil {
         fmt.Fprintln(out,"")
      }
      else {
         fmt.Fprintln(out,tag)
      }
   }
   return 0
}