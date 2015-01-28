# Go-Domains

A simple RESTful API in Go for storing domain names into a MongoDB collection.

## Running the Server

Make sure you have MongoDB installed and running on a standard port.

```
src/domains/ $ go install
src/domains/ $ domains
[martini] listening on :3000 (development)
```

## Running the Tests (WIP)

You'll need MongoDB running for these as well.

```
src/domains/ $ ginkgo -r --randomizeAllSpecs -cover
```