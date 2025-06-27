# client-server

Project that tests different configurations of http client. It is influenced by
[connection churn](https://dev.to/gkampitakis/http-connection-churn-in-go-34pl) blogpost.

## prerequisites
- build `go build`
- netstat command has to be present
 
## usage

- default, body not closed and not read `./client-server`
- body closed but not read `./client-server -close-body true`
- body read but not closed `./client-server -read-body true`
- ...

### available flags

- `-close-body bool` close response body
- `-concurrency int` concurrent connections (default 2)
- `-max-idle-conn int` max idle connections (default 100)
- `-max-idle-conn-host int` max idle connections per host (default 2)
- `-port int` server port (default 8080)
- `-read-body` read response body

## example

```
./client-server

flags set to: {Port:8080 Concurrency:2 MaxIdleConnsPerHost:2 MaxIdleConns:100 ReadBody:false CloseBody:false}
server listening on localhost:8080
requests        connections     TIME_WAIT       CLOSE_WAIT      ESTABLISHED
    1164              15159         12803                0             2480
    2249              17304         12803                0             4498
    3012              18831         12803                0             6025
    3864              20551         12803                0             7892
    4798              21789         13421                0             8508
    5575              22384         14454                0             7704
    6285              22663         15515                0             6864
     ...                ...           ...              ...              ...
```
