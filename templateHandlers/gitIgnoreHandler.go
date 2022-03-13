package templateHandler

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type GitIgnores struct {
	Ignores []string
}

type IDE struct {
	Name    string
	Enabled bool
}

var Vscode = IDE{
	Name:    "visualstudiocode",
	Enabled: false,
}

var Vim = IDE{
	Name:    "vim",
	Enabled: false,
}

var Emacs = IDE{
	Name:    "emacs",
	Enabled: false,
}

func FetchGitIgnore(name string) (string, error) {
	resp, err := http.Get("https://www.toptal.com/developers/gitignore/api/" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func FetchGitIgnores(names []string) (gitIgnores GitIgnores, err error) {
	for _, name := range names {
		gitIgnore, err := FetchGitIgnore(name)
		if err != nil {
			return gitIgnores, err
		}
		gitIgnores.Ignores = append(gitIgnores.Ignores, gitIgnore)

	}
	return gitIgnores, nil
}

func FetchGitIgnoresBulk(names []string) (gitIgnore string, err error) {
	keys := strings.Join(names, ",")
	gitIgnore, err = FetchGitIgnore(keys)
	if err != nil {
		return gitIgnore, err
	}

	return gitIgnore, nil
}

func GenerateGitIgnoreFromTemplate(names []string) string {
	IDEs := []IDE{Vscode, Vim, Emacs}
	for _, ide := range IDEs {
		if ide.Enabled {
			names = append(names, ide.Name)
		}
	}
	gitIgnores, err := FetchGitIgnores(names)
	if err != nil {
		panic(err)
	}
	t, err := template.ParseFiles("templates/gitIgnore.tmpl")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", gitIgnores)
	var buffer bytes.Buffer
	t.Execute(&buffer, gitIgnores)
	return buffer.String()
}

// Get gitignore list from https://www.toptal.com/developers/gitignore/api/list
func FetchGitIgnoreList() ([]string, error) {
	resp, err := http.Get("https://www.toptal.com/developers/gitignore/api/list?format=lines")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(body), "\n"), nil
}
