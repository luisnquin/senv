{
  installShellFiles,
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  version = "v0.6.0";
  commit = "0142cc2fa61b85e3364b7512753732e8d69a6921";
in
  buildGoModule {
    pname = "senv";
    inherit version;

    src = builtins.path {
      name = "senv-switcher";
      path = ./.;
    };

    vendorSha256 = "sha256-YDzN1WUzhSHWvgxc7JKo5kIiQJDKOFYtvZQFXSl6ztU=";
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
