{ pkgs ? import <nixpkgs> { } }:

let
  rustPlatform = pkgs.rustPlatform;
  go = pkgs.go;
in
pkgs.stdenv.mkDerivation rec {
  pname = "exchain";
  version = "1.6.8.5";

  src = ./.;

  nativeBuildInputs = [ go rustPlatform.cargoBuildHook ];
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
}
