
package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"net/http"
	//"log"
	"fmt"
)

func main() {
	
	http.HandleFunc("/testApiPing", apiPing)
	http.HandleFunc("/testApiDB", apiDB)
    http.ListenAndServe(":8082", nil)

   
}

func apiPing(w http.ResponseWriter, r * http.Request){
	
	param1 := r.URL.Query().Get("param1")

    w.Write([]byte(Reverse(param1)));
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}


func apiDB(w http.ResponseWriter, r * http.Request){
	
	userId := r.URL.Query().Get("user")
	
 	db, err := sql.Open("mysql", "root@/test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	stmt, err := db.Prepare("SELECT content FROM testtable WHERE idtesttable = ?")

    res, err := stmt.Query(userId)

	if err != nil {
		panic(err.Error())
	}

	var col1 string
	res.Next();
	err = res.Scan(&col1);

	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Response: ", col1)
}