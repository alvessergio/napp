package routes

import (
	"log"
	"net/http"

	"github.com/alvessergio/pan-integrations/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
