package proxy

import (
	"expvar"
)

var (
	TotalRequests     = expvar.NewInt("equilibrium_requests_total")
	ActiveConnections = expvar.NewInt("equilibrium_active_connections")
)
