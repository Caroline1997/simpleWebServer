# 服务计算Golang作业--编程web服务程序类似[cloudgo应用](https://github.com/pmlpml/golang-learning/tree/master/web)
## 框架
使用[gorilla/mux](https://github.com/gorilla/mux) + [codegangsta/negroni](https://github.com/urfave/negroni)

Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.

Negroni is an idiomatic approach to web middleware in Go. It is tiny, non-intrusive, and encourages use of net/http Handlers.
Negroni tries to play well with all of them by fully supporting net/http. 

## 使用curl测试
go run main.go
```
caro@ubuntu:~/gopath/src/hw/web$ go run main.go
[negroni] listening on :9999
[negroni] 2017-11-08T14:41:37+08:00 | 200 | 	 46.514µs | localhost:9999 | GET / 
[negroni] 2017-11-08T14:41:53+08:00 | 200 | 	 49.95µs | localhost:9999 | GET /student/15331157 
[negroni] 2017-11-08T14:42:05+08:00 | 200 | 	 120.608µs | localhost:9999 | GET /teacher/pan 
```
three tests
```
caro@ubuntu:~$ curl -v http://localhost:9999/
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 9999 (#0)
> GET / HTTP/1.1
> Host: localhost:9999
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 08 Nov 2017 06:41:37 GMT
< Content-Length: 26
< Content-Type: text/plain; charset=utf-8
< 
Welcome to the home page!
* Connection #0 to host localhost left intact
caro@ubuntu:~$ curl -v http://localhost:9999/student/15331157
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 9999 (#0)
> GET /student/15331157 HTTP/1.1
> Host: localhost:9999
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 08 Nov 2017 06:41:53 GMT
< Content-Length: 18
< Content-Type: text/plain; charset=utf-8
< 
Student: 15331157
* Connection #0 to host localhost left intact
caro@ubuntu:~$ curl -v http://localhost:9999/teacher/pan
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 9999 (#0)
> GET /teacher/pan HTTP/1.1
> Host: localhost:9999
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 08 Nov 2017 06:42:05 GMT
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
< 
Teacher: pan
* Connection #0 to host localhost left intact
```
## 使用ab测试
```
//对http://localhost:9999/student/15331157 进行共1000次请求，一次同时100个并发请求
caro@ubuntu:~$ ab -n 1000 -c 100 http://localhost:9999/student/15331157
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            9999

Document Path:          /student/15331157
Document Length:        18 bytes

Concurrency Level:      100
//time for the whole test
Time taken for tests:   0.231 seconds
//number of completed requests
Complete requests:      1000
Failed requests:        0
Total transferred:      135000 bytes
HTML transferred:       18000 bytes
//平均每秒处理请求4322个请求
Requests per second:    4322.75 [#/sec] (mean)
//平均每个请求处理时间为23ms，将100次并发请求作为一个整体
Time per request:       23.133 [ms] (mean)
//平均每个并发请求处理的时间为0.23ms=23ms/100
Time per request:       0.231 [ms] (mean, across all concurrent requests)
Transfer rate:          569.89 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.9      0       5
Processing:     0   21  14.3     18      71
Waiting:        0   21  14.3     17      71
Total:          0   22  14.3     18      71

Percentage of the requests served within a certain time (ms)
  //在共1000个请求中，有50%在18ms内完成
  50%     18
  //在共1000个请求中，有66%在24ms内完成
  66%     24
  75%     28
  80%     31
  90%     38
  95%     57
  98%     63
  99%     71
  //共1000个请求,71ms完成
 100%     71 (longest request)
 ```
 ab command parameter
 ```
Usage: ab [options] [http[s]://]hostname[:port]/path  

options:
//总的请求数
-n requests Number of requsts to perform
//一次同时并发的请求数
-c concurrncy Number of multiple requests to make
 ``` 