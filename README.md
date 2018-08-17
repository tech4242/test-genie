[![Build Status](https://travis-ci.org/tech4242/test-genie.svg?branch=master)](https://travis-ci.org/tech4242/test-genie)

# test-genie
A small Go server to help you unit test your frontend by fetching "real" mock data from your live server. The whole point is to let you unit test your frontend with up-to-date data instead of static json files in your frontend code.

![Test Genie](/img/test-genie.png)

## Setup

### Config

Add the `url` that you want to route to in `config.yml`. `live: true` will run the "live" reverse proxy server that will route your `url` through `localhost:9000`. Setting `live: false` will allow you to run a local mock server with the saved data from the live server _but this feature is still under development_.

Duplicate `config.yml.dist` as `config.yml` (in .gitignore):

```
cp config.yml.dist config.yml
```

An example .yml file:

```
host:
  url: "https://www.google.com"
  live: true

```

### Run with make

```
make deps
make
make run
```

This will run on `localhost:9000` and map to the `url` in the `config.yml`

### Run with Docker

```
docker build . --tag=test-genie
docker run -p 9000:9000 test-genie:latest
```

### Postman

You can just create your normal Postman requests now and you can pass your token (Bearer or otherwise) to the live server as we only pass the request to your actual server.

## Wishlist

This list will hopefully get larger and not smaller over time :) The current to-do's are not here but instead under "Projects" in a Kanban board! 

* Build a generic parser for Swagger and Postman, so you can do `make run` and depending on your .yml config run all your existing APIs instead of calling them one by one from the frontend.
* Spend more than 30 min on the logo.

## Credits

Special thanks to @turbinenreiter for the help with the Makefile and to @signalkraft for the initial idea for how to solve this! Kudos to Google for golang and thanks to Takuya Ueda for the gopher vector under CC!
