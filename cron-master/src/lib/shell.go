package lib
import (
  "bytes"
   "fmt"
   "log"
   "os/exec"
   "time"
   "context"
)

func ExecShell(s string){
  
  ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
  //ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  //cmd := exec.Command("/bin/bash","-c",s)
  
  cmd := exec.CommandContext(ctx, "/bin/bash", "-c", s)

   
  var out bytes.Buffer
  cmd.Stdout = &out

  
  err := cmd.Start()
   if err !=nil{
    log.Fatal(err)

  }

 
  time.Sleep(10 * time.Second)
  fmt.Println("退出程序中...", cmd.Process.Pid)
  //cancel()

   err = cmd.Wait()
 
   t := time.Now()
   fmt.Printf("exec complete")
   fmt.Println(t.Unix())


  
  fmt.Printf("wait res:%v", err)
}
