package main

import (
	"log"
	"net/http"
	"os"

	"equilibrium/internal/config"
	"equilibrium/internal/proxy"
	"equilibrium/internal/ui"
	_ "expvar"
)

func main() {

	go func() {
		log.Println("Metryki expvar dostępne pod adresem: http://localhost:2112/debug/vars")
		http.ListenAndServe(":2112", nil)
	}()

	cfg := config.LoadConfig("config.json")

	serverPool := proxy.NewServerPool(cfg.Strategy)

	for _, backendUrl := range cfg.Backends {
		serverPool.AddBackend(backendUrl)
	}

	go proxy.StartHealthCheckLoop(serverPool, cfg.HealthCheckInterval)

	server := http.Server{
		Addr: cfg.Port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			peer := serverPool.GetNextPeer(r)

			if peer != nil {
				peer.ReverseProxy.ServeHTTP(w, r)
			} else {
				http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			}
		}),
	}

	log.SetOutput(os.Stderr)

	go ui.StartDashboard(serverPool, cfg.Port, cfg.Strategy)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
