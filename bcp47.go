package main

import (
   "flag"
   "fmt"
   "io"
   "os"
   "golang.org/x/text/language"
)

func main() {
   os.Exit(checkBCP47(os.Stdout))
}

func checkBCP47(out io.Writer) int {
   name := flag.String("name", "", "Your Name")
   flag.Parse()
   if *name == "" {
       fmt.Fprintf(out, "Missing flag -name\n")
       return 1
   }
   fmt.Fprintf(out, language.Make(*name))
   return 0
}