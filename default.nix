{
  installShellFiles,
  buildGoModule,
  lib,
}: let
  version = "1.1.0";
  commit = "89fb9098443a779e152cb8ac2fa8696904eea479";
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
