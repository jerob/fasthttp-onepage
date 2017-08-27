# fasthttp-onepage

Webserver for one page only.

A great way to save bandwidth.

Test with AB :

Concurrency Level:      20
Complete requests:      5000

## FastHTTP-onepage :

Command : docker run -p 8000:80 -v $($pwd)/file:/file:ro -d jerob/fasthttp-onepage
Time taken for tests:   2.717 seconds
Total transferred:      1455000 bytes

## Nginx :

Command : docker run -p 8000:80 -v $($pwd)/file:/usr/share/nginx/html/index.html:ro -d nginx
Time taken for tests:   7.480 seconds
Total transferred:      1850000 bytes

## Apache:

Command : docker run -dit --name my-apache-app -p 8000:80 -v $($pwd)/file:/usr/local/apache2/htdocs/index.html:ro httpd:2.4
Time taken for tests:   4.382 seconds
Total transferred:      1915000 bytes
