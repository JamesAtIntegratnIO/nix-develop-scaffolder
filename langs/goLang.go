package langs

var GoLang = Lang{
	Name:    "go",
	Version: "1.16",
	BuildInputs: `  pkgs.go_{{.Version}}
	    pkgs.gotools
	    pkgs.golangci-lint
	    pkgs.gopls
	    pkgs.go-outline
	    pkgs.gopkgs
	`,
}
