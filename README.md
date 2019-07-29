# spidy
A simple web crawler written in go.

## Description

This tool tries to crawl through the links on any given page and list
the urls it finds.

## Dependencies and Building

It uses the [goquery](https://github.com/PuerkitoBio/goquery) to parse
the HTML page and fetch the links from them. There is no other external
dependencies for the application.

You can run `make dep` to install the dependencies.

Note: If you want to run `make lint` you need to have
[golint](https://github.com/golang/lint) installed.

## Building

To build a binary for Mac (Darwin) you can simply run `make build`. But
if you want to build the binary for Linux run the below command on a linux
instance.

```bash
go build -i -o spidy
```

## Running

Simply run the application with url you want o crawl as the first arg.

```bash
./spidy https://xkcd.com
```

It will print all the links which belong to the same site. So in the example
above, all the links of https://xkcd.com domain will be listed.

## Considerations and Limitations

As of now, it doesn't limit the concurrency. But it is a good idea to limit
the cocurrency to the number of cpu cores available the machine.

Also, while testing the application I quickly found that crawling through the
links recursively can take hours or not even end. So it might be a great idea
to limit the layers/depth of crawling.

## Planned Enhancements

1. Sanitise the user input and reject incorrect starting URL.
1. Limit the concurrency of the process.
1. Set the depth until which the the crawling should be done.
1. Modify the result struct so that a sitemap of some kind can be printed.
1. Add more unit tests.
1. Setup a CI/CD system which runs some validations before merging.
