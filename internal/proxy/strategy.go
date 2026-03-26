package proxy

import (
	"hash/fnv"
	"math"
	"net/http"
	"sync/atomic"
)

type BalancingStrategy interface {
	NextBackend(backends []*Backend, r *http.Request) *Backend
}

type RoundRobinStrategy struct {
	current uint64
}

func (s *RoundRobinStrategy) NextBackend(backends []*Backend, r *http.Request) *Backend {
	next := atomic.AddUint64(&s.current, 1)
	l := uint64(len(backends))

	for i := uint64(0); i < l; i++ {
		idx := next % l
		if backends[idx].IsAlive() {
			return backends[idx]
		}
		next = atomic.AddUint64(&s.current, 1)
	}
	return nil
}

type LeastConnStrategy struct{}

func (s *LeastConnStrategy) NextBackend(backends []*Backend, r *http.Request) *Backend {
	var leastConnBackend *Backend
	minConns := int64(math.MaxInt64)

	for _, b := range backends {
		if b.IsAlive() {
			conns := b.GetActiveConnections()
			if conns < minConns {
				minConns = conns
				leastConnBackend = b
			}
		}
	}
	return leastConnBackend
}

type IPHashStrategy struct{}

func (s *IPHashStrategy) NextBackend(backends []*Backend, r *http.Request) *Backend {
	if len(backends) == 0 {
		return nil
	}

	clientIP := r.RemoteAddr

	h := fnv.New32a()
	h.Write([]byte(clientIP))
	hashValue := h.Sum32()

	idx := int(hashValue) % len(backends)

	startIdx := idx
	for {
		if backends[idx].IsAlive() {
			return backends[idx]
		}
		idx = (idx + 1) % len(backends)
		if idx == startIdx {
			break
		}
	}
	return nil
}
