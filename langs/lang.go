package langs

import (
	"bytes"
	"html/template"
)

type Lang struct {
	Name        string
	Version     string
	BuildInputs string
}

func ExecuteTemplateToString(templateString string, data interface{}) string {
	t := template.New("")
	t, _ = t.Parse(templateString)
	var buffer bytes.Buffer
	t.Execute(&buffer, data)
	return buffer.String()
}

func ExecuteBuildInputsToString(lang Lang) string {
	return ExecuteTemplateToString(lang.BuildInputs, lang)
}

func ExecuteMultipleBuildInputsToString(langs []Lang) string {
	var buffer bytes.Buffer
	for _, lang := range langs {
		buffer.WriteString(ExecuteBuildInputsToString(lang))
	}
	return buffer.String()
}
