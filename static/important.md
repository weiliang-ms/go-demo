> 服务端的最大连接数取决于文件描述符，而非支持的最大监听端口数

	https://stackoverflow.com/questions/2332741/what-is-the-theoretical-maximum-number-of-open-tcp-connections-that-a-modern-lin

	If a client has many connections to the same port on the same destination, 
	then three of those fields will be the same - only source_port varies to differentiate the different connections. 
	Ports are 16-bit numbers, therefore the maximum number of connections any given client can have to any given host port is 64K.

	However, multiple clients can each have up to 64K connections to some server's port, and if the server has multiple ports or either is multi-homed then you can multiply that further.

	So the real limit is file descriptors. Each individual socket connection is given a file descriptor, so the limit is really the number of file descriptors that the system has been configured to allow and resources to handle. The maximum limit is typically up over 300K, but is configurable e.g. with sysctl.



> nginx最大打开文件个数

	(worker_processes * worker_connections * 2) + (shared libs, log files, event pool) = max open files

> keepalive工作原理

	Keep Alive's work between requests.

	When you download a webpage it downloads the HTML page and discovers it needs another 20 resources say (CSS files, javascript files, images, fonts... etc.).
	
	Under HTTP/1.1 you can only request one of these resources at once so typically the web browser fires up another 5 connections (giving 6 in total) and requests 6 of those 20 resources. Then it requests the remaining 14 resources as those connections free up. Yes keep-alives help in between those requests but that's not its only use as we'll discuss below. The overhead of setting up those connections is small but noticeable and there is a delay in only being able to request 6 resources of those 20 at a time. This is why HTTP/1.1 is inefficient for today's usage of the web where a typical web page is made up of 100 resources.
	
	Under HTTP/2 we can fire off all 20 requests at once on the same connection so some good gains there. And yes technically you don't really benefit from keep-alives in between those as connection is still in use until they all arrive - though still benefit from small delay between first HTML request and the other 20.
	
	However after that initial load there are likely to be more requests. Either because you are browsing around the site or because you interact with the page and it makes addition XHR api calls. Those will benefit from keep-alives whether on HTTP/1.1 or HTTP/2.
	
	So HTTP/2 doesn't negate need for keep-alives. It negates need for multiple connections (amongst other things).
	
	So the answer is to always use keep-alives unless you've a very good reason not to. And what type of benchmarking are you doing to say it makes no difference?

> nginx-keepalive

	Please keep in mind that keepalive is a feature of HTTP 1.1, NGINX uses HTTP 1.0 per default for upstreams.

	nginx与客户端采用HTTP1.1，具有keepalive能力
	nginx与上游服务采用HTTP1.0，不具有keepalive能力


