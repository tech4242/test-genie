# test-genie
A small Go server to help you unit test your frontend by fetching mock data from your live server.

## Run locally

```
make
REDIRECTURL=http://www.google.com/ make run
```

This will let you do a GET to `localhost:9000`, which will reroute you to google.com


## Credits

Special thanks to @turbinenreiter for the help with the Makefile and to @signalkraft for the initial idea for how to solve this! 
