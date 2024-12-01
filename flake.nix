{
  inputs = {
    nixpkgs-unstable.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs-unstable,
    utils,
    ...
  }:
    utils.lib.eachDefaultSystem (
      system: let
        overlays = [];

        pkgs = import nixpkgs-unstable {
          inherit system overlays;

          config = {};
        };
      in {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            delve
          ];
        };
      }
    );
}
