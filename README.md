# Gomall
This is a simplified-ver tiktok mall demo based CloudWeGo.

We made some changes to the original demo, the reference address as follow --
[gomall](https://github.com/cloudwego/biz-demo/tree/main/gomall)

Thanks very much for Bytedance's technical lessons!

## Technology Stack
| technology | introduce |
|---------------|----|
| cwgo          | -  |
| kitex         | -  |
| [bootstrap](https://getbootstrap.com/docs/5.3/getting-started/introduction/) | Bootstrap is a powerful, feature-packed frontend toolkit. Build anything—from prototype to production—in minutes.  |
| Hertz         | -  |
| MySQL         | -  |
| Redis         | -  |
| ES            | -  |
| Prometheus    | -  |
| Jaeger        | -  |
| Docker        | -  |


## Biz Logic
- [x] The pages check auth
- [x] Register
- [x] Login
- [x] Logout
- [x] Product categories
- [x] Products
- [x] Add to cart
- [x] The number badge of cart products
- [x] Checkout
- [x] Payment
- [x] Orders center

## How to use
### Prepare 
List required
- Go
- IDE / Code Editor
- Docker
- [cwgo](https://github.com/cloudwego/cwgo)
- kitex `go install github.com/cloudwego/kitex/tool/cmd/kitex@latest`
- [Air](https://github.com/cosmtrek/air)
- ...

### Clone code
```
git clone ...
```

### Copy `.env` file
```
make init
```
*Note:*`You must generate and input SESSION_SECRET random value for session`

### Download go module
```
make tidy
```

### Start Docker Compose
```
make env-start
```
if you want to stop their docker application,you can run `make env-stop`.

### Run Service
This cmd must appoint a service.

*Note:* `Run the Go server using air. So it must be installed`
```
make run svc=`svcName`
```

If you want to run the service selectively, you can edit the *~/scripts/start.sh* file and execute 
```
make run_all
```

Its default content is shown below:

```
#!/bin/bash

make run svc=cart &
make run svc=checkout &
make run svc=email &
make run svc=frontend &
make run svc=order &
make run svc=payment &
make run svc=product &
make run svc=user &
wait
```

### View Gomall Website
```
make open-gomall
```
### Check Registry
```
make open-consul
```
### Make Usage
```
make
```
## Contributors
- [rogerogers](https://github.com/rogerogers)
- [baiyutang](https://github.com/baiyutang)
