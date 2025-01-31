package main

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)


func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Debug().Msg("Hello, World!")
	 log.Log().
        Str("foo", "bar").
        Msg("")

	err := errors.New("seems we have an error here")
	log.Error().Err(err).Msg("oh no")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	err = outer()
	log.Error().Stack().Err(err).Msg("")

	var a int = 1
	var b int = 2
	b, a = a, b
	log.Error().Int("a", a).Int("b", b).Msg("")
	var c int = 1
	c = 2
	log.Error().Int("c", c).Msg("")
	m1, _ := GetMoney()
	log.Error().Int("m1", m1).Msg("")
	http.ListenAndServe(":8080", nil)

}

func GetMoney() (int,int) {
	return 1,2
}
func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}