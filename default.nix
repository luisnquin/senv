{
  installShellFiles,
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  version = "v0.7.0";
  commit = "d7a4eaafe77f24049c1799cb330f53bfd4197e4c";
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
