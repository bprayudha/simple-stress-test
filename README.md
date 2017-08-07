### Description
This is a simple stress test using [ApacheBench](http://httpd.apache.org/docs/current/programs/ab.html). This test is performing 100,000 requests with 100 concurrency.

### Request
```
GET /
```

### Response
```javascript
{
  "hello": "world"
}
```

### Command to Test
```
$ ab -klc 100 -n 100000 http://127.0.0.1:3000/
```

### Result
|                    |   Laravel  | Ruby on Rails | Express JS | Play (Scala) |  Golang  |
|--------------------|-----------:|--------------:|-----------:|-------------:|---------:|
| Request per second |         21 |           848 |       9324 |        18613 |    36443 |
| Time to complete   | 4660.058 s |     117.859 s |   10.693 s |      5.372 s |  2.744 s |
| Time per request   |  46.601 ms |      1.179 ms |   0.107 ms |     0.054 ms | 0.027 ms |
