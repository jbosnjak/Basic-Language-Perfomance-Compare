<?php
$init = round(microtime(true) * 1000);
//opens 100 connections to the database
for ($i = 0; $i < 100; $i++) {
	$mysqli = new mysqli("127.0.0.1", "root", "", "test");
	$res = $mysqli->query("SELECT idtestTable FROM testtable WHERE content = 'a test content with BMW'");
    $row = $res->fetch_assoc();
}

$end= round(microtime(true) * 1000);
echo ($end - $init);