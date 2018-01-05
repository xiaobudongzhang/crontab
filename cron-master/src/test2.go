package main

import(
  "fmt"
  "os"
  "time"
)
func main(){
   time.Sleep(60 * time.Second)
   for{
   }
   fmt.Println("退出程序中...", os.Getpid())
}