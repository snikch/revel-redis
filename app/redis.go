// This module configures a redis connect for the application
package revelRedis

import (
	"github.com/gosexy/redis"
	"github.com/robfig/revel"
)

var (
	Redis *redis.Client
)

func Init() {
	// Read configuration.
	var found bool
	var host string
	var port int
	if host, found = revel.Config.String("redis.host"); !found {
		revel.ERROR.Fatal("No redis.host found.")
	}
	if port, found = revel.Config.Int("redis.port"); !found {
		port = 6379
	}

	Redis = redis.New()

	// Open a connection.
	var err error
	err = Redis.Connect(host, uint(port))
	if err != nil {
		revel.ERROR.Fatal(err)
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
}
