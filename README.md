# Unix IPC Testing

Testing the performance of Unix IPC.
Adapted from [blog post by Eli Bendersky](https://eli.thegreenplace.net/2019/unix-domain-sockets-in-go/).

## Running

First build
```
go build ./bin/server
go build ./bin/client
```

Run the client and server programs separately
```
./server --unixdomain --msgsize 10000000 &
./client --unixdomain --msgsize 10000000
```

## Some test results

Unix Domain Sockets on an AWS c5n.large instance.

**128 byte messages**
```
50000 pingpongs took 790975853 ns
avg. rt latency 15819 ns
avg. thoughput 0.007536 GB/sec
```

**10,000,000 byte messages**
```
50000 pingpongs took 37768188364 ns
avg. rt latency 755363 ns
avg. thoughput 12.329458 GB/sec
```
