package cmd

import (
	"os"
	"runtime/pprof"

	"github.com/PicoTools/plan/pkg/engine"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type App struct {
	CpuProfile bool
}

func (g *App) RegisterFlags(f *pflag.FlagSet) {
	f.BoolVar(&g.CpuProfile, "cpuprofile", false, "")
}

func (a *App) Run(cmd *cobra.Command, args []string) error {
	if a.CpuProfile {
		f, err := os.Create("cpu.pprof")
		if err != nil {
			return err
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			return err
		}
		defer pprof.StopCPUProfile()
	}
	err := engine.Evaluate(args[0])
	if err != nil {
		return err
	}
	return nil
}
