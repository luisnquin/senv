{
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  version = "v0.5.1";
  commit = "d1eaa3e287f5b619424cc6cdaf1e62791d49a277";
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
