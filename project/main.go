package main



import(
	"Project/src/routes"
	"net/http"
)

func main(){
	port := ":8080"
	routes.SetUpRoutes()
	http.ListenAndServe(port,nil)
}