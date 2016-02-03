
package main

import (
    //"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	
	http.HandleFunc("/apiPing", apiPing)
    http.ListenAndServe(":8082", nil)

   
}

func apiPing(w http.ResponseWriter, r * http.Request){
	
	userId := r.URL.Query().Get("user")
	
 	//db, err := sql.Open("mysql", "root@/test")
    //if err != nil {
    //    panic(err.Error())
    //}
    //defer db.Close()

	//query, err := db.Query("SELECT maketxt FROM testtable WHERE idtesttable = 1")
	//if err != nil {
	//	panic(err.Error())
	//}

	//var col1 int
	//query.Next();
	//err = query.Scan(&col1);
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer query.Close()


    w.Write([]byte(userId));
}