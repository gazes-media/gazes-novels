{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }: 
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          go_1_21
          gotools
          gopls
          xclip
          eza
        ];

        shellHook = ''
          export DB_HOST=""
          export DB_USER=""
          export DB_PASSWORD=""
          export DB_NAME=""
          export DB_PORT=""

          alias gdc='git diff --cached | xclip -selection clipboard'
        '';
      };
    };
}
