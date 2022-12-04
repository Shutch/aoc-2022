{
  description = "Demonstrating The Pipelines Pattern In Golang";

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
            pkgs.bashInteractive
            pkgs.go
            pkgs.go-tools
            pkgs.golangci-lint
            pkgs.gopls
            pkgs.go-outline
            pkgs.gopkgs
            pkgs.delve
            pkgs.gocode
            pkgs.gocode-gomod
            pkgs.godef
            pkgs.gotests
          ];
          hardeningDisable = [ "fortify" ];
        };
      });
}
