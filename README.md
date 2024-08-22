# AVAD: Web Directory Scanner Tool

> Overview

## Avad is a multi-threaded tool designed for scanning directories on web servers. It allows you to send both GET and POST requests to specific URLs or a list of URLs, making it easier to test the availability of various endpoints. The tool is particularly useful for penetration testers and developers who need to check the structure and behavior of web servers.

## Features -

  - **Concurrent Requests:** Perform multiple requests simultaneously with customizable concurrency.
  - **Support for GET and POST:** Sends GET requests and automatically converts them to POST requests with sample data.
  - **Custom URL and Directory Input:** Scan a single URL or a list of URLs and directories.
  - **Comprehensive Output:** Get detailed information about the status and size of the responses.

__Installation__

Clone the repository: ``git clone https://github.com/avyaysec/avad.git``

Navigate to the project directory: ``cd avad``

Build the project: ``go build``

Run the scanner: ``./avad.go -u http://example.com -d /admin``

![avad](https://github.com/user-attachments/assets/553724bc-edf1-4b2d-9567-8491e2ee95df)


Usage
Command-Line Options

    -u string
    Specify a single URL to scan (e.g., http://example.com).

    -U string
    Specify the path to a list of URLs (e.g., urllist.txt).

    -d string
    Specify a single directory to scan (e.g., /admin).

    -D string
    Specify the path to a list of directories (e.g., dirlist.txt).

    -c int
    Set the number of concurrent requests (default is 10).
