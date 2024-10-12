## How to run go file?

```
 go run [name of file]
 ex) go run hello.go
```

## To turn a Go file into an executable, a build process is required.

### To build Go file

step 1 : Need to make mod

```
go mod init [name of module/file]
ex) go mod init goprojects/hello
```

step 2 : Build Go file

```
go build
```

step 3: Execute the file

```
./hello
```

go mod init goprojects/hello 명령어는 Go 프로젝트에서 모듈 초기화를 위해 사용됩니다. 이 명령어를 실행하면 해당 디렉토리에서 Go 모듈을 설정하고, 모듈과 관련된 정보를 담고 있는 go.mod 파일을 생성합니다. 이 명령어의 역할을 좀 더 구체적으로 설명하자면:

1. 모듈 초기화 (go mod init)

   • 모듈은 Go에서 의존성 관리를 위해 사용하는 개념입니다. 모듈은 하나의 프로젝트 또는 라이브러리를 의미하며, 여러 패키지로 구성될 수 있습니다.
   • go mod init 명령어는 현재 디렉토리를 새로운 Go 모듈로 설정합니다.
   • 이 명령어를 실행하면 **go.mod**라는 파일이 생성되며, 이 파일은 모듈의 이름과 버전 정보, 그리고 의존성 목록을 관리합니다.

2. goprojects/hello

   • goprojects/hello는 이 모듈의 이름을 지정하는 부분입니다. 보통 Go 모듈 이름은 실제 저장소의 경로나 프로젝트 이름으로 설정됩니다. 예를 들어, GitHub에 프로젝트를 올린다면 github.com/username/projectname처럼 설정하는 것이 일반적입니다.
   • 그러나 로컬에서 작업할 때는 goprojects/hello와 같은 임의의 이름을 사용해도 됩니다.
   • 이 이름은 나중에 이 모듈을 다른 곳에서 불러올 때 사용됩니다. 즉, 다른 프로젝트가 이 모듈을 의존성으로 추가할 때 이 모듈 이름을 이용해 가져오게 됩니다.

3. 결과

   • 이 명령어를 실행한 후, 프로젝트 루트에 go.mod 파일이 생기고, 그 파일 안에는 모듈 이름 (goprojects/hello)과 Go 버전 정보가 들어가게 됩니다. 예시:

```
module goprojects/hello

go 1.20
```

이제 이 모듈은 Go 프로젝트로서 관리되며, 패키지를 추가하거나 의존성을 관리할 수 있게 됩니다.

쉽게 말해, 이 명령어는 현재 프로젝트를 Go 모듈로 설정하고, Go에서 의존성을 관리할 수 있도록 준비해주는 과정입니다.

### Tip build까지 다하고 수정이 이루어 지면 어떻게 변경?

=> 다시 go build만 진행 하면 된다.
