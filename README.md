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
for i in {1..5}
do
curl http://localhost:8000
done
```

### Testing
To test the load balancer for race conditions:
```bash
go test -race
```

## Author
Chaitanya Mittal, 2023
