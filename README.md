# ⚖️ Equilibrium Load Balancer

![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)

**Equilibrium** is a lightweight, concurrent Layer 7 (HTTP) Load Balancer built from scratch in Go. It demonstrates core distributed systems concepts including Reverse Proxying, Round-Robin scheduling, and Active Health Checks.

The project runs entirely in Docker and features a built-in **Real-time Terminal Dashboard (TUI)** to monitor traffic distribution and server health instantly.

---

## 🚀 Use Cases

This project serves as a production-ready infrastructure skeleton for:
1.  **High Availability:** Automatically routes traffic away from crashed servers.
2.  **Zero-Downtime Deployments:** Allows updating backend services one by one without disrupting users.
3.  **Horizontal Scaling:** Distributes heavy traffic load across multiple application instances.
4.  **Chaos Engineering:** Perfect for testing how your system behaves when components fail.

---

## 🛠️ Architecture

```mermaid
graph LR
    Client(User / Browser) -->|HTTP :8000| LB{Equilibrium LB}
    
    subgraph "Docker Network"
        LB -->|Round Robin| App1(Backend 1)
        LB -->|Round Robin| App2(Backend 2)
        LB -->|Round Robin| App3(Backend 3)
    end

    LB -.->|Health Check Loop| App1
    LB -.->|Health Check Loop| App2
    LB -.->|Health Check Loop| App3