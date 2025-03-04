package runtimedef

import (
	"os/user"

	"m2y/utils/userutil"
)

var _executer user.User

func initExecuter() error {
	u, err := userutil.GetRealUser()
	if err != nil {
		return err
	}
	_executer = *u
	return nil
}

func GetExecuter() user.User {
	return _executer
}
