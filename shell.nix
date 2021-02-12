with (import <nixpkgs> {});


mkShell {
  name = "go-shell";
  buildInputs = [
  pkgs.go
    
  ];

}
