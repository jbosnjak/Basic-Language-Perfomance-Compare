# Basic-Language-Perfomance-Compare

This project is a basic tool to test different language performance. At this moment we are testing 4 languages: Java, PHP, Golang and Node.js.

#Prerequisites

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
	
#Tests included


	 
