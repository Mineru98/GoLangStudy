# GoLangStudy
Go Language Study
```
cd /workspace/GoLangStudy/src/github.com/Mineru98/GoLangStudy

// 빌드
export GOPATH="/workspace/GoLangStudy"
go build -o /workspace/GoLangStudy/src/autobench-server/bin//main $(find /workspace/GoLangStudy/src/autobench-server/ -name *.go) && /workspace/GoLangStudy/src/autobench-server/bin//main

// Singup
https://autobench-server.run.goorm.io/api/account/signup
{
	"_id":1,
	"uid":"daslkjdfsaslkjfklasdjf",
	"pw": "123456789",
	"uname":"Test",
	"profile_img":"",
	"point":2
}

// Singin
https://autobench-server.run.goorm.io/api/account/signin
{
	"uname":"Test",
	"pw":"123456789",
	"uid":"daslkjdfsaslkjfklasdjf"
}
```