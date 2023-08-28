
latest_git_tag := `git describe --tags --abbrev=0`
latest_git_commit := `git rev-parse origin/main`

user := `echo $USER`

build dst='./build/main':
    @go build -ldflags="-X main.version={{latest_git_tag}} -X main.commit={{latest_git_commit}}" -o {{dst}} ./cmd/senv/

install: (build '/home/$USER/go/bin/senv')

clean:
    rm -rf ./result