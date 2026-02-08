# Equilibrium Load Balancer

**Equilibrium** is a high-performance Layer 7 (HTTP) Load Balancer engineered in Go. It acts as a reverse proxy that distributes incoming network traffic across multiple backend servers to ensure efficiency and reliability.

The system is fully containerized using Docker and features a **Real-time Terminal Dashboard (TUI)**. This dashboard provides instant visibility into traffic throughput and server health statuses without the need for complex logging setups. Core capabilities include:
* **Round-Robin Scheduling:** Evenly distributes requests across healthy instances.
* **Active Health Checks:** Periodically pings backends and automatically removes unresponsive servers from the rotation.
* **Fault Tolerance:** Prevents HTTP 500 errors by routing traffic only to "UP" servers.

## Use Cases
Equilibrium is designed as a foundational infrastructure component suitable for various scenarios:

* **High Availability (HA):** Ensures your service remains accessible even if one or more backend server instances crash.
* **Zero-Downtime Deployments:** Facilitates "Rolling Updates" where backend services are updated one by one. The load balancer routes traffic away from the instance being updated, ensuring users experience no service interruption.
* **Horizontal Scaling:** Allows you to handle increased traffic loads by simply adding more container instances to the backend pool.
* **Chaos Engineering:** Provides a safe local environment to simulate server failures (e.g., stopping a Docker container) and observe system recovery mechanisms in real-time.

## Requirements
To run this project, you do not need to install Go or manual dependencies. The entire environment is containerized.

**Prerequisites:**
* **Docker Desktop** (includes Docker Engine and Docker Compose).
* **Git** (optional, for cloning the repository).

*No local installation of Golang is required.*

## Configuration & Custom Integration
This section explains how to configure the load balancer and integrate your own custom application (e.g., Python, Node.js, Java) instead of the default test application.

## Core Configuration (`config.json`)
The `config.json` file in the root directory controls the load balancer's behavior.

```json
{
  "port": ":8000",
  "backends": [
    "http://backend1:8081",
    "http://backend2:8081",
    "http://backend3:8081"
  ],
  "health_check_interval": "5s"
}
