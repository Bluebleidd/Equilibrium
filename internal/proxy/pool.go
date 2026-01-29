package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type ServerPool struct {
	backends []*Backend
	current  uint64
}

func NewServerPool() *ServerPool {
	return &ServerPool{}
}

func (s *ServerPool) AddBackend(serverUrl string) {
	u, err := url.Parse(serverUrl)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		log.Printf("[%s] %s\n", u.Host, e.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service Unavailable"))
	}

	s.backends = append(s.backends, &Backend{
		URL:          u,
		ReverseProxy: proxy,
		Alive:        true,
	})
}

func (s *ServerPool) GetNextPeer() *Backend {
	next := s.current
	l := uint64(len(s.backends))

	for i := uint64(0); i < l; i++ {
		next = atomic.AddUint64(&s.current, 1)
		idx := next % l

		if s.backends[idx].IsAlive() {
			if i != 0 {
				atomic.StoreUint64(&s.current, next)
			}
			s.backends[idx].IncrementRequests()
			return s.backends[idx]
		}
	}
	return nil
}
func (s *ServerPool) GetBackends() []*Backend {
	return s.backends
}
