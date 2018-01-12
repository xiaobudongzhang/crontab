package lib

type DataProcess struct{
     Pid int
}


type DataConfig struct {
    Id int
    Cron string
    Cmd string
    Process map[int] DataProcess 
}
//var DataProcess map[int] DataProcessBase

var DataConfigs map[int] DataConfig
   //DataConfigs = make(map[string] DataConfig, 200)
 