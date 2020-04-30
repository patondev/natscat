# natscat
Building program like kafkacat but for NATS and STAN

# Sample Commands
All of these sample commands below, assumed the NATS/STAN service is running on localhost with standard ports

## List all subjects online
Before you run this command, make sure the port 8222 is opened and monitoring has been enabled.
```
./natscat --addr http://127.0.0.1:8222 subjects
```

# Notes
For testing purpose, better to run in localhost to avoid unexpected problems (like network congestion, etc). You may change the address for testing on `internal/nats/global.go`