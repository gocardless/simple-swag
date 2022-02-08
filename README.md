# simple-swag

Dead simple Swagger / OpenApi docs server, under the covers we use
https://github.com/Redocly/redoc to present the docs.

## Compatibilty

At the present time this server will work for OpenApi/Swagger docs version 2 
or later.

## Install

The easiest way to install is using go install

```
go install github.com/gocardless/simple-swag@latest
```

## Usage

```
simple-swag
  -filename string
    	a path to the swagger/openapi spec (default "~/swagger.json")
  -host string
    	host ip to serve using (default "127.0.0.1")
  -port int
    	port to serve http over (default 9000)
```

## Troubleshooting

If you get the following error when attempting to run `simple-swag` it is 
likely that Go isn't initialised properly in your shell.

Add the below lines to your `~/.zshrc` file.

```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```
Run `source ~/.zshrc`

Try running simple-swag again

## License & Contributing

* SimpleSwag is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
* Bug reports and pull requests are welcome on GitHub at https://github.com/gocardless/simple-swag

GoCardless ♥ open source. If you do too, come [join us](https://gocardless.com/about/careers/).
