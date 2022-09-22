package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	directory := http.Dir("./public")
	fileServer := http.FileServer(directory)

    router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	// router
	// get
	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/form-add-project", getAddProject).Methods("GET")
	router.HandleFunc("/contact-me", getContactMe).Methods("GET")
	router.HandleFunc("/project/{projectId}", getProjectDetail).Methods("GET")
	// post
	router.HandleFunc("/add-project", postAddProject).Methods("POST")

	fmt.Println("running on port 5000")
	http.ListenAndServe("localhost:5000", router)

}

func getHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]string {
		"test": "aduhdeh",
	}
	
	var view, err = template.ParseFiles("views/index.html")	
	if err != nil {
		panic(err.Error())
	}
	view.Execute(w, data)
}

func getAddProject(w http.ResponseWriter, r *http.Request) {
	var view, err = template.ParseFiles("views/project.html")	
	if err != nil {
		panic(err.Error())
	}
	view.Execute(w, nil)
}

func getContactMe(w http.ResponseWriter, r *http.Request) {
	var view, err = template.ParseFiles("views/contact.html")	
	if err != nil {
		panic(err.Error())
	}
	view.Execute(w, nil)
}

func getProjectDetail(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		title = "Default"
	}
	data := map[string]interface{} {
		"title": title,
	}
	var view,err = template.ParseFiles("views/projectDetail.html")
	if err != nil {
		panic(err.Error())
	}
	view.Execute(w, data)

}

func postAddProject(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	data := map[string]interface{} {
		"name": r.PostForm["name"],
		"description": r.PostForm["description"],
		"start-date": r.PostForm["start-date"],
		"end-date": r.PostForm["end-date"],
		"checkbox": r.PostForm["checkbox"],
	}



	fmt.Println(data)

	http.Redirect(w,r,"/form-add-project", http.StatusFound)
}