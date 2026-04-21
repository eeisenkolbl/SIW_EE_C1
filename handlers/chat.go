package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"

	"github.com/eeisenkolbl/SIW_EE_C1/models"
)

// decisionTree maps keyword sets to a RespuestaID.
// Add or reorder entries here to change chatbot behavior.
var decisionTree = []models.FAQ{
	{Keywords: []string{"horario", "hora", "atencion", "abierto", "abre", "cierra"}, RespuestaID: 2},
	{Keywords: []string{"ubicacion", "direccion", "donde", "llegar", "mapa"}, RespuestaID: 3},
	{Keywords: []string{"precio", "alquiler", "costo", "cuanto", "mensualidad", "renta"}, RespuestaID: 4},
	{Keywords: []string{"pago", "pagar", "transferencia", "efectivo", "tarjeta", "deposito"}, RespuestaID: 5},
	{Keywords: []string{"mantenimiento", "reparacion", "arreglo", "roto", "problema", "fuga", "luz"}, RespuestaID: 6},
	{Keywords: []string{"requisito", "documento", "contrato", "garantia", "alquilar", "inquilino"}, RespuestaID: 7},
	{Keywords: []string{"area", "comun", "piscina", "gimnasio", "estacionamiento", "salon", "lavanderia"}, RespuestaID: 8},
	{Keywords: []string{"contacto", "telefono", "llamar", "email", "correo", "whatsapp"}, RespuestaID: 9},
	{Keywords: []string{"mascota", "perro", "gato", "animal"}, RespuestaID: 10},
	{Keywords: []string{"incluido", "servicio", "agua", "internet", "basura", "expensa", "gasto"}, RespuestaID: 11},
	// Easter egg
	{Keywords: []string{"urgencia", "urgencias"}, RespuestaID: 99},
}

var greetingKeywords = []string{"hola", "buenas", "buenos", "buen dia", "buenas tardes", "buenas noches"}

func normalize(s string) string {
	s = strings.ToLower(s)
	var result strings.Builder
	for _, r := range norm.NFD.String(s) {
		if !unicode.Is(unicode.Mn, r) {
			result.WriteRune(r)
		}
	}
	return strings.TrimSpace(result.String())
}

func respond(w http.ResponseWriter, id int) {
	resp := RespuestaMap[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ChatResponse{ID: resp.ID, Response: resp.Texto})
}

// PostChat godoc
// @Summary      Enviar mensaje al chatbot
// @Description  Procesa el mensaje del usuario y retorna la respuesta FAQ más relevante usando el árbol de decisión
// @Tags         chat
// @Accept       json
// @Produce      json
// @Param        message  body      models.ChatRequest   true  "Mensaje del usuario"
// @Success      200      {object}  models.ChatResponse
// @Failure      400      {object}  models.ChatResponse
// @Router       /api/chat [post]
func PostChat(w http.ResponseWriter, r *http.Request) {
	var req models.ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, 0)
		return
	}

	normalized := normalize(req.Message)

	for _, g := range greetingKeywords {
		if strings.Contains(normalized, g) {
			respond(w, 1)
			return
		}
	}

	words := strings.Fields(normalized)
	bestID := 0
	bestScore := 0

	for _, faq := range decisionTree {
		score := 0
		for _, keyword := range faq.Keywords {
			for _, word := range words {
				if strings.Contains(word, keyword) {
					score++
					break
				}
			}
		}
		if score > bestScore {
			bestScore = score
			bestID = faq.RespuestaID
		}
	}

	respond(w, bestID)
}
