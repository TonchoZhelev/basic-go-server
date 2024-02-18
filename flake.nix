{
  description = "A basic Go server";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    systems.url = "github:nix-systems/default";
    flake-parts.url = "github:hercules-ci/flake-parts";
    treefmt-nix.url = "github:numtide/treefmt-nix";
    treefmt-nix.inputs.nixpkgs.follows = "nixpkgs";
    fourmolu-nix.url = "github:jedimahdi/fourmolu-nix";
  };

  outputs = inputs:
    inputs.flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;
      imports = [
        inputs.treefmt-nix.flakeModule
        inputs.fourmolu-nix.flakeModule
      ];
      perSystem = { self', system, lib, config, pkgs, ... }: {
        # Default shell.
        devShells.default = pkgs.mkShell {
          buildInputs = [ pkgs.go ];
        };
      };
  };
}

