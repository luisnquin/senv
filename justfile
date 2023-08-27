
latest_git_tag := `git describe --tags --abbrev=0`
latest_git_commit := `git rev-parse origin/main`

success_emoji := `printf "\\033[1;32m✔\\033[0m"`
error_emoji := `printf "\\033[0;31m✘\\033[0m"`

build dst='./build/main':
    @go build -ldflags="-X main.version={{latest_git_tag}} -X main.commit={{latest_git_commit}}" -o {{dst}} ./cmd/senv/

clean:
    rm -rf ./result ./build/*

install: (build '/home/$USER/go/bin/senv')

test:
    @echo "Source code testing results:"
    @go clean -testcache && go test ./...

test-suite: test
    @printf "Does it purely build? "
    @if just build; then echo "{{success_emoji}}"; else echo "{{error_emoji}}"; fi

    @printf "Does it builds with Nix?"
    @if command -v nix >/dev/null; \
        then if nix build; then echo "Yeap it builds with Nix {{success_emoji}}" ; else echo "Nix failed to build {{error_emoji}}"; fi \
    fi
