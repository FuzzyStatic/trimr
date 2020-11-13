package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// withCmdVersion option inserts command to get current version
func withCmdVersion() Option {
	return func(t *Trimr) {
		t.cmds = append(t.cmds, t.newCmdVer())
	}
}

func (t *Trimr) newCmdVer() *cobra.Command {
	cmdVer := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Print version",
		Long:    `Prints the current version of this application.`,
		Args:    cobra.NoArgs,
		Run:     t.runCmdVer(),
	}

	return cmdVer
}

func (t *Trimr) runCmdVer() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		t.printVersion()
	}
}

func (t *Trimr) printVersion() {
	fmt.Printf("%s version %s\n", t.progName, t.version)
}
