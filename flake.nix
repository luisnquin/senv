{
  description = "Switch your .env file from the command line";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    pkgs = import nixpkgs {
      inherit system;
    };

    system = "x86_64-linux";
    app = pkgs.callPackage ./default.nix {};
  in {
    devShell = pkgs.mkShell {
      buildInputs = [app];
    };

    defaultPackage.${system} = app;

    packages.${system}."senv" = app;
  };
}
