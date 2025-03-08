# REST Beer 🍻

A Golang REST API sample

![CI](https://github.com/artursilveiradev/rest-beer/actions/workflows/ci.yml/badge.svg)

## API requests 

### Create beer
```
curl -X "POST" "http://localhost:8080/v1/beers" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
  "name": "Heineken",
  "type": 2,
  "style": 6
}'
```

### Read beer
```
curl -X "GET" "http://localhost:8080/v1/beers/1" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Read beers
```
curl -X "GET" "http://localhost:8080/v1/beers" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Update beer
```
curl -X "PATCH" "http://localhost:8080/v1/beers/1" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
  "name": "Budweiser"
}'
```

### Delete beer
```
curl -X "DELETE" "http://localhost:8080/v1/beers/1"
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```
