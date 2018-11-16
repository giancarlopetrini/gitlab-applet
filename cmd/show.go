package cmd

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show information regarding a project",
	Long: `used to get information about a project. Could be commits, pull requests,
project variables, etc. 

gitlab-applet show --variables`,

	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("Show CMD called....")
		switch args[0] {
		case "variables":
			variables()
		}
	},
}

func variables() {
	opt := &gitlab.ListProjectsOptions{
		Search:  gitlab.String(rootCmd.Flag("project").Value.String()),
		OrderBy: gitlab.String("name"),
		Sort:    gitlab.String("asc"),
	}
	projects, _, err := GitClient.Projects.ListProjects(opt)
	if err != nil {
		log.Fatalln(err)
	}

	if len(projects) == 0 {
		log.Fatalln(color.RedString("No matching projects found for: %s", rootCmd.Flag("project").Value))
	}

	project := projects[0]
	color.Green("fetching variables for project: %s ...", project.WebURL)

	varList, _, err := GitClient.ProjectVariables.ListVariables(project.ID)
	if err != nil {
		log.Fatalln(err)
	}

	if len(varList) == 0 {
		color.Yellow("No project variables found")
		return
	}

	for _, v := range varList {
		data, err := base64.StdEncoding.DecodeString(v.Value)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("----%v-----\n", v.Key)
		fmt.Println(string(data))
	}
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().String("item", "", "desired information (variables, commits, etc)")
}
