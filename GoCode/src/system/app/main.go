package main
import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)
func main(){
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
