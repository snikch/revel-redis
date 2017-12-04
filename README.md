
# Revel Redis
A simple Redis module for the [Revel Framework](http://revel.github.io/). Adds a shared [gosexy/redis](https://github.com/gosexy/redis) client to any controller, with a shared connection across all requests.

## Usage

To get started, add the revel-redis module and config to `conf/app.conf`

**NOTICE**:
You'd better add this in the body of the `conf/app.conf`, 
not append it at the bottom, which leads to empty value. For example you
may just copy this and paste it under the `http.sslkey` section.

```
############# Redis Connection Info ##########
module.redis=github.com/snikch/revel-redis
redis.host = 127.0.0.1
#Optional
redis.port = 6379
#Optional
redis.password =
##############################################
```

Now in any controller you want some Redis action, import the library, and add the `RedisController` to your controller struct.

```go

import "github.com/snikch/revel-redis"

type MyController struct {
	*revel.Controller
	revelRedis.RedisController
}
```

Now you can feel free to access Redis on your controller

```go
func (c *MyController) DoStuff() revel.Result{
	return c.RenderText(c.Redis.Keys("*"))
}
```

### ENV
If you want to use an environment variable instead of a hardcoded conf value, Revel Redis will prioritise `REDIS_URL` in the format `redis://hostname:port`, over the conf values.

## Who

Created with love by [Mal Curtis](http://github.com/snikch)

Twitter: [snikchnz](http://twitter.com/snikchnz)

## License

MIT. See license file.

## Todo

*  Handle global processes when starting another instance of hack


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
