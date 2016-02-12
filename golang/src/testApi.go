package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/testApiPing", apiPing)
	r.GET("/testApiDB", apiDB)
    r.Run(":8082")
}

func apiDB(c *gin.Context){
	userId := c.Query("user")
	
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

	c.String(http.StatusOK, "Response: %s", col1)
}


func apiPing(c *gin.Context){
	param1 := c.Query("param1")
	responseMessage := Reverse(param1)
    c.String(http.StatusOK, responseMessage )
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/*func apiDB(w http.ResponseWriter, r * http.Request){
	
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
	r.Header.Set("Connection", "close")
}*/