package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func init() {
	rootCmd.AddCommand(docsCmd)
}

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate Docs for this cli",
	Long: `
		Generate Docs for this cli.
		`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating docs")
		genDocs()
	},
}

func genDocs() {
	err := doc.GenMarkdownTree(rootCmd, "./docs")
	if err != nil {
		fmt.Println(err)
	}
}
