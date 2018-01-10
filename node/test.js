var schedule = require('node-schedule');
const { spawn } = require('child_process');
const fs = require('fs');

//执行命令，将结果放入到日志文件
fs.readFile('config.json','utf8',function (err, data) {
    if(err) console.log(err);
    var config=JSON.parse(data);
    if (!config) {
        console.log("config parse error")
    }
    run(config);
});

function run(config) {
    config.forEach((val) => {
            var j = schedule.scheduleJob(val['cron'], function(){

                const ls = spawn(val['cmd']);

                ls.stdout.on('data', (data) => {
                    console.log(`stdout: ${data}`);
                });

                ls.stderr.on('data', (data) => {
                    console.log(`stderr: ${data}`);
                });

                ls.on('close', (code) => {
                    console.log(`child process exited with code ${code}`);
                });
          });
    });
}

