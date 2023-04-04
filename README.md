# blog-maintenance tool

[ta2mo's blog](https://ta2mo.github.io/)
Generate static html tool .

## Development

### build
```shell script
$ go build 
```

### run blog in local
```
$ ./blog-maintenance convert
$ cd nuxt-template
$ npm install
$ npm run dev

# open browser http://localhost:13000/
```

## blog maintenance command

### create new post

```shell script
$ ./blog-maintenance new # or n
Please input post information
post title       : test
post categories  : programming, golang
generete success : ./post/2006-01-02-new.md

# edit post
$ vim ./post/2006-01-02-new.md
```

### convert from markdown files to html

```shell script
$ ./blog-maintenance convert # or c
```

### generate and update search index

```shell script
$ ALGOLIA_API_KEY=1234abcd ./blog-maintenance index # or i
```
