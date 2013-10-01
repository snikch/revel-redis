// This module configures a redis connect for the application
package revelRedis

import (
	"github.com/gosexy/redis"
	"github.com/robfig/revel"
	"os"
	"regexp"
	"strings"
	"strconv"
)

var (
	Redis *redis.Client
)

func Init() {
	// Read configuration.
	var found bool
	var host string
	var port int

	// First look in the environment for REDIS_URL
	url := os.Getenv("REDIS_URL")

	// Check it matches a redis url format
	if match, _ := regexp.MatchString("^redis://.*:[0-9]+$", url); match {

		// Remove the scheme
		url = strings.Replace(url, "redis://", "", 1)

		// Split to get the port off the end
		parts := strings.Split(url, ":")
		port64, _ := strconv.ParseInt(parts[len(parts)-1], 0, 0)
		if port64 > 0{
			port = int(port64)
		}

		// Remove the port part and join to get the hostname
		parts = parts[:len(parts)-1]
		host = strings.Join(parts, ":")
	}

	// Then look into the configuration for redis.host and redis.port
	if len(host) == 0 {
		if host, found = revel.Config.String("redis.host"); !found {
			revel.ERROR.Fatal("No redis.host found.")
		}
	}
	if port == 0 {
		if port, found = revel.Config.Int("redis.port"); !found {
			port = 6379
		}
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
