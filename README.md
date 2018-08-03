# test-genie
A small Go server to help you unit test your frontend by fetching "real" mock data from your live server.

## Run locally

Duplicate `config.yaml.dist` as `config.yaml` (in .gitignore) and add the `url` that you want to route to.

```
make deps
make
make run
```

This will run on `localhost:9000/` and map to the `url` in the `config.yaml`

## Credits

Special thanks to @turbinenreiter for the help with the Makefile and to @signalkraft for the initial idea for how to solve this!
