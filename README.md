# Gecko

A simple Docker file for Golang + Echo web framework.

## Installation + Building

```
docker build -t gecko .
```

## Running

```
docker run -i -t -v ~/path/to/local/app:/app -p 3000:3000 gecko
```
