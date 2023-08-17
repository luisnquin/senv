{
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  version = "v100";
  commit = "e86581e201799a819a6fca77935df9e7d2a207b9";
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
