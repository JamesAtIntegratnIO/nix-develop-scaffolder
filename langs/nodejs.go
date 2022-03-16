package langs

var NodeJS = Lang{
	Name: "node",
	BuildInputs: ` pkgs.nodejs_{{.Version}}_x
		pkgs.yarn
	`,
}
