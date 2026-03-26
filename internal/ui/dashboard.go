package ui

import (
	"equilibrium/internal/proxy"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"

	ClearScreen = "\033[2J"
	MoveTopLeft = "\033[H"
)

func Render(pool *proxy.ServerPool, port string, strategy string) {
	fmt.Print(ClearScreen + MoveTopLeft)

	line := "====================================================="

	fmt.Println(ColorCyan + line + ColorReset)
	fmt.Printf("   EQUILIBRIUM LOAD BALANCER (Port: %s)   \n", port)
	fmt.Printf("   STRATEGY: %s \n", strings.ToUpper(strategy))
	fmt.Println(ColorCyan + line + ColorReset)
	fmt.Println("")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

	fmt.Fprintln(w, "SERVER\tSTATUS\tACTIVE\tTOTAL REQ")
	fmt.Fprintln(w, "------\t------\t------\t---------")

	for _, b := range pool.GetBackends() {
		statusText := "DOWN"
		color := ColorRed

		if b.IsAlive() {
			statusText = "UP"
			color = ColorGreen
		}

		fmt.Fprintf(w, "%s\t%s%s%s\t%d\t%d\n",
			b.URL,
			color, statusText, ColorReset,
			b.GetActiveConnections(),
			b.GetRequests(),
		)
	}

	w.Flush()

	fmt.Println("\n" + ColorCyan + line + ColorReset)
	fmt.Println("[Ctrl+C] Stop System | [Ctrl+P, Ctrl+Q] Detach")
}

func StartDashboard(pool *proxy.ServerPool, port string, strategy string) {
	for {
		Render(pool, port, strategy)
		time.Sleep(500 * time.Millisecond)
	}
}
