{
  installShellFiles,
  buildGoModule,
  lib,
}: let
  version = "0.8.1";
  commit = "856437ef35f194199276f38f0601715764cffe83";
in
  buildGoModule {
    pname = "senv";
    inherit version;

    src = builtins.path {
      name = "senv-switcher";
      path = ./.;
    };

    vendorHash = "sha256-GtFvRGUkmh639zRi/V2sSuVhcHzQf1I0g4IXLuht2Lg=";
    doCheck = true;

    buildTarget = ".";
    ldflags = ["-X main.version=v${version} -X main.commit=${commit}"];

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
