{
  installShellFiles,
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  version = "v0.8.0";
  commit = "7ec50b6552c458fb3392945f4e3fb43bfb914a77";
in
  buildGoModule {
    pname = "senv";
    inherit version;

    src = builtins.path {
      name = "senv-switcher";
      path = ./.;
    };

    vendorSha256 = "sha256-GtFvRGUkmh639zRi/V2sSuVhcHzQf1I0g4IXLuht2Lg=";
    doCheck = true;

    buildTarget = ".";
    ldflags = ["-X main.version=${version} -X main.commit=${commit}"];

    nativeBuildInputs = [
      installShellFiles
    ];

    postInstall = ''
      installShellCompletion --cmd senv \
        --zsh <($out/bin/senv completion zsh)
    '';

    meta = with lib; {
      description = "Switch your .env file from the command line";
      homepage = "https://github.com/luisnquin/senv";
      license = licenses.mit;
      maintainers = with maintainers; [luisnquin];
    };
  }
