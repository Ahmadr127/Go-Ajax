package main

import (
	"net/http"

	"github.com/Ahmadr127/go-affiliate-web/controllers/productcontroller"
	
)

func main() {
	http.HandleFunc("/", productcontroller.Index)
	http.HandleFunc("/product/get_form", productcontroller.GetForm)
	http.HandleFunc("/product/store", productcontroller.Store)
	http.HandleFunc("/product/delete", productcontroller.Delete)
	


	http.ListenAndServe(":8000", nil)

}
