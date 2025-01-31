{ pkgs ? import <nixpkgs> {} }:
{
  deno = pkgs.callPackage ./deno.nix { };
}
