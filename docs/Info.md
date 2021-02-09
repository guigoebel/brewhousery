# Setup

+ Install Go

  `brew install go`

+ Install Go tools

  `go get -u golang.org/x/tools/...`

+ Install `jq` to format `curl` output

  `brew install jq`

+ Check existing project remote

  `git remote -v`

+ Create a module by associating it with remote repository

  **(NOTE:** Recommended to create one module per repository, which is located at the project root (filename: `go.mod`)

  `go mod init github.com/username/projectname`
