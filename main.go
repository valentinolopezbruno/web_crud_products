package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var temp = template.Must(template.ParseGlob("front/*"))

type Product struct {
	Id          int
	Name        string
	Description string
	Img         string
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)
	fmt.Println("SERVIDOR INICIADO")
	http.ListenAndServe(":8080", nil)
}

func connectDB() (conn *sql.DB) {
	Name := "productss"
	conn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/"+Name) //nameBD == productss

	if err != nil {
		panic(err.Error())
	}
	return conn
}

func Home(w http.ResponseWriter, r *http.Request) {

	connect := connectDB()
	getProduct, err := connect.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	sliceProduct := []Product{}

	for getProduct.Next() {
		var id int
		var name, description, img string
		err = getProduct.Scan(&id, &name, &description, &img)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Img = img

		sliceProduct = append(sliceProduct, product)
	}

	temp.ExecuteTemplate(w, "home", sliceProduct)
}

func Create(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		img := r.FormValue("img")

		connect := connectDB()
		createProduct, err := connect.Prepare("INSERT INTO products(name,description,img) VALUES(?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		createProduct.Exec(name, description, img)

		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
	temp.ExecuteTemplate(w, "create", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	getIdProduct := r.URL.Query().Get("id")
	connect := connectDB()
	editProduct, err := connect.Query("SELECT * FROM products WHERE id=?", getIdProduct)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for editProduct.Next() {
		var id int
		var name, description, img string
		err = editProduct.Scan(&id, &name, &description, &img)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Img = img
	}

	fmt.Println(product)
	temp.ExecuteTemplate(w, "edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		img := r.FormValue("img")

		connect := connectDB()
		updateProduct, err := connect.Prepare("UPDATE products SET name=?, description=?, img=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		updateProduct.Exec(name, description, img, id)

		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	connect := connectDB()
	getIdProduct := r.URL.Query().Get("id")
	fmt.Println("ID DEL PRODUCTO A ELIMINAR: ", getIdProduct)

	deleteProduct, err := connect.Prepare("DELETE FROM products WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(getIdProduct)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
