var express = require("express");
var app = express();
var fs = require("fs");

var mysql = require("mysql");

app.get('/listUsers', function(req, res) {
    var con = mysql.createConnection({
	host : "localhost",
	user : "root",
	password : "",
	database : "test"
    });
    con.connect(function(err) {
	if (err) {
	    console.log('Error connecting to Db');
	    return;
	}
    });
    
    con.query("SELECT maketxt FROM testtable WHERE idtesttable  = "+req.query.user, function(err, row) {
	res.send(row);
    });
    
    con.end(function(err) {
    });
    
})

var server = app.listen(8081, function() {

    var host = server.address().address
    var port = server.address().port

    console.log("Example app listening at http://%s:%s", host, port)

})