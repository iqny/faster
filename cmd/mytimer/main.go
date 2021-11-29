package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds()) //创建一个cron实例
	//c := cron.New()
	//执行定时任务（每5秒执行一次）
	c.AddFunc("*/1 * * * * *", print5)

	//启动/关闭
	c.Start()
	defer c.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}
}

//执行函数
func print5() {
	fmt.Println("每5s执行一次cron")
}
