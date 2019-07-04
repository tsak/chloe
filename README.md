# Chloe

An image generator for the meat conscious.

This started on a friends Slack with someone posting [this image](https://i.redd.it/mfztzrcfct131.jpg)
*(Trigger warning: an image of wrapped meat on a super market shelf)* and then the idea of an image generator
was born.

You can try it on [cow.name](https://cow.name)

# Build

```bash
go build
```

# Assets

To update the assets in `assets.go`, install [go-bindata](https://github.com/go-bindata/go-bindata)

```bash
go get github.com/go-bindata/go-bindata
go-bindata -o assets.go assets
```