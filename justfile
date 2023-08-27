# The latest tag in the local repository.
latest_git_tag := `git describe --tags --abbrev=0`
# The latest commit in the local repository.
latest_git_commit := `git rev-parse origin/main`

# When everything works as expected.
success_emoji := `printf "\\033[1;32m✔\\033[0m"`
# When something wrong happens.
error_emoji := `printf "\\033[0;31m✘\\033[0m"`

# Build the project in the build directory by default.
build dst='./build/main':
    @go build -ldflags="-X main.version={{latest_git_tag}} -X main.commit={{latest_git_commit}}" -o {{dst}} ./cmd/senv/

# Remove all build files.
clean:
    rm -rf ./result ./build/*

# Install the program in your $GOPATH/bin.
install: (build '$GOPATH/bin/senv')

# Run all the available Go tests in the source code.
test:
    @echo "Source code testing results:"
    @go clean -testcache && go test ./...

# Test and build the source code, if possible try to compile it with Nix.
test-suite: test
    @printf "Does it purely build? "
    @if just build; then echo "{{success_emoji}}"; else echo "{{error_emoji}}"; fi

    @printf "Does it builds with Nix? "
    @if command -v nix >/dev/null; then \
            if [[ $(nix show-config 2>/dev/null | grep "experimental-features") =~ "flakes" ]]; then \
                printf "\\033[38;2;209;214;114mflakes feature is not enabled\\033[0m\n"; \
            else \
                if nix build; \
                    then echo "Yeap it builds with Nix {{success_emoji}}"; \
                    else echo "Nix failed to build {{error_emoji}}"; \
                fi; \
            fi \
    else \
        printf "\\033[38;2;209;214;114m'nix' cannot be found in PATH\\033[0m\n"; \
    fi
