{
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  version = "v0.5.0";
  commit = "c1d04585d83ec229fa4ce2033c45ba41af6f9981";
in
  buildGoModule rec {
    pname = "senv";
    inherit version;

    src = builtins.path {
      name = "senv-switcher";
      path = ./.;
    };

    vendorSha256 = "sha256-qTgdWQKCkNCOxFLM77wMvbPvwHOkBOPzRP8a10hZuqw=";
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
