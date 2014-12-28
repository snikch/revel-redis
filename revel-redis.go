// This module configures a redis connect for the application
package revelRedis

import (
	"github.com/gosexy/redis"
	"github.com/revel/revel"
	"os"
	"regexp"
	"strings"
	"strconv"
	"fmt"
)

var (
	Redis *redis.Client
)

func Init() {
	// Read configuration.
	var found bool
	var host string
	var password string
	var port int

	// First look in the environment for REDIS_URL
	url := os.Getenv("REDIS_URL")

	// Check it matches a redis url format
	if match, _ := regexp.MatchString("^redis://(.*:.*@)?[^@]*:[0-9]+$", url); match {

		// Remove the scheme
		url = strings.Replace(url, "redis://", "", 1)
		parts := strings.Split(url, "@")

		// Smash off the credentials
		if len(parts) > 1 {
			url = parts[1]
			password = strings.Split(parts[0], ":")[1]
		}

		// Split to get the port off the end
		parts = strings.Split(url, ":")
		if len(parts) != 2 {
			revel.ERROR.Fatal(fmt.Sprintf("REDIS_URL format was incorrect (%s)", url))
		}

		// Get the host and possible password
		var port64 int64
		host = parts[0]
		port64, _ = strconv.ParseInt(parts[1], 0, 0)
		if port64 > 0{
			port = int(port64)
		}
	}

	// Then look into the configuration for redis.host and redis.port
	if len(host) == 0 {
		if host, found = revel.Config.String("redis.host"); !found {
			revel.ERROR.Fatal("No redis.host found.")
		}
	}
	if len(password) == 0 {
		password, _ = revel.Config.String("redis.password")
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

	// Attempt to authenticate
	if len(password) != 0 {
		m, err := Redis.Auth(password)
		if err != nil {
			revel.ERROR.Fatal(fmt.Sprintf("Could not authenticate redis: %s", m))
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
}
