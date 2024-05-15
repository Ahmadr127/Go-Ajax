package productmodel

import (
	"database/sql"

	"github.com/Ahmadr127/go-affiliate-web/config"
	"github.com/Ahmadr127/go-affiliate-web/entities"
)


// Digunakan untuk membuat model method (API) yang akan digunakan

type ProductModel struct {
	db *sql.DB
}

func New() *ProductModel {
	db, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &ProductModel{db: db}
}

func (m *ProductModel) FindAll(product *[]entities.Product) error {
	rows, err := m.db.Query("select * from product")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Product
		rows.Scan(
			&data.Id,
			&data.NamaProduct,
			&data.Harga,
			&data.JumlahTerjual,
			&data.Kota,
			&data.Url)
		*product = append(*product, data)
	}

	return nil
}

func (m *ProductModel) Create(product *entities.Product) error {
	result, err := m.db.Exec("insert into product (nama_product, harga, jumlah_terjual, kota, url) values(?,?,?,?,?)",
		product.NamaProduct, product.Harga, product.JumlahTerjual, product.Kota, product.Url)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	product.Id = lastInsertId
	return nil
}

// scan data berdasarkan id
func (m *ProductModel) Find(id int64, product *entities.Product) error {
	return m.db.QueryRow("select * from product where id = ?", id).Scan(
		&product.Id,
		&product.NamaProduct,
		&product.Harga,
		&product.JumlahTerjual,
		&product.Kota,
		&product.Url)
}
func (m *ProductModel) Update(product entities.Product) error {

	_, err := m.db.Exec("update product set nama_product = ?, harga = ?, jumlah_terjual = ?, kota = ?, url = ? where id = ?",
		product.NamaProduct, product.Harga, product.JumlahTerjual, product.Kota, product.Url, product.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *ProductModel) Delete(id int64) error {
	_, err := m.db.Exec("delete from product where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}


