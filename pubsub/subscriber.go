package pubsub

type (
	subscriber chan interface{} // 订阅者通道
)

// Listen 监听通道信息
func (s *subscriber) Listen(handler func(v interface{})) {
	for v := range *s {
		handler(v)
	}
}
