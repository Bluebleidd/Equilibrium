package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ServerPool struct {
	backends []*Backend
	Strategy BalancingStrategy
}

func NewServerPool(strategyName string) *ServerPool {
	pool := &ServerPool{}

	switch strategyName {
	case "least_conn":
		pool.Strategy = &LeastConnStrategy{}
		log.Println("Używana strategia: Least Connections")
	case "ip_hash":
		pool.Strategy = &IPHashStrategy{}
		log.Println("Używana strategia: IP Hash")
	default:
		pool.Strategy = &RoundRobinStrategy{}
		log.Println("Używana strategia: Round Robin (Domyślna)")
	}

	return pool
}

func (s *ServerPool) AddBackend(serverUrl string) {
	u, err := url.Parse(serverUrl)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(u)

	backend := &Backend{
		URL:          u,
		ReverseProxy: proxy,
		Alive:        true,
	}

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		backend.DecrementActiveConns()
		return nil
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		log.Printf("[%s] %s\n", u.Host, e.Error())
		backend.DecrementActiveConns()
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service Unavailable"))
	}

	s.backends = append(s.backends, backend)
}

func (s *ServerPool) GetNextPeer(r *http.Request) *Backend {
	if s.Strategy == nil {
		return nil
	}

	backend := s.Strategy.NextBackend(s.backends, r)
	if backend != nil {
		backend.IncrementRequests()
		backend.IncrementActiveConns()
	}
	return backend
}

func (s *ServerPool) GetBackends() []*Backend {
	return s.backends
}
