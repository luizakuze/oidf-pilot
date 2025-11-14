# Example RP using go-oidfed library
This is an example RP that uses the 
[go-oidfed library](https://github.com/go-oidfed/lib).

It is very rudimentary and by no-means meant as a production RP. It's just 
for demonstration purposes and may crash on any error.

It also does not do much useful, but you can do a login using oidfed.
[OFFA](https://oidfed.github.com/offa) might be a more useful alternative.

## How to deploy / setup

The whoami-rp is provided as a docker image
[`oidfed/whoami-rp`](https://hub.docker.com/r/oidfed/whoami-rp).

One can run it like:
```bash
docker run -v gorp-keys:/keys -v /path/to/config/config.yaml:/config.yaml oidfed/whoami-rp
```

However, you probably want to run it with other components, so a docker compose
might make more sense.
An example [docker-compose file](docker-compose.yaml) can be found in this 
repository. However, it is very minimal and does not contain other services.

For configuration please refer to the [config-example.yaml](config-example.
yaml) file.

---

Please note that support and development effort for this dummy RP is limited.

[OFFA](https://oidfed.github.com/offa) might be an alternative.
