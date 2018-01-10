package main
import (
  "fmt"
  "cron"
  "time"
  "lib"
)


const OneSecond = 1*time.Second + 10*time.Millisecond

func main(){
  //test
  cron := cron.New()
  cron.AddFunc("0 * * * * *", func() {
    lib.ExecShell("date")
  })


	cron.Start()
	defer cron.Stop()

	// Give cron 2 seconds to run our job (which is always activated).
	select {
	case <-time.After(OneSecond*600):
		 fmt.Printf("error")
	}

}
