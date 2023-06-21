# Load Balancer in Go

## About
A simple implementation in Go to better understand the working of a load balancer.

## Usage
Open up your terminal and run the following:
```bash
go run .
```

In another terminal window:
```bash
for i in {1..6}
do
curl http://localhost:8000
done
```

### Configuration
The `config.json` file looks something like this:
```json
{
  "lb": {
    "host": "localhost",
    "port": 8000
  },
  "servers": [
    {
      "host": "localhost",
      "port": 8001,
      "weight": 3,
      "uri": "http://localhost:8001"
    },
    {
      "host": "localhost",
      "port": 8002,
      "weight": 2,
      "uri": "http://localhost:8002"
    },
    {
      "host": "localhost",
      "port": 8003,
      "weight": 1,
      "uri": "http://localhost:8003"
    }
  ],
  "protocol": 0
}
```

#### Protocols
0: Weighted Round Robin (Just specify weight = 1 to all servers for Round Robin)

### Testing
To test the load balancer for race conditions:
```bash
go test -race
```

## TODO
1. Support for all methods (currently supports only GET & POST)
2. Dynamic Load Balancing algorithms

## Author
Chaitanya Mittal, 2023
