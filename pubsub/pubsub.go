// Package pubsub 提供了发布者-订阅者模式
package pubsub

type (
	topicFunc func(v interface{}) bool // 主题过滤器
)
