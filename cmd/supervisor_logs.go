package cmd

import (
	"fmt"

	helper "github.com/home-assistant/cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var supervisorLogsCmd = &cobra.Command{
	Use:     "logs",
	Aliases: []string{"log", "lg"},
	Short:   "View the log output of the Home Assistant Supervisor",
	Long: `
Allowing you to look at the log output generated by the Home Assistant
Supervisor running on your Home Assistant system.`,
	Example: `
  ha supervisor logs`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("args", args).Debug("supervisor logs")

		section := "supervisor"
		command := "logs"

		url, err := helper.URLHelper(section, command)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		request := helper.GetRequest()
		resp, err := request.SetHeader("Accept", "text/plain").Get(url)

		if err != nil {
			fmt.Println(err)
			ExitWithError = true
		} else {
			fmt.Println(resp.String())
		}
	},
}

func init() {
	supervisorCmd.AddCommand(supervisorLogsCmd)
}
