{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/c33c7c3d5fa99fb13d39fd9af265b38a161abaf1.tar.gz") {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
  ];
}
