package lib
import (
  "bytes"
   "fmt"
   "log"
   "os/exec"
   "time"
   "context"
   "os"
)


func ExecShell(s string){
  
  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  //ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  //cmd := exec.Command("/bin/bash","-c",s)
  

  fmt.Printf(s)
  cmd := exec.CommandContext(ctx, "/bin/bash", "-c", s)

   
  var out bytes.Buffer
  cmd.Stdout = &out

  
  err := cmd.Start()
   if err !=nil{
    log.Fatal(err)

  }

  err = cmd.Wait()
   
  t := time.Now()
  fmt.Printf("shell end time:")
  fmt.Println(t.Unix());

  fmt.Printf("shell err:%v", err)
}
func main(){
   
   fmt.Println("进程id.", os.Getpid())  
   //time.Sleep(20 * time.Second)
   
   ExecShell("sleep 60")
   //time.Sleep(10 * time.Second)
   //cancel()  
}