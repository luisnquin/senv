{
  fetchFromGitHub,
  buildGoModule,
  lib,
}: let
  owner = "luisnquin";
  version = "0.4.4";
  commit = "0ab42428cd5e689160a7592496d11288bf756441";
in
  buildGoModule rec {
    pname = "senv";
    inherit version;

    src = fetchFromGitHub {
      inherit owner;

      repo = pname;
      rev = "v${version}";
      sha256 = "100q65248sc32qrmx5385p0cw7ixyxaw3a6wa2rnq4kaahwmhsyf";
    };

    vendorSha256 = "sha256-C33Kj6PXoXa3OuH1ZP5kDJGR+BNaqbDrDGNtVpYgHZU=";
    doCheck = true;

    buildTarget = ".";
    ldflags = ["-X main.version=${version} -X main.commit=${commit}"];

    meta = with lib; {
      description = "Switch your .env file from the command line";
      homepage = "https://github.com/${owner}/${pname}";
      license = licenses.mit;
      maintainers = with maintainers; [luisnquin];
    };
  }
