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





func (d MyJob) Run() {
     
}


func main(){
      
       
         //解析json
	configData,err := ioutil.ReadFile("./config.json")
	if err!=nil{
	   fmt.Println("config read err",err.Error())
	   log.Fatal(err)
	}
	
	var configs = lib.DataConfigs
	errjson := json.Unmarshal(configData, &configs)
	if errjson != nil {
	    fmt.Println("error:", err.Error())
	    log.Fatal(err)
	}
	lib.DataConfigs = configs
	

	lib.PrintNow("start-time")
	
        cron := cron.New()
       
	//loop
	for _,v := range configs{
	    //闭包
	    vtmp  := v
	    cron.AddFunc(v.Cron, func() {
		lib.ExecShell(vtmp.Cmd,vtmp.Id)
	    })
	}
	
		 
	cron.Start()
	defer cron.Stop()

	
        //远程控制
	lib.Server(cron);

	//最终超时兜底
	select {
	       case <-time.After(OneSecond*1200):
	       	 //   lib.PrintNow("end-time")
		   // fmt.Printf("main error")
	}

}
