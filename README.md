# probes

### Build
```
docker build -t probes:v1 .
docker tag probes:v1 us.icr.io/order-mgmt-np/oms/probes:v1
docker push us.icr.io/order-mgmt-np/oms/probes:v1
```

### Usage
```
./probes <wait-time-in-minutes>

eg., ./probes 10
```
