package main

import (
	pkgConfig "device-manager/pkg/flags"
	"device-manager/pkg/restapi"
	"device-manager/pkg/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

func main(){
	conf := &pkgConfig.Config{}

	parser := flags.NewParser(conf, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	zerolog.MessageFieldName = "MESSAGE"
	zerolog.LevelFieldName = "LEVEL"

	var output io.Writer
	if conf.UseSyslog == 1 {
		var err error
		output, err = os.OpenFile(conf.LogDir, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	} else {
		output = os.Stderr
	}

	log.Logger = log.Output(output).With().Str("PROGRAM", "device-manager").Logger()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatal().Msg("loading swagger spec failed")
	}

	api := operations.NewDeviceManagerAPI(swaggerSpec)
	api.Logger = log.Printf
	server := restapi.NewServer(api)

	defer server.Shutdown()

	//redis
	//redisOpt := redis.Options{Addr: conf.RedisHost}
	//redisClient := redis.NewClient(&redisOpt)
	//if _, err = redisClient.Ping().Result(); err != nil {
	//	panic(err)
	//}

	handler := configureAPI(api, conf)
	server.SetHandler(handler)
	server.Port = conf.AppPort

	if err := server.Serve(); err != nil {
		log.Fatal().Msg("server run failed")
	}
}