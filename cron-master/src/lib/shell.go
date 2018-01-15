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


func ExecShell(s string,id int){
  
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

  //将当前进程加入数据config中
  nowPid := cmd.Process.Pid

  datap := DataProcess{nowPid}
  log.Println(datap)
  log.Println(DataConfigs[id].Process)
  log.Println(nowPid)
  DataConfigs[id].Process[111] = datap

  err = cmd.Wait()
  //进程已结束，将当前进程数据信息从数据中移除  
  delete(DataConfigs[id].Process, nowPid)
  
  t := time.Now()
  fmt.Printf("shell end time:")
  fmt.Println(t.Unix());

  fmt.Printf("shell err:%v", err)
}
func main(){
   
   fmt.Println("进程id.", os.Getpid())  
   //time.Sleep(20 * time.Second)
   
   //ExecShell("sleep 60")
   //time.Sleep(10 * time.Second)
   //cancel()  
}