package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hinha/reavers/provider"
)

// Run is a command to run web engine
type Run struct {
	engine provider.WebEngine
}

// NewRun return CLI to run web engine
func NewRun(engine provider.WebEngine) *Run {
	return &Run{engine: engine}
}

// Use return how the command used
func (r *Run) Use() string {
	return "run:web"
}

// Example of the command
func (r *Run) Example() string {
	return "run:web"
}

// Short description about the command
func (r *Run) Short() string {
	return "Run WEB engine"
}

// Run ...
func (r *Run) Run(args []string) {
	go func() {
		_ = r.engine.Run()
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 3 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// omit the error
	_ = r.engine.Shutdown(ctx)

	fmt.Println("\nGracefully shutdown the server...")
}
