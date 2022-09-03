package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

func prettyPrint(s string, c string){
  fmt.Println(c, "=> ", s)
}

func printError(s string){
  cc_red := "\033[31m"
  fmt.Println(cc_red, s)
}

func printHelp(){
  fmt.Println("\033[34m==> RGEN is a Random Generator  <==\033[0m")
  fmt.Println("\n\033[32mCommands:\033[0m")
  fmt.Println("-----------------------------------------------")
  fmt.Println("\033[32muuid         Generates UUID.")
  fmt.Println("pwd          Generates strong password.")
  fmt.Println("hex *size    Generates hex with input size.")
  fmt.Println("base64 *size Generates Base64 with input size.")
  fmt.Println("prime *size  Generates prime number with input size.\033[0m\n")
}

func main(){
  cc_green := "\033[32m"

  if len(os.Args) < 2 {
    printHelp()
  } else {
    switch os.Args[1] {
    case "uuid":
      prettyPrint(uuid.NewString(), cc_green)
    case "pwd":
      prettyPrint(randstr.String(12), cc_green)
    case "hex":
      if len(os.Args) < 3 {
        printError("Error: Plase enter a lenght!")
        return
      } else {
        size, err := strconv.Atoi(os.Args[2])
        if err != nil {
          printError("Error: Please enter a valid length!")
          return
        }
        prettyPrint(randstr.Hex(size), cc_green)
      }
    case "base64":
      if len(os.Args) < 3 {
        printError("Error: Plase enter a lenght!")
        return
      } else {
        size, err := strconv.Atoi(os.Args[2])
        if err != nil {
          printError("Error: Please enter a valid length!")
          return
        }
        prettyPrint(randstr.Base64(size), cc_green)
      }
    case "prime":
      if len(os.Args) < 3 {
        printError("Error: Please enter a length")
        return
      } else {
        size, err := strconv.Atoi(os.Args[2])
        if err != nil {
          printError("Error: Please enter a valid length!")
          return
        }
        p, err := rand.Prime(rand.Reader, size)
        if err != nil {
          printError(err.Error())
          return
        }
        prettyPrint(p.String(), cc_green)
      }
    default:
      printHelp()
    }
  }
}
