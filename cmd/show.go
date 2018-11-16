package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var supportedArgs = []string{"variables", "test"}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show information regarding a project",
	Long: `used to get information about a project. Could be commits, pull requests,
project variables, etc. 

gitlab-applet show variables`,

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New(color.RedString("requires exactly one argument"))
		}

		argOk := Contains(supportedArgs, args[0])
		if !argOk {
			return errors.New(color.RedString(`arg: "%s" not implemented
supported arguments: %v
			`, args[0], supportedArgs))
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("SHOW command invoked....")
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
	color.Cyan("fetching variables for project: %s ...", project.WebURL)

	varList, _, err := GitClient.ProjectVariables.ListVariables(project.ID)
	if err != nil {
		log.Fatalln(err)
	}

	if len(varList) == 0 {
		color.Yellow("No project variables found")
		return
	}
	color.Green("variables fetched!\n*******")

	for _, v := range varList {
		// a lot of our variables are base64 encoded
		data, err := base64.StdEncoding.DecodeString(v.Value)
		if err != nil {
			// if err, data was not base 64, but still grab
			fmt.Println("Key: ", v.Key)
			fmt.Println("Value: ", v.Value)
			color.Green("*******")
			continue
		}

		fmt.Println("Key: ", v.Key)
		fmt.Printf("Value: \n%s\n", string(data))
		color.Green("*******")
	}
}

func init() {
	rootCmd.AddCommand(showCmd)
}
