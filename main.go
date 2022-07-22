package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strconv"
)

func main() {

    f, err := os.Create("data.log")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()
    i := 0
    for {
      fmt.Println("Delak log 22222")
      dt := time. Now()
      _, err2 := f.WriteString(dt.String() + " log entry" + strconv.Itoa(i) + "\n")
      i = i + 1
      if err2 != nil {
          log.Fatal(err2)
      }
      time.Sleep(1 * time.Second)
    }
    fmt.Println("done")
}
