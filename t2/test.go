package main
import (
  //"fmt"
  //"mycron/src/cron"
  //"time"
  "mycron2/src/lib"
  //"encoding/json"
  //"io/ioutil"
  //"log"
)
func main(){	

	lib.PrintNow("start-time")
		
        //远程控制
	lib.Server();

	//最终超时兜底
	//select {
	   
	//}

}


