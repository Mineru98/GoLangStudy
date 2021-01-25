# GoLangStudy

## 초기 설정
```sh
touch go.mod
touch go.sum
```

## 빌드
```sh
// 크로스 플렛폼 배포 빌드 코드
GOOS=linux GOARCH=amd64 go build main.go
```

## 패키지 설치(예시)
```sh
go get github.com/golang/example/appengine-hello
go get github.com/golang/example/gotypes
go get github.com/golang/example/outyet
go get github.com/golang/example/stringutil
go get github.com/golang/example/template
```

## 참고 자료
https://brownbears.tistory.com/category/%EC%96%B8%EC%96%B4/Golang

https://mingrammer.com/go-modules-private-repo/

https://ingeec.tistory.com/106

## Go CLI Argument
- build : 패키지 및 종속성을 컴파일
- clean : 오브젝트 파일 및 캐시된 파일 제거
- doc : 패키지 또는 기호에 대한 설명서 표시
- env : Go 환경 정보 인쇄
- bug : 버그리포트를 시작
- fix : 새 API를 사용하도록 패키지 업데이트
- fmt : gfmt 패키지 소스
- generate : 처리 소스를 통해 이동 파일 생성
- get : 패키지 및 종속성 다운로드 및 설치
- install : 패키지 및 종속성을 컴파일 및 설치
- list : 패키지를 나열
- mod : 모듈 유지 관리
- run : Go 프로그램 실행
- test : 테스트 패키지
- tool : 지정된 Go 도구를 실행
- version : Go 버전 출력
- vet : 패키지에 있을 법한 실수를 보고

- buildconstraint : 제약조건을 생성
- buildmode : 모드를 구축
- c : Go와 C 사이에서 호출
- cache : 빌드 및 테스트 캐슁
- environment : 환경 변수
- filetype : 파일 형식
- go.mod : 그 go.mod 파일
- gopath : GOPASS 환경 변수
- gopath-get : 기존 GOPASS Go get
- goproxy : 모듈 프록시 프로토콜
- importpath : 가져오기 경로 구문
- modules : 모듈, 모듈 버전 등
- module-get : Go 모듈 인식
- module-auth : Go를 사용한 모듈 인증.합계를 내다
- module-private : 비공용 모듈에 대한 모듈 구성
- packages : 패키지 나열
- testflag : 테스트 플래그
- testfunc : 테스트 함수