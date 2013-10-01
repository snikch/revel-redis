
# Revel Redis
A simple Redis module for the [Revel Framework](http://robfig.github.io/revel/). Adds a shared [gosexy/redis](https://github.com/gosexy/redis) client to any controller, with a shared connection across all requests.

## Usage

To get started, add the revel-redis module and config to `conf/app.conf`

```
module.redis=github.com/snikch/revel-redis

redis.host=127.0.0.1
redis.port=6379 // Optional
redis.password=abc123 // Optional
```

Now in any controller you want some Redis action, import the library, and add the `RedisController` to your controller struct.

```go

import "github.com/snikch/revel-redis/app"

type MyController struct {
	*revel.Controller
	revelRedis.RedisController
}
```

Now you can feel free to access Redis on your controller

```go
func (c *MyController) DoStuff() revel.Result{
	return c.RenderText(c.Redis.Keys())
}
```

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