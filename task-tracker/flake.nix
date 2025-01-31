{
  description = "A nix flake that packages the task-cli command";
  inputs = {
    nixpkgs = {
      url = "github:nixos/nixpkgs?ref=nixos-unstable";
    };
    deno = {
      type = "github";
      owner = "denoland";
      repo = "deno";
      flake = false;
    };
  };
  outputs = { self, nixpkgs, deno }: {
    formatter.x86_64-linux = nixpkgs.legacyPackages.x86_64-linux.nixpkgs-fmt;
    packages.x86_64-linux.default = with import nixpkgs { system = "x86_64-linux"; };
      stdenv.mkDerivation {
        name = "task-cli";
        src = self;
        buildPhase = "deno build app";
        installPhase = "mkdir -p $out/bin; install -t $out/bin task-cli";
      };
  };
}
