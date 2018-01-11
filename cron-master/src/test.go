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
      //test
       cron := cron.New()
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
	
	//loop
	//for _,v := range configs{
	   // fmt.Printf(v.Cmd)
	    //cron.AddFunc(v.Cron, func() {

	    	lib.ExecShell("sleep 40")	    			
	    //})
	//}	     	
	//start
        t := time.Now()
	fmt.Printf("start-time:")	 
	fmt.Println(t.Unix());
		 
	cron.Start()
	defer cron.Stop()
	
	
	select {
	       case <-time.After(OneSecond*60*100):
		    fmt.Printf("error")
	}

}
