package lib


type DataConfig struct {
    Id int
    StartTime int64
    Sleep int
    MaxCount int
    CallUrl string
    CallData string
    FailCount int
}

var DataConfigs map[int] DataConfig
 