{
  installShellFiles,
  buildGoModule,
  lib,
}: let
  version = "1.1.1";
  commit = "cc1d60c86c0e9ba5e8ff3ac99475520f6dc1397c";
in
  buildGoModule {
    pname = "senv";
    inherit version;

    src = builtins.path {
      name = "senv-switcher";
      path = ./.;
    };

    vendorHash = "sha256-FTxvzV7Gu1OYIV2RAWNX/ciVHZhrpK4aQpnXGuTarTI=";
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
