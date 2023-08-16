package api

import "gvd_server/api/user_api"

type Api struct {
	UserApi user_api.UserApi
}

var App = new(Api)
