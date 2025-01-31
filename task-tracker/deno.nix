{
  stdenv,
  fetchFromGitHub
}:
stdenv.mkDerivation {
  pname = "deno";
  version = "2.1.9";

  src = fetchFromGitHub {
    owner = "denoland";
    repo = "deno";
    rev = "v2.1.9";
    sha256 = "0wk3zzp1qnpjq25chpadnxb7825rhbk6jb1ys41ikx9cj88hqykb";
  };
}
