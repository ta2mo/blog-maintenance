# blog-maintenance tool

## Development

```
$ dep ensure
$ ./go run main.go generate
$ cd nuxt-template
$ yarn install
$ yarn run dev

# open browser http://localhost:13000/
```

## command

### create post scaffold

```
$ go run main.go new $TITLE $CATEGORIES
```
