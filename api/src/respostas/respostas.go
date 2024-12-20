package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statuscode int, dados interface{}) {
	w.WriteHeader(statuscode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func Erro(w http.ResponseWriter, statuscode int, erro error) {
	JSON(w, statuscode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
