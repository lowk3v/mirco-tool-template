package utils

import "github.com/author_name/project_name/config"

func HandleError(err error, customMsg string) bool {
	if err != nil {
		config.Log.
			WithField("msg", customMsg).
			Error(err)
		return true
	}
	return false
}
