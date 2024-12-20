# REST Beer 🍻

A Golang REST API sample

![CI](https://github.com/artursilveiradev/rest-beer/actions/workflows/ci.yml/badge.svg)

## API requests 

### Store beer
```
curl -X "POST" "http://localhost:8080/v1/beer" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
  "name": "Heineken",
  "type": 2,
  "style": 6
}'
```

### Show beer
```
curl "http://localhost:8080/v1/beer/1" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Show beers
```
curl "http://localhost:8080/v1/beer" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```