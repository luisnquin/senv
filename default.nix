{
  installShellFiles,
  buildGoModule,
  lib,
}: let
  version = "0.8.2";
  commit = "5a1d7df33cdefc91287c720b53c6a416e3f722ea";
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
