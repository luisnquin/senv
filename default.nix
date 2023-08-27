{
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  version = "v0.5.2";
  commit = "28e7f0baf4dd7ff1232ba98143e157d2c435faf5";
in
  buildGoModule rec {
    pname = "senv";
    inherit version;

    src = builtins.path {
      name = "senv-switcher";
      path = ./.;
    };

    vendorSha256 = "sha256-B6e1U8eDxXaB+3Skt/DxwWlF/33jJ07F+RT1ABCFiSw=";
    doCheck = true;

    buildTarget = "./cmd/senv";
    ldflags = ["-X main.version=${version} -X main.commit=${commit}"];

    meta = with lib; {
      description = "Switch your .env file from the command line";
      homepage = "https://github.com/luisnquin/${pname}";
      license = licenses.mit;
      maintainers = with maintainers; [luisnquin];
    };
  }
