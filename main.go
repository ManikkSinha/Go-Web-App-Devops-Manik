package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/courses", coursesHandler)

	fmt.Println("🚀 Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to Manik's DevOps Project</h1><p>This is a demo Go web app used for learning Docker, Kubernetes, and CI/CD pipelines.</p>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About This Project</h1><p>This project is created and customized by Manik Sinha. It is a sample Go web application used to demonstrate DevOps practices.</p>")
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Learning Modules</h1><ul><li>Go Web Application Basics</li><li>Docker & Containerization</li><li>Kubernetes Deployment</li><li>CI/CD Pipeline Integration</li></ul>")
}
