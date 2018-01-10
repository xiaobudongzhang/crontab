package lib
import (
  "bytes"
   "fmt"
   "log"
   "os/exec"
)
func ExecShell(s string){
  cmd := exec.Command("/bin/bash","-c",s)
  var out bytes.Buffer

  cmd.Stdout = &out

  err := cmd.Run()
  if err !=nil{
    log.Fatal(err)
  }

  fmt.Printf("%s", out.String())
}
