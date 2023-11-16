{ pkgs ? import <nixpkgs> {} }:

with pkgs;
mkShell {
  nativeBuildInputs = [
    go
    gopls
    tmux
  ];
}
