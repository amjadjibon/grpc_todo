package utils

import (
	"github.com/alexandrevicenzi/unchained"
	"github.com/rs/zerolog/log"
)

func CreateHash(password string) string {
	hash, err := unchained.MakePassword(password, unchained.GetRandomString(12), "default")
	if err != nil {
		log.Err(err).Msg("error encoding password")
	}
	return hash
}

func CheckPassword(password, hash string) bool {
	valid, err := unchained.CheckPassword(password, hash)

	if valid {
		log.Info().Msg("password is valid")
	} else {
		if err == nil {
			log.Warn().Msg("password is invalid")
		} else {
			log.Err(err).Msg("error decoding password")
		}
	}
	return valid
}