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

Example of current output text:
```owen@pfhor:~/go/src/getcert$ ./getcert www.example.com
Downloading certificate from www.example.com
Connected
3 certificates in chain
Cert # 0
--------
Valid From: "2018-11-28 00:00:00 +0000 UTC"
Expiry:     "2020-12-02 12:00:00 +0000 UTC"
Subject Alt. Names:
	# 0 www.example.org
	# 1 example.com
	# 2 example.edu
	# 3 example.net
	# 4 example.org
	# 5 www.example.com
	# 6 www.example.edu
	# 7 www.example.net


Cert # 1
--------
Valid From: "2013-03-08 12:00:00 +0000 UTC"
Expiry:     "2023-03-08 12:00:00 +0000 UTC"
Subject Alt. Names:
 	- None -


Cert # 2
--------
Valid From: "2006-11-10 00:00:00 +0000 UTC"
Expiry:     "2031-11-10 00:00:00 +0000 UTC"
Subject Alt. Names:
 	- None -
```
And the files you get as a result...
```
owen@pfhor:~/go/src/getcert$ ls -lh
total 4.5M
-rwxr-xr-x 1 owen owen 4.5M Jan 19 20:16 getcert
-rw-rw-r-- 1 owen owen 2.0K Jan 18 12:46 getcert.go
-rw-rw-r-- 1 owen owen  516 Jan 19 20:22 README.md
-rw-r--r-- 1 owen owen 2.6K Jan 19 20:25 www.example.com_00
-rw-r--r-- 1 owen owen 1.7K Jan 19 20:25 www.example.com_01
-rw-r--r-- 1 owen owen 1.4K Jan 19 20:25 www.example.com_02
```
