# check cli
curl -v -X POST "localhost:8001/work?delay=1s&name=foo"

# expect output
```
Starting the dispatcher
Starting worker 1
Starting worker 2
Starting worker 3
Starting worker 4
Registering the collector
HTTP server listening on 127.0.0.1:8001
2019/08/06 11:31:17 Work request queued
Received work requeust
Dispatching work request
2019/08/06 11:31:17 Doing shiit
```