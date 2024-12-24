package main

import (
	"context"
	"os"

	"github.com/PicoTools/plan/cmd/plan/internal/cmd"
	"github.com/PicoTools/plan/internal/zapcfg"
	"github.com/fatih/color"
	"github.com/go-faster/sdk/zctx"
	"github.com/spf13/cobra"
)

func main() {
	var err error

	lg, err := zapcfg.New().Build()
	if err != nil {
		panic(err)
	}

	flush := func() {
		_ = lg.Sync()
	}

	exit := func(code int) {
		flush()
		os.Exit(code)
	}

	a := cmd.App{}
	ctx := zctx.Base(context.Background(), lg)

	root := &cobra.Command{
		Use:           "plan [flags] <file>",
		Short:         "PLAN runtime executor",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.MinimumNArgs(1),
		RunE:          a.Run,
		PersistentPostRun: func(*cobra.Command, []string) {
			flush()
		},
	}

	root.CompletionOptions.DisableDefaultCmd = true
	a.RegisterFlags(root.PersistentFlags())

	if err = root.ExecuteContext(ctx); err != nil {
		color.Red("%v", err)
		exit(2)
	}
}
