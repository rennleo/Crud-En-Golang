package main

import(
//"fmt"
"log"
"net/http"
"text/template"
)
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main(){
	http.HandleFunc("/", Inicio)

	log.Println("Servidor corrriendo")
	http.ListenAndServe(":8080",nil)
}
func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"hola develoteca")
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

