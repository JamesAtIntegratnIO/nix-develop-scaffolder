package templateHandler

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type License struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

// Fetch license by name from https://api.github.com/licenses/${LICENSE}
// with header Accept: application/vnd.github.v3+json
// returns string
func FetchLicense(name string) (License, error) {
	var license License
	req, err := http.NewRequest("GET", "https://api.github.com/licenses/"+name, nil)
	if err != nil {
		return license, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return license, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return license, err
	}
	err = json.Unmarshal(body, &license)
	if err != nil {
		return license, err
	}

	return license, nil
}

// Populate license template with name and year
// returns string
func PopulateLicenseTemplate(lname string, name string, year string) (string, error) {
	license, err := FetchLicense(lname)
	if err != nil {
		return "", err
	}
	license.Body = strings.ReplaceAll(license.Body, "[year]", year)
	license.Body = strings.ReplaceAll(license.Body, "[fullname]", name)

	t, err := template.New("license").Parse(license.Body)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	err = t.Execute(&buffer, struct {
		Year     string
		FullName string
	}{
		Year:     year,
		FullName: name,
	})
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// Fetch License types from https://api.github.com/licenses
func FetchLicenses() ([]License, error) {
	var licenses []License
	req, err := http.NewRequest("GET", "https://api.github.com/licenses", nil)
	if err != nil {
		return licenses, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return licenses, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return licenses, err
	}
	err = json.Unmarshal(body, &licenses)
	if err != nil {
		return licenses, err
	}

	return licenses, nil
}
