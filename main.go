package main

import (
	"log"
	"net/http"
	"os"

	"equilibrium/internal/config"
	"equilibrium/internal/proxy"
	"equilibrium/internal/ui"
)

func main() {
	cfg := config.LoadConfig("config.json")

	serverPool := proxy.NewServerPool()
	for _, backendUrl := range cfg.Backends {
		serverPool.AddBackend(backendUrl)
	}

	go proxy.StartHealthCheckLoop(serverPool, cfg.HealthCheckInterval)

	server := http.Server{
		Addr: cfg.Port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			peer := serverPool.GetNextPeer()
			if peer != nil {
				peer.ReverseProxy.ServeHTTP(w, r)
			} else {
				http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			}
		}),
	}

	log.SetOutput(os.Stderr)

	go ui.StartDashboard(serverPool, cfg.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
