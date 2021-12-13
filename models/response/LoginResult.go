package response

import (
	"firstGin/models/request"
)

type LoginResult struct {
	User  request.User
	Token string
}
