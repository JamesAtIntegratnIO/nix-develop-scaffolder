package cmd

import (
	"fmt"
	"nix-project-generator/helpers"
	templateHandler "nix-project-generator/templateHandlers"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(licenseCmd)
}

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Fetch a license and add it to your project",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		license := args[0]
		name := args[1]
		year := strconv.Itoa(time.Now().Year())
		if len(args) > 2 {
			year = args[2]
		} else {

		}
		fmt.Println("Fetching license:", license)
		createLicenseFile(license, name, year)
	},
}

func createLicenseFile(l string, name string, year string) {
	fmt.Println("Creating license file")
	license, err := templateHandler.PopulateLicenseTemplate(l, name, year)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = helpers.WriteFile("./LICENSE", license)
	if err != nil {
		fmt.Println(err)
		return
	}
}
