# Setup

+ Install Go

  `brew install go`

+ Install project dependencies

  ```
  go get -u golang.org/x/tools/...
  go get -u github.com/gorilla/mux
  go get github.com/go-playground/validator/v10
  ```

+ Install `jq` to format `curl` output

  `brew install jq`

+ Check existing project remote

  `git remote -v`

+ Create a module by associating it with remote repository

  **(NOTE:** Recommended to create one module per repository, which is located at the project root (filename: `go.mod`)

  `go mod init github.com/username/projectname`
