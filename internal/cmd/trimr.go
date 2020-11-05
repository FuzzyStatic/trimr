package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cfgName = ".trimrconfig"
	cfgType = "yaml"
)

// Flags
var (
	FlagRepoPath string
)

type Option func(c *Trimr)

type (
	Trimr struct {
		// program information
		progName, version, buildTime, buildHost string
		// root command
		rootCmd *cobra.Command
		// slice of commands from opts
		cmds []*cobra.Command
		// protected branches read from TRIMR configuration
		protectedBranches []string
		// selected repository
		repo *git.Repository
	}
)

func NewTrimr(progName, version, buildTime, buildHost string, opts ...Option) (*Trimr, error) {
	var (
		t   Trimr
		err error
	)

	t.progName = progName
	t.version = version
	t.buildTime = buildTime
	t.buildHost = buildHost

	opts = append([]Option{withCmdVersion()}, opts...)

	// apply the list of options to Cmd
	for _, opt := range opts {
		opt(&t)
	}

	t.rootCmd = &cobra.Command{
		Use: t.progName,
		Run: t.trimr(),
	}

	for _, cmd := range t.cmds {
		t.rootCmd.AddCommand(cmd)
	}

	t.rootCmd.Flags().StringVarP(&FlagRepoPath, "path", "p", "", "path to the repository to trim (required)")
	_ = t.rootCmd.MarkFlagRequired("path")

	err = t.readInConfig()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Trimr) Execute() error {
	return t.rootCmd.Execute()
}

func (t *Trimr) trimr() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		t.trimRepo()
	}
}

func (t *Trimr) trimRepo() {
	fileInfo, err := os.Stat(FlagRepoPath)
	if os.IsNotExist(err) {
		log.Fatal("path does not exist")
	}

	if !fileInfo.IsDir() {
		log.Fatal("path is not a directory")
	}

	t.repo, err = git.PlainOpen(FlagRepoPath)
	if err != nil {
		log.Fatal(err)
	}

	err = t.removeBranches()
	if err != nil {
		log.Fatal(err)
	}
}

func (t *Trimr) removeBranches() error {
	refs, err := t.repo.Branches()
	if err != nil {
		return err
	}

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		var protected bool
		if ref.Name().IsBranch() {
			for _, protectedBranch := range t.protectedBranches {
				if strings.EqualFold(ref.Name().Short(), protectedBranch) {
					protected = true
				}
			}

			if !protected {
				// Prompt for branch deletion
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Delete branch \"%s\"", ref.Name().Short()),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					// Return if y is not input
					return nil
				}

				// Delete branch
				err = t.repo.Storer.RemoveReference(ref.Name())
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *Trimr) readInConfig() error {
	viper.SetConfigName(cfgName)
	viper.SetConfigType(cfgType)
	viper.AddConfigPath("../configs")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	t.readInProtectedBranches()

	return nil
}

func (t *Trimr) readInProtectedBranches() {
	protectedBranches := viper.Get("branches.protected").([]interface{})
	for _, protectedBranch := range protectedBranches {
		t.protectedBranches = append(t.protectedBranches, protectedBranch.(string))
	}
}
