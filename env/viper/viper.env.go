package env_viper

import (
	"bytes"
	"errors"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func parseName(path string) (formatName, configName string, err error) {
	splitPaths := strings.Split(path, "/")
	if len(splitPaths) > 0 {
		for i := 0; i < len(splitPaths); i++ {
			configName = splitPaths[i]
		}
	}
	splitNames := strings.Split(configName, ".")
	if len(splitNames) < 2 {
		err = errors.New("failed to parse config name")
		return
	}
	formatName = splitNames[len(splitNames)-1]
	return
}

// InitViper initialize Viper to use the config file as env variable
func InitViper(path string, env interface{}, log *zerolog.Logger) error {
	formatName, configName, err := parseName(path)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return err
	}
	viper.SetConfigName(strings.TrimRight(configName, "."+formatName))
	viper.SetConfigType(formatName)
	viper.AddConfigPath(strings.TrimRight(path, configName))
	if err = viper.ReadInConfig(); err != nil {
		log.Error().Msgf("Configs file: %v", err)
		return err
	}
	if err = viper.Unmarshal(&env); err != nil {
		log.Error().Msgf("Configs unmarshar error: %v", err)
	}
	return err
}

// InitRemoteViper initialize Viper using remote config
func InitRemoteViper(user, pass, url string, env interface{}, log *zerolog.Logger) (err error) {
	client := resty.New()
	res, err := client.R().
		SetBasicAuth(user, pass).
		Get(url)
	if err != nil {
		log.Err(err).Send()
		return
	}
	if res.IsError() {
		err = errors.New(res.String())
		if res.StatusCode() == http.StatusUnauthorized {
			err = errors.New("wrong config's credential")
		}
		log.Error().Msg(res.String())
		return errors.New(res.String())
	}
	formatName, _, err := parseName(url)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	viper.SetConfigType(strings.ToLower(formatName))
	if err = viper.ReadConfig(bytes.NewReader(res.Body())); err != nil {
		log.Warn().Msgf("Remote configs: %v", err)
		return
	}

	if err = viper.Unmarshal(&env); err != nil {
		log.Warn().Msgf("Configs unmarshar error: %v", err)
	}
	return
}
