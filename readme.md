##### Initiate a new project
```bash
go mod init github.com/username/projectname
```

##### Install a package
```bash
go get -u gorm.io/gorm 
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/joho/godotenv
go get -u github.com/golang-jwt/jwt/v4
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon@latest
go get -u gorm.io/driver/postgres
go get -u github.com/gin-gonic/gin
```

#### Run with CompileDaemon

```bash
compiledaemon --command="./ref"
```

To use air for auto reload

Ensure you have your project folder structure in the C:\Users\{username}\go\src folder
e.g: /c/Users/user/go/src/gotodo


To use air for auto reload

Ensure you have your project folder structure in the C:\Users\{username}\go\src folder
e.g: /c/Users/user/go/src/gotodo


Install air

```bash
go get -u github.com/cosmtrek/air  
```



In your root folder add a .air.toml file

```toml
# .air.toml

root = "."
tmp_dir = "tmp"
exclude_dir = ["tmp", "vendor"]
include_ext = [".go"]
build_args = ["-tags", "development"]
debug = true
notify_delay = 5

[build]
  bin = "./tmp/main.exe"
  cmd = "go build -o ./tmp/main.exe ."
  delay = 1000
  exclude_dir = ["tmp", "vendor"]

[[event]]
  delay = 1000
  path = "**/*.go"
  exclude_paths = []
  command = "go run ."
```

Then run air in your terminal

```bash
air
```

If you have any issues
Run 

```bash
export PATH="$PATH:$GOPATH/bin"
```

Save the source
    
```bash
source ~/.bashrc
```

Verify the air executable is in your path

```bash
ls $GOPATH/bin/air
```

Ensure your GOBIN is set, check by running the following command:-

```bash
go env GOBIN
```

If the output is empty, set the GOBIN by running the following command:-

```bash
go env -w GOBIN=$GOPATH/bin or export GOBIN=$HOME/.local/bin
```

Run:
  
```bash
go env GOBIN
```

If the file is found, it means air is properly installed and accessible. If not, try reinstalling air by running the following command:-
    
```bash
go install github.com/cosmtrek/air@latest
```

```bash
air
```


---

#### set the environment variables in your .bash_profile


```bash 
nvim ~/.bashrc or nvim ~/.bash_profile
```

#### add the following lines to your .bash_profile

```bash
export GOPATH=C:\Users\user\go
```

#### apply the changes

```bash
source ~/.bashrc
```