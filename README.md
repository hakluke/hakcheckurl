你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# hakcheckurl
Takes a list of URLs and returns their HTTP response codes

This tool was written to be chained with [hakrawler](https://github.com/hakluke/hakrawler) to easily check the response codes of discovered URLs.

# Installation
```
go get github.com/hakluke/hakcheckurl
```

# Usage
- `-t 100` use 100 threads

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
