package cmd

import (
	"fmt"
	"nix-project-generator/helpers"
	templateHandler "nix-project-generator/templateHandlers"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gitIgnoreCmd)
	gitIgnoreCmd.AddCommand(gitIgnoreAddCmd)
	gitIgnoreCmd.AddCommand(gitIgnoreListCmd)
}

var gitIgnoreCmd = &cobra.Command{
	Use:   "gitignore",
	Short: "Generates a gitignore file",
	Long: `
		Generates a gitignore file.
		`,
}

var gitIgnoreAddCmd = &cobra.Command{
	Use:   "add [gitignore key...]",
	Short: "Add a gitignore key to your project",
	Long: `
		Add a gitignore key to your project.
		The gitignore key is the key that is returned from the {gitignore list} command
		`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addGitIgnore(args)
	},
}

var gitIgnoreListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available gitignore keys",
	Long: `
		List all available gitignore keys. Fair warning, this is a long list.
		`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all available gitignore keys")
		listGitIgnoreOptions()
	},
}

func addGitIgnore(keys []string) {
	gitIgnore, err := templateHandler.FetchGitIgnoresBulk(keys)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gitIgnoreFile := "./.gitignore"
	err = helpers.WriteFile(gitIgnoreFile, gitIgnore)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func listGitIgnoreOptions() {
	list, err := templateHandler.FetchGitIgnoreList()
	if err != nil {
		fmt.Println(err)
	}
	chunkedList := helpers.ChunkSlice(list, 7)
	for _, chunk := range chunkedList {
		fmt.Printf("%v\n", chunk)
	}
}
