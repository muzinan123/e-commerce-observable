package common

import (
	"strconv"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {

	consulSource := consul.NewSource(

		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),

		consul.WithPrefix(prefix),

		consul.StripPrefix(true),
	)

	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}

	err = conf.Load(consulSource)
	return conf, err
}
