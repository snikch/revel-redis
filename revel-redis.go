// This module configures a redis connect for the application
package revelRedis

import (
	"github.com/gosexy/redis"
	"github.com/revel/revel"
	"github.com/revel/revel/logger"
)

var (
	Redis        *redis.Client
	moduleLogger logger.MultiLogger
)

func Init() {
	// Then look into the configuration for redis.host and redis.port
	host, found := revel.Config.String("redis.host")
	if !found {
		moduleLogger.Fatal("No redis.host found.")
	}

	port := revel.Config.IntDefault("redis.port", 6379)

	password, _ := revel.Config.String("redis.password")

	Redis = redis.New()

	// Open a connection.
	var err error
	err = Redis.Connect(host, uint(port))
	if err != nil {
		moduleLogger.Fatal(err.Error())
	}

	// Attempt to authenticate
	if len(password) != 0 {
		m, err := Redis.Auth(password)
		if err != nil {
			moduleLogger.Fatalf("Could not authenticate redis: %s", m)
		}
	}
}

type RedisController struct {
	*revel.Controller
	Redis *redis.Client
}

func (c *RedisController) Begin() revel.Result {
	c.Redis = Redis
	return nil
}

func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*RedisController).Begin, revel.BEFORE)
	revel.RegisterModuleInit(func(module *revel.Module) {
		moduleLogger = module.Log
	})
}
