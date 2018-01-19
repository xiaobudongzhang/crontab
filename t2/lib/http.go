package lib

import (
  "net/http"
  "fmt"
  "log"
  "html"
  
  "time"
   "encoding/json"
   //"net/url"
   //"strconv"
   //"os"
   //"context"
   //"os/exec"
   //"bytes"
   "io/ioutil"
)



const OneSecond = 1*time.Second

func Server(){

  http.HandleFunc("/addTask", func(res http.ResponseWriter, req *http.Request) {
      body,err := ioutil.ReadAll(req.Body)
      if err != nil{
      	 panic(err)
      } 
      
      var config DataConfig
      
      err = json.Unmarshal(body, &config)

      if err != nil {
      	 panic(err);
      }
      //log.Println(config.StartTime)

      NewRw(config);
      fmt.Fprintf(res, "%s", html.EscapeString("200"))	
  })

  log.Fatal(http.ListenAndServe("169.254.87.130:8080", nil))
}

