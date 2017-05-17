# Build drone-cmds-image
```
docker build -t my-drone-cmds-image .
```

# Run drone-cmds in interactive mode
```
docker run -it --rm --name my-drone-cmds-app -p 3000:3000 my-drone-cmds-image

curl -i -X POST http://localhost:3000/api/cmds/telemetry \
     -H "Content-Type: application/json" \
     -d "{\"drone_id\":\"drone666\",\"battery\":72,\"uptime\":6941,\"core_temp\":21}"
```

# Drone-cmds + rabbitmq
### Connect with docker rabbitmq instance
- Change rabbitmq url in server.go to 
  ```
  url := "amqp://guest:guest@rabbit:5672/"
  ```

### Build modified docker image
```
docker build -t my-drone-cmds-image .
```

### Run containers
```
# Run rabbitmq container
docker run -d --name my-rabbit -p 8080:15672 -p 4369:4369 -p 5672:5672 rabbitmq:3-management

# Run drone-cmds and link to rabbitmq
docker run -d --name my-drone-cmds-app --link my-rabbit:rabbit -p 3000:3000 my-drone-cmds-image

# Test with a post request
curl -i -X POST http://localhost:3000/api/cmds/telemetry \
     -H "Content-Type: application/json" \
     -d "{\"drone_id\":\"drone666\",\"battery\":72,\"uptime\":6941,\"core_temp\":21}"

# Check rabbitmq admin to check if message was inserted
```
