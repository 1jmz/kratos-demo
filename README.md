## Start consul
```
consul agent -dev
```

## Start calc_engine/calc_gateway micro_service
calc_gateway提供http服务，calc_engine提供gRPC服务，calc_gateway调用calc_engine
```
kratos run
```

## TEST
``` 
http://127.0.0.1:8000/sum/47/90333

// output:
{
    "a": "47",
    "b": "90333",
    "sum": "90380"
}
```



