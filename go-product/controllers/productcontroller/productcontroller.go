package productcontroller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Ahmadr127/go-affiliate-web/entities"
	"github.com/Ahmadr127/go-affiliate-web/models/productmodel"
)

var prductModel = productmodel.New()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/product/index.html")
	temp.Execute(w, data)
}

func GetData() string {

	buffer := &bytes.Buffer{}

	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/product/data.html")

	var product []entities.Product
	err := prductModel.FindAll(&product)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"product": product,
	}

	temp.ExecuteTemplate(buffer, "data.html", data)

	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)
	var data map[string]interface{}
	var product entities.Product

	if err != nil {
		data = map[string]interface{}{
			"title": "Tambah Data Product",
			"product": product,
		}
	} else {
		
		err := prductModel.Find(id, &product)
		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"title":     "Edit Data product",
			"product": product,
		}
	}

	temp, _ := template.ParseFiles("views/product/form.html")
	temp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		r.ParseForm()
		var product entities.Product

		product.NamaProduct = r.Form.Get("nama_product")
		product.Harga = r.Form.Get("harga")
		product.JumlahTerjual = r.Form.Get("jumlah_terjual")
		product.Kota = r.Form.Get("kota")
		product.Url = r.Form.Get("url")


		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		var data map[string]interface{}

		if err != nil {
			// insert data
			err := prductModel.Create(&product)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data berhasil disimpan",
				"data":    template.HTML(GetData()),
			}
		} else {
			// mengupdate data
			product.Id = id
			err := prductModel.Update(product)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data berhasil diubah",
				"data":    template.HTML(GetData()),
			}
		}

		ResponseJson(w, http.StatusOK, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	err = prductModel.Delete(id)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"message": "Data berhasil dihapus",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"error": message})
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}






