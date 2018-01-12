package lib

import (
  "net/http"
  "fmt"
  "log"
  "html"
  "mycron/src/cron"
  "time"
   "encoding/json"
   "net/url"
   "strconv"
   "os"
   //"context"
   "os/exec"
   //"bytes"
)

type Config struct {
    Id int
    Cron string
    Cmd string
    Process []int
}


const OneSecond = 1*time.Second

func Server(c *cron.Cron){
  http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
     c.Start()
   fmt.Fprintf(w, "%s", html.EscapeString("200"))
  })

  http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
    c.Stop()
 fmt.Fprintf(w, "%s", html.EscapeString("200"))
  })


  http.HandleFunc("/setConfig", func(w http.ResponseWriter, r *http.Request) {
    
   fmt.Fprintf(w, "%s", html.EscapeString("200"))
    ReLoad(c)    
     
  })
  
  http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
  		fmt.Fprintf(w, "%s", html.EscapeString("200"))
  })

  http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "%s", html.EscapeString("201"))
  })

  http.HandleFunc("/processFromCmd", func(w http.ResponseWriter, req *http.Request) {
  		query, err := url.ParseQuery(req.URL.RawQuery)
		 if err != nil {
                    fmt.Println("error:", err)
                }

		 id ,_:= strconv.Atoi(query["id"][0])
		 
		 
 		if _,ok := DataConfigs[id];!ok{
		     fmt.Println("error:", err)
		//     DataConfigsLess := DataConfigs
		}
		
		DataConfigsLess := DataConfigs[id].Process
		
		
  		b, err := json.Marshal(DataConfigsLess)
		if err != nil {
		    fmt.Println("error:", err)
		}

                fmt.Fprintf(w, "%s", string(b))
  })

  http.HandleFunc("/tree", func(w http.ResponseWriter, r *http.Request) {
        ppid :=os.Getpid()
	
  //ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  //defer cancel()
  //cmd := exec.CommandContext(ctx, "pstree", "-p", strconv.Itoa(ppid))
 cmd := exec.Command("pstree", "-p", strconv.Itoa(ppid))

  //var out bytes.Buffer
  //cmd.Stdout = &out


  //err := cmd.Start()
   //if err !=nil{
    //fmt.Println("error:", err)
  //}

  outdata, err := cmd.Output()
  //err = cmd.Wait()


  if err != nil {
       fmt.Println("error tree:", err)
  }
  //fmt.Printf("%s\n",outdata)
  fmt.Fprintf(w, "%s", outdata)

  })


  log.Fatal(http.ListenAndServe("169.254.87.130:8080", nil))
}



func ReLoad(c *cron.Cron){
    
        str := []byte(`[{"Id": 11,"Cron": "*/5 * * * * *","Cmd": "sleep 33","Process": []}]`)
     var configs []Config
     json.Unmarshal(str, &configs)
    
  	     
        //loop
        for _,v := range configs{
            //闭包
            vtmp  := v
            c.AddFunc(v.Cron, func() {
                ExecShell(vtmp.Cmd,vtmp.Id)
            })
        }
	
}
