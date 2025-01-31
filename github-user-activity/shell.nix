{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    gopls
    go
  ];
  shellHook = ''
    source ~/.bashrc
  '';
}
