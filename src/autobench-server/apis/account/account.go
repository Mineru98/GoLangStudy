package account

import (
	"autobench-server/apis/account/signin"
	"autobench-server/apis/account/signup"
	"autobench-server/apis/api"
)

// AccountAPIs 는 Account에 대한 API 리스트입니다
var AccountAPIs []api.API

func init() {
	AccountAPIs = make([]api.API, 0)

	apis := []api.API{
		new(signin.API),
		new(signup.API),
	}

	for i := 0; i < len(apis); i++ {
		AccountAPIs = append(AccountAPIs, apis[i])
	}
}