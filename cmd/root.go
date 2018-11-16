package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

// GitClient contains gitlab connection/project details
var GitClient *gitlab.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitlab-applet",
	Short: "facilitate with interactions with gitlab repos and api",
	Long: `Performs actions against specified gitlab api/url provided when
	supplied with credentials. Example:
	gitlab-applet show variables`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("token", "", "gitlab provided auth token")
	rootCmd.PersistentFlags().String("giturl", "", "enter gitlab url to use or query")
	rootCmd.PersistentFlags().String("project", "", "name of gitlab project/repo")
	rootCmd.MarkPersistentFlagRequired("token")
	rootCmd.MarkPersistentFlagRequired("giturl")
	rootCmd.MarkPersistentFlagRequired("project")

	cobra.OnInitialize(initConfig)
}

// initConfig sets up gitlab client
func initConfig() {
	token := rootCmd.Flag("token").Value.String()
	giturl := rootCmd.Flag("giturl").Value.String()
	GitClient = gitlab.NewClient(nil, token)
	GitClient.SetBaseURL(giturl + "/api/v4/")
}
