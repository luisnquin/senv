{
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  version = "v0.4.5";
  commit = "0ab42428cd5e689160a7592496d11288bf756441";
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
