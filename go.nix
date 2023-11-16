{ pkgs ? import <nixpkgs> {} }:

with pkgs;
mkShell {
  nativeBuildInputs = [
    go
    gopls
    tmux
  ];
  shellHook = ''
    alias ll="ls -l"
    export FOO=bar
  '';
}
