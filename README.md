# Balancing proxy
Simple proxy server with round robin balancing.

Before run server need to create config file with routes. Config file sample:

```
{
  "routes": [
    {
      "local_addr": "0.0.0.0:4242",
      "remote_addrs": [
        "127.0.0.1:8080",
        "127.0.0.1:8081",
        "127.0.0.1:8082"
      ]
    },
    {
      "local_addr": "0.0.0.0:4243",
      "remote_addrs": [
        "127.0.0.1:8090",
        "127.0.0.1:8091",
        "127.0.0.1:8092"
      ]
    }
  ]
}
```
