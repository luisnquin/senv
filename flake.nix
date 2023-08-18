{
  description = "Switch your .env file from the command line";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {
          inherit system;
        };

        defaultPackage = pkgs.callPackage ./default.nix {};
      in {
        inherit defaultPackage;

        defaultApp = flake-utils.lib.mkApp {
          drv = defaultPackage;
        };

        devShell = pkgs.mkShell {
          buildInputs = [defaultPackage];
        };
      }
    );
}
