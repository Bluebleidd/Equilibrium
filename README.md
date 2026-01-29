# Equilibrium Load Balancer

![Go](https://img.shields.io/badge/Go-1.21-00ADD8?logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?logo=docker&logoColor=white)
![Status](https://img.shields.io/badge/Status-Active-success)

**Equilibrium** to lekki, wydajny Load Balancer warstwy 7 (HTTP) napisany w języku Go. Projekt demonstruje działanie systemów rozproszonych, w tym mechanizmy Reverse Proxy, algorytm Round-Robin oraz aktywne sprawdzanie stanu serwerów (Health Checks).

Całość działa w kontenerach Docker i posiada wbudowany **graficzny panel w terminalu (Dashboard)** do monitorowania ruchu na żywo.

---

## Do czego to służy? (Zastosowania)

Ten projekt jest gotowym szkieletem infrastruktury, który możesz wykorzystać do:

1.  **Zwiększenia stabilności aplikacji (High Availability):** Jeśli jeden serwer padnie, Equilibrium automatycznie przekieruje ruch do pozostałych.
2.  **Aktualizacji bez przestojów (Zero-Downtime Deployment):** Możesz aktualizować backendy pojedynczo, a użytkownicy nie zauważą przerwy w działaniu.
3.  **Skalowania ruchu:** Pozwala obsłużyć więcej użytkowników poprzez rozłożenie zapytań na wiele instancji tej samej aplikacji.
4.  **Testowania odporności (Chaos Engineering):** Idealne środowisko do nauki – możesz celowo wyłączać kontenery i obserwować, jak system radzi sobie z awariami.

---

## Wymagania Lokalne

Dzięki konteneryzacji, nie musisz instalować języka Go ani zależności na swoim komputerze.

* **Wymagane:** Zainstalowany [Docker Desktop](https://www.docker.com/products/docker-desktop).
* **Opcjonalnie:** Git (do sklonowania repozytorium).

---

## Konfiguracja

Główna konfiguracja znajduje się w pliku `config.json`. Możesz tam zmienić porty lub dodać więcej serwerów.

```json
{
  "port": ":8000",           // Port, na którym nasłuchuje Load Balancer
  "health_check_interval": "5s", // Jak często sprawdzać czy serwery żyją
  "backends": [              // Lista adresów Twoich aplikacji
    "http://backend1:8081",
    "http://backend2:8081",
    "http://backend3:8081"
  ]
}