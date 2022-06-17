package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func doConsumerTask() {
	// 1. 创建消费者
	config := nsq.NewConfig()
	q, errNewCsmr := nsq.NewConsumer("topic_demo", "first", config)
	if errNewCsmr != nil {
		fmt.Printf("fail to new consumer!, topic=%s, channel=%s", "topic_demo", "first")
		return
	}

	// 2. 添加处理消息的方法
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		//todo 将数据写入到ES （为方便演示，消费端只做打印）
		log.Printf("time: %s, message: %v", time.Now().Format("2006-01-02 15:04:05"), string(message.Body))
		message.Finish()
		return nil
	}))

	// 3. 通过http请求来发现nsqd生产者和配置的topic（推荐使用这种方式）
	lookupAddr := []string{
		"10.20.23.107:4161",
	}
	err := q.ConnectToNSQLookupds(lookupAddr)
	if err != nil {
		log.Panic("[ConnectToNSQLookupds] Could not find nsqd!")
	}

	// 4. 接收消费者停止通知
	<-q.StopChan

	// 5. 获取统计结果
	stats := q.Stats()
	fmt.Println(fmt.Sprintf("message received %d, finished %d, requeued:%s, connections:%s",
		stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections))
}

func main() {
	doConsumerTask()
}
