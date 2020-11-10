package cmd

import (
	"github.com/robocorp/rcc/common"
	"github.com/robocorp/rcc/operations"

	"github.com/spf13/cobra"
)

func createWorkarea() {
	if len(directory) == 0 {
		common.Exit(1, "Error: missing target directory")
	}
	err := operations.InitializeWorkarea(directory, templateName, forceFlag)
	if err != nil {
		common.Exit(2, "Error: %v", err)
	}
}

func listTemplates() {
	common.Log("Template names:")
	for _, name := range operations.ListTemplates() {
		common.Log("- %v", name)
	}
}

var initializeCmd = &cobra.Command{
	Use:     "initialize",
	Aliases: []string{"init"},
	Short:   "Create a directory structure for a robot.",
	Long:    "Create a directory structure for a robot.",
	Run: func(cmd *cobra.Command, args []string) {
		if common.Debug {
			defer common.Stopwatch("Initialization lasted").Report()
		}
		if listFlag {
			listTemplates()
		} else {
			createWorkarea()
		}
		common.Log("OK.")
	},
}

func init() {
	robotCmd.AddCommand(initializeCmd)
	initializeCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Root directory to create the new robot in.")
	initializeCmd.Flags().StringVarP(&templateName, "template", "t", "standard", "Template to use to generate the robot content.")
	initializeCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force the creation of the robot and possibly overwrite data.")
	initializeCmd.Flags().BoolVarP(&listFlag, "list", "l", false, "List available templates.")
}
