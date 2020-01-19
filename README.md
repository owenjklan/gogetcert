# Go, Get Cert! #
This is one of my first programs written in Go. Yes, it's a
tiny toy program. However, I find it very useful in my day job so
being able to pull it down off github is useful for me.

You use it like:
```
./getcert server
```

Where server is currently assumed to be an HTTPS server, listening
on port 443.

What you should get is a collection of one or more certificates,
representing each certificate from the chain of trust that the
server has sent. These will be named like *www.example.com_XX*.