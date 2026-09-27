package products

import(
	"fmt"
	"net/http"
)

func GetAllProducts(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w, "Products Fetched")
}