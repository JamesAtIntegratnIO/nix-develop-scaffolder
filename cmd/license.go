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
	licenseCmd.AddCommand(licenseAddCmd)
	licenseCmd.AddCommand(licensesListCmd)
}

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Fetch a list of licenses or add it to your project",
}

var licenseAddCmd = &cobra.Command{
	Use:   "add [license] [name] [year]",
	Short: "Add a license to your project",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		license := args[0]
		name := args[1]
		year := strconv.Itoa(time.Now().Year())
		if len(args) > 2 {
			year = args[2]
		} else {

		}
		fmt.Println("Adding license:", license)
		createLicenseFile(license, name, year)
	},
}

var licensesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available licenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all available licenses")
		listLicenses()
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

func listLicenses() {
	licenses, err := templateHandler.FetchLicenses()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, license := range licenses {
		fmt.Printf("License Key: %-20s License Name: %08s , \n", license.Key, license.Name)
	}
}
