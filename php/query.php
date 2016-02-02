<?php
$time = microtime(true);
$mysqli = new mysqli("127.0.0.1", "root", "", "test");
for ($i = 0; $i < 100; $i++) {
	$res = $mysqli->query("SELECT idtestTable FROM testtable WHERE MAKETXT = 'BMW'");
    $row = $res->fetch_assoc();
}
echo microtime(true) - $time;