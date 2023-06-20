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
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gin-gonic/gin
```

#### Run with CompileDaemon

```bash
compiledaemon --command="./ref"
```