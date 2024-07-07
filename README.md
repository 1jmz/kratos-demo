## Start consul
```
consul agent -dev
```

## Start calc_engine/calc_engine micro_service
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



