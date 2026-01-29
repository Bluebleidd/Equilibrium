package proxy

import (
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	URL          *url.URL
	ReverseProxy *httputil.ReverseProxy
	Alive        bool
	mux          sync.RWMutex
	Requests     uint64
}

func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

func (b *Backend) IsAlive() bool {
	b.mux.RLock()
	defer b.mux.RUnlock()
	return b.Alive
}

func (b *Backend) GetRequests() uint64 {
	return atomic.LoadUint64(&b.Requests)
}

func (b *Backend) IncrementRequests() {
	atomic.AddUint64(&b.Requests, 1)
}
