{
  description = "Exchain Nix Flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        inherit (pkgs) rustPlatform go;
      in
      rec {
        exchain = pkgs.stdenv.mkDerivation rec {
          pname = "exchain";
          version = "1.6.8.5";

          src = self;

          nativeBuildInputs = [ go rustPlatform.cargo ];
          buildInputs = [ pkgs.rocksdb ];

          buildPhase = ''
            make exchain
          '';

          installPhase = ''
            mkdir -p $out/bin
            cp cmd/exchaind/exchaind $out/bin/
            cp cmd/exchaincli/exchaincli $out/bin/
          '';

          meta = with pkgs.lib; {
            description = "Exchain blockchain node and CLI";
            homepage = "https://github.com/okx/exchain";
            license = licenses.asl20;
            maintainers = [ ];
            platforms = platforms.all;
          };
        };

        defaultPackage = exchain;
      });
}
