// +build windows darwin

package main

import (
	"os"
	"path/filepath"

	"github.com/concourse/atc"
	"github.com/concourse/baggageclaim/baggageclaimcmd"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
)

func (cmd *WorkerCommand) naiveBaggageclaimRunner(logger lager.Logger) (ifrit.Runner, error) {
	volumesDir := filepath.Join(cmd.WorkDir, "volumes")

	err := os.MkdirAll(volumesDir, 0755)
	if err != nil {
		return atc.Worker{}, nil, err
	}

	bc := &baggageclaimcmd.BaggageclaimCommand{
		BindIP:   baggageclaimcmd.IPFlag(cmd.Baggageclaim.BindIP),
		BindPort: baggageclaimcmd.IPFlag(cmd.Baggageclaim.BindPort),

		VolumesDir: volumesDir,

		Driver: "naive",

		ReapInterval: cmd.Baggageclaim.ReapInterval,

		Metrics: cmd.Metrics,
	}

	return bc.Runner(nil), nil
}
