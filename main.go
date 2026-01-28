package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

// Struktura pomocnicza dla pojedynczego backendu
type Backend struct {
	URL          *url.URL
	ReverseProxy *httputil.ReverseProxy
}

// ServerPool trzyma listę backendów i licznik
type ServerPool struct {
	backends []*Backend
	current  uint64
}

// Dodawanie nowego backendu do puli
func (s *ServerPool) AddBackend(serverUrl string) {
	u, err := url.Parse(serverUrl)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	// Modyfikacja odpowiedzi (opcjonalna - dla debugowania)
	proxy.ModifyResponse = func(r *http.Response) error {
		r.Header.Set("X-Proxy", "Equilibrium-RR")
		return nil
	}

	s.backends = append(s.backends, &Backend{
		URL:          u,
		ReverseProxy: proxy,
	})
}

// Algorytm Round-Robin: Zwraca następny serwer w kolejce
func (s *ServerPool) GetNextPeer() *Backend {
	// Atomic Add zapewnia bezpieczeństwo przy wielu wątkach (zwiększa licznik o 1)
	next := atomic.AddUint64(&s.current, 1)
	// Operacja modulo (%) pozwala wracać do początku listy (0, 1, 2, 0, 1, 2...)
	len := uint64(len(s.backends))
	idx := next % len
	return s.backends[idx]
}

func main() {
	// 1. Konfiguracja puli serwerów
	serverPool := &ServerPool{}

	// Tutaj definiujemy nasze "ofiary"
	serverPool.AddBackend("http://localhost:8081")
	serverPool.AddBackend("http://localhost:8082")
	serverPool.AddBackend("http://localhost:8083")

	// 2. Główny Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Wybieramy serwer algorytmem Round Robin
		peer := serverPool.GetNextPeer()

		if peer != nil {
			fmt.Printf("[Equilibrium] RR Index: %d -> Przekierowanie do %s\n", serverPool.current, peer.URL)
			peer.ReverseProxy.ServeHTTP(w, r)
		}
	})

	// 3. Start Load Balancera
	port := ":8000"
	log.Printf("Startowanie Load Balancera (Round Robin) na %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
