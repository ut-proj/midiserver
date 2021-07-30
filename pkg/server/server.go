package server

import (
	"context"
	"sync"
	"syscall"

	"github.com/geomyidia/erl-midi-server/internal/cli"
	"github.com/geomyidia/erl-midi-server/internal/util"
	"github.com/geomyidia/erl-midi-server/pkg/erl/exec"
	"github.com/geomyidia/erl-midi-server/pkg/erl/port"
	"github.com/geomyidia/erl-midi-server/pkg/types"
	log "github.com/sirupsen/logrus"
)

func Serve(ctx context.Context, key types.ParserKey, parserFlag string) {
	log.Info("Starting the server ...")
	ctx, cancel := util.SignalWithContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		parser := exec.ProcessMessage
		if parserFlag == cli.PortParser {
			parser = port.ProcessMessage
		}
		ProcessMessages(ctx, parser, ProcessCommand, key)
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	cancel()
	log.Info("Shutting down gracefully, press Ctrl+C again to force")
	log.Info("Waiting for wait groups to finish ...")
	wg.Wait()
	log.Info("Application shutdown complete.")
}
