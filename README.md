# simple-swag

Dead simple swagger/openapi docs server, under the covers we use 
https://github.com/Redocly/redoc to present the docs.

## Compatibilty

At the present time this server will work for OpenApi/swagger docs version 2 
or later.

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
