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
      "weight": 0,
      "uri": "http://localhost:8001"
    },
    {
      "host": "localhost",
      "port": 8002,
      "weight": 0,
      "uri": "http://localhost:8002"
    },
    {
      "host": "localhost",
      "port": 8003,
      "weight": 0,
      "uri": "http://localhost:8003"
    }
  ],
  // 0: round robin, 1: weighted round robin (WIP)
  "protocol": 0
}

```

### Testing
To test the load balancer for race conditions:
```bash
go test -race
```

## Author
Chaitanya Mittal, 2023
