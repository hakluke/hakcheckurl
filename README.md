# hakcheckurl
Takes a list of URLs and returns their HTTP response codes

This tool was written to be chained with [hakrawler](https://github.com/hakluke/hakrawler) to easily check the response codes of discovered URLs.

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
