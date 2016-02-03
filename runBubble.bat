@echo off
Echo Bubble Execution - 1 millon cycles
echo Node time:
node node.js/testBubble.js
echo PHP Time:
php php/testBubble.php
cd  java
echo  ms 
echo Java Time:
java TestBubble
cd ..
echo Golang Time:
go run golang/src/testBubble.go
