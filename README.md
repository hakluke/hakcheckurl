# hakcheckurl

Takes a list of URLs and returns their HTTP response codes. This tool is perfect for quickly determining the status of multiple web pages, especially when combined with other tools.

This tool was written to be chained with [hakrawler](https://github.com/hakluke/hakrawler) to easily check the response codes of discovered URLs.

## Features
- **Concurrent Processing**: Utilize multiple threads to speed up the checking process.
- **Configurable Timeout**: Define how long each request should wait before timing out.
- **Retry Mechanism**: Automatically retry failed requests to handle temporary network glitches.

This tool was written to be chained with [hakrawler](https://github.com/hakluke/hakrawler) to easily check the response codes of discovered URLs.

# Installation
```
go install github.com/hakluke/hakcheckurl@latest
```

# Usage
```
-t : Specify the number of threads to use.
-retry : Define how many times to retry failed requests.
-timeout : Set a timeout for each request.
-retry-sleep : Set the duration to sleep between retries.
```
# Sample Usage
```
hakluke~$ assetfinder google.com | hakrawler -plain | hakcheckurl | grep -v 404
200 http://mw1.google.com/transit?
200 http://mw1.google.com/places/
200 http://mw1.google.com/notebook/search?
200 http://mw1.google.com/reader/
200 http://mw1.google.com/views?
200 http://mw1.google.com/sprint_xhtml
200 http://mw1.google.com/sprint_wml
200 http://mw1.google.com/scholar
200 https://area120.google.com/
200 https://area120.google.com/static/js/main.min.js
...
```
