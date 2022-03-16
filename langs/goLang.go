package langs

var GoLang = Lang{
	Name: "go",
	BuildInputs: `  pkgs.go_{{.Version}}
	    pkgs.gotools
	    pkgs.golangci-lint
	    pkgs.gopls
	    pkgs.go-outline
	    pkgs.gopkgs
	`,
}
