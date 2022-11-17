# gateway-go
A lightweight Gateway implemented with Go.

It currently supports load balancing algorithms:\
`-round-robin`\
`-random`\
`-power of 2 random choice`\
`-consistent hash`\
`-consistent hash with bounded`\
`-ip-hash`\
`-least-load`

This project based on Gin.

## Install
`> git clone https://github.com/zehuamama/balancer.git`

## Characteristic
### Wildcard (/* or /**)

eg :

config.yaml
```location:                     
  - pattern: /gate/v1/api/product/*
    method: GET
    remove_prefix: true                # remove "/gate"
    proxy_pass:                  
      - "http://localhost:8081"
      - "http://localhost:8082"
    balance_mode: round-robin     

  - pattern: /gate/v1/api/consumer/**
    method: GET
    remove_prefix: true
    proxy_pass:                
      - "http://localhost:8081"
      - "http://localhost:8082"
    balance_mode: round-robin
```

`/gate/v1/api/product/*` can match `/gate/v1/api/product/a`, `/gate/v1/api/product/b` ...\
`/gate/v1/api/product/**` can match `/gate/v1/api/product/a/b`, `/gate/v1/api/product/c/d` ...

## Run
you need to configure the config.yaml file, and run main.go.

## Acknowledgement
zehuamama/balancer.git
gin-gonic/gin