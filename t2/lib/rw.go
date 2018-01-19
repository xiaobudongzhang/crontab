package lib

import (
      "time"
      "log"
      //"strconv"
)

func NewRw(config DataConfig){
    
   
   now := time.Now().Unix()
   //nowint := strconv.FormatInt(now, 10) 
   end := config.StartTime
   
   spend := (config.StartTime - now)*time.Second
   
   time.AfterFunc(spend,func(){
   	log.Println("config.Id")
   })

}



