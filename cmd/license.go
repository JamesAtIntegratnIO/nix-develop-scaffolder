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
	Use:   "add [license key] [name] [year]",
	Short: "Add a license to your project",
	Long: `
		Add a license to your project.
		The license key is the key that is returned from the {license list} command
		`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		name := args[1]
		year := strconv.Itoa(time.Now().Year())
		if len(args) > 2 {
			year = args[2]
		}
		fmt.Println("Adding license:", key)
		createLicenseFile(key, name, year)
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

func createLicenseFile(key string, name string, year string) {
	fmt.Println("Creating license file")
	license, err := templateHandler.PopulateLicenseTemplate(key, name, year)
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
