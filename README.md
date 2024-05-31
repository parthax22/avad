This Go program is a web directory scanner designed to concurrently send HTTP GET requests to specified URLs and directories. Additionally, it converts each GET request into a POST request and sends it with example data. The program provides insights into the status codes and content sizes of the responses.

Features:

  >Concurrent scanning of URLs and directories
  >Conversion of GET requests to POST requests
  >Output of response details for each request

Usage:
go run avad.go -u http://example.com -d /admin -c 10

Output Example:
URL: http://example.com, Directory: /admin, Status: 200 OK, Size: 1234
Converted POST Request: URL: http://example.com, Directory: /admin, Status: 200 OK, Size: 5678
URL: http://example.com, Directory: /test, Status: 404 Not Found, Size: -1
Converted POST Request: URL: http://example.com, Directory: /test, Status: 404 Not Found, Size: 0
URL: http://example.net, Directory: /admin, Status: 301 Moved Permanently, Size: 3456
Converted POST Request: URL: http://example.net, Directory: /admin, Status: 301 Moved Permanently, Size: 7890


This program provides a versatile tool for analyzing web directories, offering both GET and POST request functionalities.
Its concurrent design ensures efficient scanning of multiple URLs and directories simultaneously, making it suitable for various testing and analysis tasks.

