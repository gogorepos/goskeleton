package pubsub

import (
	"sync"
	"time"
)

// Publisher 发布者实体
type Publisher struct {
	m           sync.RWMutex             // 读写锁
	buffer      int                      // 订阅队列的缓存大小
	timeout     time.Duration            // 发布超时时间
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// NewPublisher 构建一个发布者，<t> 为超时时间，<buffer> 为缓存队列长度
func NewPublisher(t time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     t,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// SubscribeTopic 根据 <topic> 添加一个新的订阅者，如果 <topic> 为 nil 表示订阅全部主题
func (p *Publisher) SubscribeTopic(topic topicFunc) *subscriber {
	ch := make(subscriber, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return &ch
}

// SubscribeAll 添加一个新的订阅者，订阅全部主题
func (p *Publisher) SubscribeAll() *subscriber {
	return p.SubscribeTopic(nil)
}

// Evict 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// Publish 发布主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// Close 关闭发布者，同时关闭所有订阅者通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// sendTopic 向订阅者发送主题，直到超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}
