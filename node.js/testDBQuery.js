var mysql = require("mysql");

var starttime = new Date().getTime();
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
    //console.log('Connection established');
});

var c = 0;
//opens 100 connections to the database
for (var i = 0; i < 100; i++) {
    con.query("SELECT idtestTable FROM testtable WHERE content = 'a test content with BMW'", function(err, row) {
	if (row.length!= 123){
	    console.log("error");
	}
	c++;
	if (c > 99) {
	    console.log((new Date().getTime() - starttime) + " ms");
	}

    });
}

con.end(function(err) {
    // The connection is terminated gracefully
    // Ensures all previously enqueued queries are still
    // before sending a COM_QUIT packet to the MySQL server.
});