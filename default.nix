{
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  version = "v0.4.5";
  commit = "c0ca2dfd25cbb102331c75ecdd8761bbff9fcecc";
in
  buildGoModule rec {
    pname = "senv";
    inherit version;

    src = ./.;

    vendorSha256 = "sha256-C33Kj6PXoXa3OuH1ZP5kDJGR+BNaqbDrDGNtVpYgHZU=";
    doCheck = true;

    buildTarget = ".";
    ldflags = ["-X main.version=${version} -X main.commit=${commit}"];

    meta = with lib; {
      description = "Switch your .env file from the command line";
      homepage = "https://github.com/luisnquin/${pname}";
      license = licenses.mit;
      maintainers = with maintainers; [luisnquin];
    };
  }
