package hpool

//import (
//	"context"
//	"errors"
//	"github.com/jhinih/hin/hinterface"
//	"sync"
//	"time"
//)
//
//// 定义接口，方便 mock 和替换
//type ClientPool interface {
//	Get() (hinterface.IClient, error)
//	Put(client hinterface.IClient)
//	Close()
//	Stats() (total, idle, active int)
//}
//
//// 配置项
//type PoolConfig struct {
//	InitialSize         int           // 初始连接数
//	MaxSize             int           // 最大连接数
//	IdleTimeout         time.Duration // 空闲超时（超过则关闭）
//	HealthCheckInterval time.Duration // 健康检查周期
//	RetryInterval       time.Duration // 重连间隔
//}
//
//// 增强型 Client 包装，带上元数据
//type pooledClient struct {
//	hinterface.IClient
//	lastUsedAt time.Time
//	isHealthy  bool
//	mu         sync.RWMutex
//}
//
//// 真正的连接池
//type AdvancedClientPool struct {
//	config    *PoolConfig
//	factory   func() hinterface.IClient   // Client 创建工厂
//	idleQueue chan *pooledClient          // 空闲队列（用带缓冲 channel）
//	active    map[hinterface.IClient]bool // 正在使用的连接（用 map 防泄漏）
//	mu        sync.RWMutex
//	ctx       context.Context
//	cancel    context.CancelFunc
//	wg        sync.WaitGroup
//	stats     struct {
//		total, idle, active int
//	}
//}
//
//// 创建高级池
//func NewAdvancedClientPool(config *PoolConfig, serverAddr string) ClientPool {
//	ctx, cancel := context.WithCancel(context.Background())
//
//	pool := &AdvancedClientPool{
//		config:    config,
//		factory:   func() hinterface.IClient { return znet.NewClient(serverAddr, 0) },
//		idleQueue: make(chan *pooledClient, config.MaxSize),
//		active:    make(map[hinterface.IClient]bool),
//		ctx:       ctx,
//		cancel:    cancel,
//	}
//
//	// 预热：创建初始连接
//	for i := 0; i < config.InitialSize; i++ {
//		if client := pool.createClient(); client != nil {
//			pool.idleQueue <- client
//		}
//	}
//
//	// 启动后台协程：健康检查 + 自动补充
//	pool.wg.Add(2)
//	go pool.healthChecker()
//	go pool.maintainer()
//
//	return pool
//}
//
//// 创建并连接 Client
//func (p *AdvancedClientPool) createClient() *pooledClient {
//	client := p.factory()
//	client.Start() // 开始连接（异步）
//
//	// 等待连接成功或失败（简单版，实际可配超时）
//	time.Sleep(100 * time.Millisecond)
//
//	return &pooledClient{
//		IClient:    client,
//		lastUsedAt: time.Now(),
//		isHealthy:  true, // 假设启动后就是健康的
//	}
//}
//
//// 获取连接（核心逻辑）
//func (p *AdvancedClientPool) Get() (hinterface.IClient, error) {
//	select {
//	case <-p.ctx.Done():
//		return nil, errors.New("pool closed")
//
//	case client := <-p.idleQueue: // 从空闲队列拿
//		client.mu.Lock()
//		client.lastUsedAt = time.Now()
//		client.isHealthy = true // 重置健康状态
//		client.mu.Unlock()
//
//		p.mu.Lock()
//		p.active[client.IClient] = true // 标记为活跃
//		p.mu.Unlock()
//
//		return client, nil
//
//	default: // 池子空了，但还能新建
//		p.mu.RLock()
//		activeCount := len(p.active)
//		p.mu.RUnlock()
//
//		if activeCount < p.config.MaxSize {
//			return p.createClient(), nil
//		}
//
//		// 超过最大连接数，阻塞等待
//		select {
//		case client := <-p.idleQueue:
//			return client, nil
//		case <-time.After(5 * time.Second):
//			return nil, errors.New("get client timeout")
//		}
//	}
//}
//
//// 归还连接（核心逻辑）
//func (p *AdvancedClientPool) Put(client hinterface.IClient) {
//	p.mu.Lock()
//	delete(p.active, client) // 从活跃 map 移除
//	p.mu.Unlock()
//
//	// 检查连接是否还健康
//	if pc, ok := client.(*pooledClient); ok {
//		pc.mu.RLock()
//		healthy := pc.isHealthy
//		pc.mu.RUnlock()
//
//		if !healthy {
//			client.Stop() // 不健康就销毁
//			return
//		}
//
//		select {
//		case p.idleQueue <- pc: // 塞回空闲队列
//		default: // 池子满了，直接关闭
//			client.Stop()
//		}
//	}
//}
//
//// 健康检查协程：定期检查连接是否还活着
//func (p *AdvancedClientPool) healthChecker() {
//	defer p.wg.Done()
//
//	ticker := time.NewTicker(p.config.HealthCheckInterval)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-p.ctx.Done():
//			return
//		case <-ticker.C:
//			p.checkConnections()
//		}
//	}
//}
//
//// 实际检查每个空闲连接
//func (p *AdvancedClientPool) checkConnections() {
//	// 复制当前空闲连接进行检查
//	tempQueue := make([]*pooledClient, 0)
//
//CLOSE:
//	for {
//		select {
//		case client := <-p.idleQueue:
//			// 简单健康检查：看连接是否断开
//			if conn := client.GetTCPConn(); conn != nil {
//				// 尝试读一个字节看是否 eof（实际可优化）
//				conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
//				buf := make([]byte, 1)
//				_, err := conn.Read(buf)
//				if err != nil { // 连接已断开
//					client.Stop()
//					continue
//				}
//			}
//
//			// 检查空闲超时
//			if time.Since(client.lastUsedAt) > p.config.IdleTimeout {
//				client.Stop()
//				continue
//			}
//
//			tempQueue = append(tempQueue, client)
//
//		default:
//			break CLOSE
//		}
//	}
//
//	// 把健康的重新塞回池子
//	for _, c := range tempQueue {
//		p.idleQueue <- c
//	}
//}
//
//// 维护者协程：保持最小连接数
//func (p *AdvancedClientPool) maintainer() {
//	defer p.wg.Done()
//
//	ticker := time.NewTicker(p.config.RetryInterval)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-p.ctx.Done():
//			return
//		case <-ticker.C:
//			// 若空闲连接数 < InitialSize，则补充
//			currentIdle := len(p.idleQueue)
//			for i := currentIdle; i < p.config.InitialSize; i++ {
//				if client := p.createClient(); client != nil {
//					select {
//					case p.idleQueue <- client:
//					default:
//						client.Stop()
//					}
//				}
//			}
//		}
//	}
//}
//
//// 关闭整个池子
//func (p *AdvancedClientPool) Close() {
//	p.cancel()
//
//	// 关闭所有活跃连接
//	p.mu.Lock()
//	for client := range p.active {
//		client.Stop()
//	}
//	p.mu.Unlock()
//
//	// 关闭所有空闲连接
//CLOSE:
//	for {
//		select {
//		case client := <-p.idleQueue:
//			client.Stop()
//		default:
//			break CLOSE
//		}
//	}
//
//	p.wg.Wait()
//}
//
//// 统计信息
//func (p *AdvancedClientPool) Stats() (total, idle, active int) {
//	p.mu.RLock()
//	active = len(p.active)
//	p.mu.RUnlock()
//
//	idle = len(p.idleQueue)
//	total = active + idle
//	return
//}
