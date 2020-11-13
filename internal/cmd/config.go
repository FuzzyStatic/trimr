package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cfgName = ".trimrconfig.yaml"
	cfgType = "yaml"
)

// withCmdConfig option inserts command to get current version
func withCmdConfig() Option {
	return func(t *Trimr) {
		t.cmds = append(t.cmds, t.newCmdCfg())
	}
}

func (t *Trimr) newCmdCfg() *cobra.Command {
	cmdCfg := &cobra.Command{
		Use:     "config",
		Aliases: []string{"cfg"},
		Short:   "Manage configuration file",
		Long:    `Manage the trimr configuration file.`,
		Args:    cobra.ExactValidArgs(1),
	}

	cmdPB := &cobra.Command{
		Use:     "protected-branch",
		Aliases: []string{"pb"},
		Short:   "Manage protected branches",
		Long:    `Manage the protected branches in the trimr configuration file.`,
		Args:    cobra.ExactValidArgs(1),
	}

	cmdPBAdd := &cobra.Command{
		Use:     "add",
		Aliases: []string{"a"},
		Short:   "Add protected branch",
		Long:    `Add protected branch the trimr configuration file.`,
		Args:    cobra.NoArgs,
		Run:     t.runCmdPBAdd(),
	}

	cmdPBAdd.Flags().StringVarP(&FlagBranchName, "name", "n", "", "name of branch to protect (require)")
	_ = cmdPBAdd.MarkFlagRequired("name")

	cmdPBList := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "List protected branches",
		Long:    `List protected branches from the trimr configuration file.`,
		Args:    cobra.NoArgs,
		Run:     t.runCmdPBList(),
	}

	cmdPBRemove := &cobra.Command{
		Use:     "remove",
		Aliases: []string{"rm"},
		Short:   "Remove protected branch",
		Long:    `Remove protected branch the trimr configuration file.`,
		Args:    cobra.NoArgs,
		Run:     t.runCmdPBRemove(),
	}

	cmdPBRemove.Flags().StringVarP(&FlagBranchName, "name", "n", "", "name of branch to unprotect (required)")
	_ = cmdPBRemove.MarkFlagRequired("name")

	cmdPB.AddCommand(cmdPBList)
	cmdPB.AddCommand(cmdPBAdd)
	cmdPB.AddCommand(cmdPBRemove)
	cmdCfg.AddCommand(cmdPB)

	return cmdCfg
}

func (t *Trimr) runCmdPBAdd() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := t.readInConfig()
		if err != nil {
			panic(err)
		}
		t.addProtectedBranch()
	}
}

func (t *Trimr) runCmdPBList() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := t.readInConfig()
		if err != nil {
			panic(err)
		}
		t.listProtectedBranches()
	}
}

func (t *Trimr) runCmdPBRemove() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := t.readInConfig()
		if err != nil {
			panic(err)
		}
		t.removeProtectedBranch()
	}
}

func (t *Trimr) readInConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName(cfgName)
	viper.SetConfigType(cfgType)
	viper.AddConfigPath(homeDir)
	err = viper.ReadInConfig()
	if err != nil {
		createDefaultConfig(homeDir)
	}

	t.readInProtectedBranches()

	return nil
}

func (t *Trimr) readInProtectedBranches() {
	t.protectedBranches = viper.GetStringSlice("branches.protected")
}

func (t *Trimr) addProtectedBranch() {
	viper.Set("branches.protected", append(viper.GetStringSlice("branches.protected"), FlagBranchName))
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}

func (t *Trimr) listProtectedBranches() {
	pb := viper.GetStringSlice("branches.protected")

	for _, v := range pb {
		fmt.Println(v)
	}
}

func (t *Trimr) removeProtectedBranch() {
	pb := viper.GetStringSlice("branches.protected")
	for i, branch := range pb {
		if strings.EqualFold(branch, FlagBranchName) {
			pb[i] = pb[len(pb)-1] // Copy last element to index i.
			pb[len(pb)-1] = ""    // Erase last element (write zero value).
			pb = pb[:len(pb)-1]
		}
	}

	viper.Set("branches.protected", pb)
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}

func createDefaultConfig(homeDir string) {
	viper.SetDefault("title", "TRIMR Configuration")
	viper.SetDefault("branches.protected", []string{"main", "master"})
	err := viper.WriteConfigAs(fmt.Sprintf("%s/.trimrconfig.yaml", homeDir))
	if err != nil {
		panic(err)
	}
}
