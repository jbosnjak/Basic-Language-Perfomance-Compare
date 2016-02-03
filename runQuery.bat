@echo off
echo validate that your mysql db is up and running
Echo DB Query Execution - 100 connections
echo Node time:
node node.js/testDBQuery.js
echo PHP Time:
php php/testDBQuery.php
echo  ms 
rem cd  java
rem echo Java Time:
rem java TestBubble
rem cd ..
echo Golang Time:
go run golang\src\testDBQuery.go
