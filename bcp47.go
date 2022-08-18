package main

import (
   "fmt"
   "os"
   "golang.org/x/text/language"
)

func main() {
   os.checkBCP47(realMain(os.Stdout))
}

func checkBCP47(out io.Writer) int {
   if len(os.Args) != 1 {
      fmt.Fprintf(out,"Missing language code\n")
      return 1
   }

   tag := language.Make(os.Args[0])
   fmt.Fprintln(out,tag)
   return 0
}