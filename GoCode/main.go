package main
import(
	DB "JSMPJ_vue_go/GoCode/src/system/db"
	"net/http"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"reflect"

)
type User struct {
	Id int64 `json:"id"`
	FName string `json:"fname"`
	LName string `json:"lname"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type DataCollection struct{
	DC []User `json:"items"`
}
var arr[] string
func main(){
	_, err := DB.Connect()
	if err != nil{
		return
	}

	fmt.Println("JSMPJ Database prctice")
	db,err := sql.Open("mysql","test1:test1@tcp(localhost:3306)/test1")
    if err != nil{
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Successfully connected to database")
	
	
	// insert, err := db.Query("INSERT INTO users1 VALUES('2','Anna','Rana','er.anna03@gmail.com','123#')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	fmt.Println("Successfully inserted into user tables")

	rows, err := db.Query("SELECT * from users1")
	if err != nil{
		panic(err.Error())
	}
	result1:=DataCollection{}
	defer rows.Close()
	result := DataCollection{}
	for rows.Next(){
		data := User{}
		err2:= rows.Scan(&data.Id,&data.FName,&data.LName,&data.Email,&data.Password)
        if err2 != nil{
			panic(err2)
		}
		result.DC = append(result.DC ,data)
		result1.DC = append(result.DC ,data)
		//  arr[0]=result.DC[0]
		fmt.Println(reflect.TypeOf(result.DC))
		fmt.Println(data)
	
	}
	fmt.Println(result1.DC)
    // for results.Next(){
	// 	var user User
	// 	err = results.Scan(&user.FName)
	// 	// err = results.Scan(&user.LName)
	// 	if err != nil{
	// 		panic(err.Error())
	// 	}
	// 	arr[0]= user.FName
	// 	fmt.Println(user.FName)
		
	// 	// fmt.Println(user.LName)
	// }
	r := mux.NewRouter()
	methods := handlers.AllowedHeaders([]string{"GET","POST","PUT","DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/",handler).Methods("GET")
	r.HandleFunc("/about",aboutHandler).Methods("GET")
	r.HandleFunc("/contact",contactHandler).Methods("GET")
	r.HandleFunc("/login",loginHandler).Methods("GET")
	http.Handle("/",r)
	http.ListenAndServe(":8080", handlers.CORS(methods,origins)(r))
}
func handler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"Home\": \"Page\"}"))
}
func aboutHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"About\": \"Page\"}"))
}
func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"Contact\": \"Page\"}"))
}
func loginHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"Login\": \"Page\"}"))
}
