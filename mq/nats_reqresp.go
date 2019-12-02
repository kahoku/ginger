package mq

import (
	"context"
	"time"
)

/*
Request()是一个简单方便的API，它提供了一个伪同步的方式，使用了超时timeout配置。它创建了一个收件箱reply（收件箱是一种主题(subject) 类型，
对请求者唯一），订阅主题(subject)，然后发布你的请求消息（消息带reply地址）,设置为收件箱的主题(subject)，然后等待响应，或者超时取消。
@param subject string			请求的主题
@param v interface{}			发送的消息
@param replyPtr interface{}		伪同步收件箱，在超时时间内接收同步响应数据
@param timeout time.Duration	等待响应的超时时间

For example：

1.Requests 发布一个主题等待响应消息
var response string
err = c.Request("help", "help me", &response, 10*time.Millisecond)
if err != nil {
    fmt.Printf("Request failed: %v\n", err)
}

2.Replying 订阅一个request主题，并向收件箱reply发布响应信息
c.Subscribe("help", func(subj, reply string, msg string) {
    c.Publish(reply, "I can help!")
})
*/
func (mq *NatsMQ) Request(subject string, v interface{}, replyPtr interface{}, timeout time.Duration) error {
	err := mq.conn.Request(subject, v, &replyPtr, timeout)
	if err != nil {
		return err
	}
	return err
}

/*
PublishRequest将执行一个Publish()操作，要求对回复主题做出响应。使用Request()自动内联等待响应。
*/
func (mq *NatsMQ) PublicRequest(subject, reply string, msg interface{}) error {
	err := mq.conn.PublishRequest(subject, reply, msg)
	return err
}

/*
RequestWithContext将创建一个收件箱，并使用提供的取消上下文和数据v的收件箱回复执行请求。响应将被解码为vPtrResponse。

*/
func (mq *NatsMQ) RequestWithContext(ctx context.Context, subject string, msg interface{}, respPtr interface{}, timeout time.Duration) error {
	err := mq.conn.RequestWithContext(ctx, subject, msg, &respPtr)
	if err != nil {
		return err
	}
	return err
}