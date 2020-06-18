# Chloe

An image generator for the meat conscious.

This started on a friends Slack with someone posting [this image](https://i.redd.it/mfztzrcfct131.jpg)
*(Trigger warning: an image of wrapped meat on a super market shelf)* and then the idea of an image generator
was born.

You can try it on [chloe.tsak.net](https://chloe.tsak.net)

## Build

```bash
go build
```

## Run

```bash
Usage of ./chloe:
  -debug
    	Debug mode
  -host string
    	Hostname (default "localhost")
  -port string
    	Port (default "8000")
```

```bash
./chloe # defaults to localhost:8000
```

## Assets

To update the assets in `assets.go`, install [go-bindata](https://github.com/go-bindata/go-bindata)

```bash
go get github.com/go-bindata/go-bindata
go-bindata -o assets.go assets
```

## Run as a systemd service

See [chloe.service](chloe.service) systemd service definition.

To install (tested on Ubuntu 16.04):

1. `adduser chloe`
2. copy `chloe` binary to `/home/chloe`
3. place systemd service script in `/lib/systemd/system/`
4. `sudo systemctl enable chloe.service`
5. `sudo systemctl start chloe`
6. `sudo journalctl -f -u chloe`

The last command will show if the service was started.
