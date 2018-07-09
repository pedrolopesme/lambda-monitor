<h1 align="center">
  AWS Lambda Monitor
  <br>
</h1>

<h4 align="center">AWS Lambda Utilities</h4>

<p align="center">
  <a href="https://travis-ci.org/pedrolopesme/lambda-monitor"> <img src="https://api.travis-ci.org/pedrolopesme/lambda-monitor.svg?branch=master" /></a>
  <a href="https://goreportcard.com/report/github.com/pedrolopesme/lambda-monitor"> <img src="https://goreportcard.com/badge/github.com/pedrolopesme/lambda-monitor" /></a>
  <a href="https://codeclimate.com/github/pedrolopesme/lambda-monitor/maintainability"><img src="https://api.codeclimate.com/v1/badges/10ba198f10121eb45cf4/maintainability" /></a>
  <a href='https://coveralls.io/github/pedrolopesme/lambda-monitor?branch=master'><img src='https://coveralls.io/repos/github/pedrolopesme/lambda-monitor/badge.svg?branch=master' alt='Coverage Status' /></a>
</p>


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build lambda-monitor.

* build: generates binaries
* test: runs all tests
* clean: removes binaries
* fmt: runs gofmt for all go files
* coverage: runs tests coverage build on Coveralls. It expects you to have 
declared COVERALLS_LAMBDAMONITOR_KEY env var.


### Running tests

Tests were write using [Testify](https://github.com/stretchr/testify). In order to run them, just type:

```shell
$ make test
```


### License

[Apache License, Version 2.0](LICENSE.md)