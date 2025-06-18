package common

import (
	"time"

	"github.com/spf13/viper"
)

func Timeout() time.Duration {
	timeoutConfig := viper.GetInt("HTTP_CLIENT_TIMEOUT")
	return time.Duration(timeoutConfig * int(time.Second))
}
