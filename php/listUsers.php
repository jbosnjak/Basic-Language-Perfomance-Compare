<?php
$userId = $_GET ["user"];

$mysqli = new mysqli ( "127.0.0.1", "root", "", "test" );
$res = $mysqli->query ( "SELECT maketxt FROM testtable WHERE idtesttable = " . $userId );
$row = $res->fetch_assoc ();
echo json_encode($row);