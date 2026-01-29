package ui

import (
	"equilibrium/internal/proxy"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"

	ClearScreen = "\033[2J"
	MoveTopLeft = "\033[H"
)

func Render(pool *proxy.ServerPool, port string) {
	fmt.Print(ClearScreen + MoveTopLeft)

	line := "============================================="

	fmt.Println(ColorCyan + line + ColorReset)
	fmt.Printf("   EQUILIBRIUM LOAD BALANCER (Port: %s)   \n", port)
	fmt.Println(ColorCyan + line + ColorReset)
	fmt.Println("")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "SERVER\tSTATUS      REQUESTS")
	fmt.Fprintln(w, "------\t------      --------")

	for _, b := range pool.GetBackends() {
		statusText := "DOWN"
		color := ColorRed

		if b.IsAlive() {
			statusText = "UP"
			color = ColorGreen
		}
		paddedStatus := fmt.Sprintf("%-12s", statusText)

		row := fmt.Sprintf("%s\t%s%s%s%d",
			b.URL,
			color, paddedStatus, ColorReset,
			b.GetRequests(),
		)

		fmt.Fprintln(w, row)
	}

	w.Flush()

	fmt.Println("\n" + ColorCyan + line + ColorReset)
}

func StartDashboard(pool *proxy.ServerPool, port string) {
	for {
		Render(pool, port)
		time.Sleep(500 * time.Millisecond)
	}
}
