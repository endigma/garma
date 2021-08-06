package common

import (
	"github.com/rs/zerolog/log"
)

func CheckErr(err error, label string) {
	if err != nil {
		log.Fatal().Err(err).Msg(label)
	}
}
