package cmd

import (
	"fmt"
	"strings"

	"nix-project-generator/helpers"
	"nix-project-generator/langs"
	templateHandler "nix-project-generator/templateHandlers"

	"github.com/spf13/cobra"
)

type Project struct {
	// Langs       []langs.Lang
	Name        string
	Description string
	Path        string
}

var Langs []langs.Lang

var P = Project{}

func init() {
	rootCmd.AddCommand(create)
	create.Flags().StringVarP(&langs.GoLang.Version, "go", "g", "", "go version")
	fmt.Println(langs.GoLang.Version)

}

var create = &cobra.Command{
	Use:   "create [project_name] [descrition] [path]",
	Short: "Creates a new go project with a flake.",
	Long:  "Creates a new go project with a flake.\n Use the language flags in order to add languages and versions to the flake.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		P.Name = args[0]
		if len(args) > 1 {
			P.Description = args[1]
		} else {
			P.Description = "My Dumb Description that I didn't write"
		}
		if len(args) > 2 {
			P.Path = args[2]
		} else {
			P.Path = "./" + P.Name
		}

		P.createProject()
	},
}

func (p *Project) createProject() {
	fmt.Println("Creating project:", p.Name)
	fmt.Println(langs.GoLang.Version)
	if langs.GoLang.Version != "" {
		langs.GoLang.Version = strings.Replace(langs.GoLang.Version, ".", "_", -1)
		Langs = append(Langs, langs.GoLang)
	}
	p.writeProjectFiles()
}

// create makefile from template
func (p *Project) writeProjectFiles() {
	helpers.MakeFolderIfNotExists(p.Path)

	nixFlake := templateHandler.GenerateNixFlake(Langs, p)

	nixFlakeFile := p.Path + "/flake.nix"
	err := helpers.WriteFile(nixFlakeFile, nixFlake)
	if err != nil {
		fmt.Println(err)
	}
}
