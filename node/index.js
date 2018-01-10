const fs = require('fs');
var parser = require('cron-parser');

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
        try {
            var interval = parser.parseExpression(val['cron']);



            while (true ) {
                try {
                    var obj = interval.next();
                    var nexttime = obj.toString()
                    var nexttimest = Date.parse(nexttime)
                    nexttimest = nexttimest/1000;

                    var timestamp = Date.parse(new Date());
                    var now = timestamp / 1000;

                    console.log('nexttimest:', nexttimest, 'now:', now);
                    if ( nexttimest == now)
                    {
                        //console.log('value:', obj.toString(), 'done:', obj.done);
                    }
                } catch (e) {
                    console.log('Error1: ' + e.message);
                    break;
                }
            }
        } catch (err) {
            console.log('Error2: ' + err.message);
        }
    });

}