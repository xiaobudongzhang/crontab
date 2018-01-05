package lib

import (
  "net/http"
  "fmt"
  "log"
  "html"
  "mycron/src/cron"
)


func Server(c *cron.Cron){
  http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
     c.Start()
     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
    c.Stop()
     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })


  log.Fatal(http.ListenAndServe("192.168.80.129:8080", nil))
}