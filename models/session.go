package models

import "github.com/gorilla/sessions"

const SESSION_ID = "go_auth_sesssion"

var Store = sessions.NewCookieStore([]byte("asdsadsaddsadsa"))
