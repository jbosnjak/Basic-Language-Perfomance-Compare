# Basic-Language-Perfomance-Compare

This project is a basic tool to test different language performance. At this moment we are testing 4 languages: Java, PHP, Golang and Node.js.

##Prerequisites

This project only includes source files, so you need to install all the following software to test it.

### Mysql

The test needs a mysql database install. 

	sudo yum install mysql-server
	sudo service mysqld start

You need to create a schema name "test"

Then you have to run the script **createTable.sql** in the `mysql/` directory.

**All the connections are hardcoded in the src files**, if you need to change it, do it in the code.
	
here is an example:

	var con = mysql.createConnection({
    host : "localhost",
    user : "root",
    password : "",
    database : "test"
	});


###PHP

For centos/RedHat

	sudo yum install php php-mysql

We recommend php 5.5+.

###Apache

	sudo yum install httpd
	sudo service httpd start

###JAVA

First install java, only if you don´t already have it:

	sudo yum install java-1.7.0-openjdk

Then, you have to build the project:

	cd java/testApi
	gradlew build


###Node.js

	wget https://nodejs.org/dist/v4.2.6/node-v4.2.6-linux-x64.tar.xz
	sudo tar --strip-components 1 -xzvf node-v* -C /usr/local

###Golang

	yum install golang
	yum install hg

###Apache Jmeter

	wget http://mirrors.dcarsat.com.ar/apache//jmeter/binaries/apache-jmeter-2.13.tgz
	tar zxvf apache-jmeter-2.13.tgz
	cd apache-jmeter-2.11/bin/
	./jmeter


##Tests included

There are 4 basic tests included in this project:


*   Bubble: we generate an array and call 1 millon times the function bubble
*   DBQuery: The programs connects 100 times to mysql database and executes a simple query
*	API Rest Ping: The programs expose an http path that receives a parameter as a string, reverse it and returns it
*	API Rest DB: The programs expose an http path that receives a parameter, executes a query to the mysql database and returns a string

The first two test are executed without a http server and the results are measure in milliseconds.

For the last two tests, a http server is needed. For this cases `apache-jmeter` is needed. In the `test/jmeter` directory you can find a test sample to run it.

###How to run the tests

After all services are up, the database is ready and loaded, you can start the test:

####Test Bubble

	./runBubble.sh
	
or
	
	node node.js/testBubble.js
	go run golang/testBubble.go
	java java/TestBubble
	php php/testBubble.php
	
####test DBQuery

	node node.js/testDBQuery.js
	go run golang/testDBQuery.go
	php php/testDBQuery.php

java does not have a test for this

####test API Rest Ping & API Rest DB 

First set up the servers:

	node node.js/testApi.js
	go run golang/testApi.go
	java/testApi/gradlew bootRun
	sudo service httpd start
	
For php is important that the document root points to the `php` directory.

Then just open the apache-jmeter, load the test project in `test/jmeter/listUsers.jmx` and run it.


	 
