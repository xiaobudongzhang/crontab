package lib
import (
  "fmt"
  "time"
)

func PrintNow(s string){
        t := time.Now()
        fmt.Printf(s)
        fmt.Println(t.Unix());


}