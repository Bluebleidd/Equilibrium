package proxy

import (
	"log"
	"net"
	"net/url"
	"time"
)

func (s *ServerPool) HealthCheck() {
	for _, b := range s.backends {
		alive := isBackendAlive(b.URL)
		b.SetAlive(alive)
		if !alive {
			log.Printf("Health Check failed: %s [DOWN]\n", b.URL)
		}
	}
}

func StartHealthCheckLoop(s *ServerPool, intervalStr string) {
	duration, err := time.ParseDuration(intervalStr)
	if err != nil {
		log.Printf("Config error: invalid duration, defaulting to 10s")
		duration = 10 * time.Second
	}

	t := time.NewTicker(duration)
	for range t.C {
		s.HealthCheck()
	}
}

func isBackendAlive(u *url.URL) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", u.Host, timeout)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
