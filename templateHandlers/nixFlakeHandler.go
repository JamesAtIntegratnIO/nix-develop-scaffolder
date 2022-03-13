package templateHandler

import (
	"bytes"
	"fmt"
	"html/template"
	"nix-project-generator/langs"
	"os"
)

func GenerateNixFlake(l []langs.Lang, data interface{}) string {
	t, err := template.ParseFiles("templates/nixFlake.tmpl")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(l) > 0 {
		t.New("buildInputs").Parse(langs.ExecuteMultipleBuildInputsToString(l))
	} else {
		fmt.Println("No language Provided. Generating flake.nix with empty build inputs")
		t.New("buildInputs").Parse("# no lang provided")
	}

	var buffer bytes.Buffer
	err = t.Execute(&buffer, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return buffer.String()

}
