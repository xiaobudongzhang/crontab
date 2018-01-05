package main
import (
  "fmt"
  "mycron/src/cron"
  "time"
  "mycron/src/lib"
  "encoding/json"
  "io/ioutil"
  "log"
)


const OneSecond = 1*time.Second
type MyJob struct{}

type Config struct {
    Id int
    Cron string
    Cmd string
    Process []int
}



func (d MyJob) Run() {
     
}


func main(){
      
       
         //解析json
	configData,err := ioutil.ReadFile("./config.json")
	if err!=nil{
	   fmt.Println("config read err",err.Error())
	   log.Fatal(err)
	}
	
	var configs []Config
	errjson := json.Unmarshal(configData, &configs)
	if errjson != nil {
	    fmt.Println("error:", err.Error())
	    log.Fatal(err)
	}
	


	lib.PrintNow("start-time")
	
       cron := cron.New()

	//loop
	for _,v := range configs{
	    //闭包
	    vtmp  := v
	    cron.AddFunc(v.Cron, func() {
		lib.ExecShell(vtmp.Cmd)
	    })
	}	     	
		 
	cron.Start()
	defer cron.Stop()
	


        //远程控制停止重启
	//time.Sleep(30 * time.Second)
        //cron.Stop()
	//time.Sleep(30 * time.Second)
	//cron.Start()
	
	lib.Server(cron);

	//最终超时兜底
	select {
	       case <-time.After(OneSecond*1200):
	       	 //   lib.PrintNow("end-time")
		   // fmt.Printf("main error")
	}

}
