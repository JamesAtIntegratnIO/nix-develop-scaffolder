{
  description = "Nix Develop Scaffolder";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.go_1_17
            pkgs.gotools
            pkgs.golangci-lint
            pkgs.gopls
            pkgs.go-outline
            pkgs.gopkgs
            pkgs.godef
            pkgs.gocode-gomod
            pkgs.gocode
          ];
        };
      });
}