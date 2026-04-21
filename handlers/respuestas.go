package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/eeisenkolbl/SIW_EE_C1/models"
	"github.com/go-chi/chi/v5"
)

// RespuestaMap is the central store of all chatbot responses, keyed by ID.
// The decision tree in chat.go maps keywords → RespuestaID → text here.
var RespuestaMap = map[int]models.Respuesta{
	0:  {ID: 0, Categoria: "fallback", Texto: "No entendí tu consulta, ¿podés reformularla? Puedo ayudarte con horarios, precios, ubicación, mantenimiento y más."},
	1:  {ID: 1, Categoria: "saludo", Texto: "¡Hola! Soy el asistente virtual de Residencial Guaraní. ¿En qué puedo ayudarte?"},
	2:  {ID: 2, Categoria: "horarios", Texto: "Nuestro horario de atención es de lunes a viernes de 08:00 a 17:00 y sábados de 08:00 a 12:00."},
	3:  {ID: 3, Categoria: "ubicacion", Texto: "Estamos ubicados en Av. Mariscal López 1234, Asunción, Paraguay."},
	4:  {ID: 4, Categoria: "precios", Texto: "Nuestros departamentos van desde Gs. 2.500.000/mes (monoambiente) hasta Gs. 5.000.000/mes (2 dormitorios). Contactanos para conocer disponibilidad actual."},
	5:  {ID: 5, Categoria: "pagos", Texto: "Aceptamos pago en efectivo, transferencia bancaria y giros desde billeteras electrónicas. El pago se realiza del 1 al 5 de cada mes."},
	6:  {ID: 6, Categoria: "mantenimiento", Texto: "Para solicitudes de mantenimiento, podés comunicarte al (021) 555-1234 o enviar un mensaje por este chat. Atendemos urgencias las 24 horas."},
	7:  {ID: 7, Categoria: "requisitos", Texto: "Para alquilar necesitás: cédula de identidad, comprobante de ingresos, referencia personal y un depósito de garantía equivalente a 2 meses de alquiler."},
	8:  {ID: 8, Categoria: "areas_comunes", Texto: "Contamos con piscina, gimnasio, salón de eventos, lavandería comunitaria y estacionamiento para inquilinos."},
	9:  {ID: 9, Categoria: "contacto", Texto: "Podés contactarnos al (021) 555-1234, por WhatsApp al +595 981 123456, o por correo a info@residencialguarani.com.py."},
	10: {ID: 10, Categoria: "mascotas", Texto: "Se permiten mascotas pequeñas (hasta 10 kg) con un depósito adicional de Gs. 500.000. Consultá las normas de convivencia en administración."},
	11: {ID: 11, Categoria: "servicios_incluidos", Texto: "El alquiler incluye agua, recolección de basura y mantenimiento de áreas comunes. Internet y electricidad corren por cuenta del inquilino."},
	99: {ID: 99, Categoria: "easter_egg", Texto: "Jaja si atendemos 😄"},
}

// GetRespuestas godoc
// @Summary      Listar todas las respuestas
// @Description  Retorna todas las respuestas del árbol de decisión del chatbot
// @Tags         respuestas
// @Produce      json
// @Success      200  {array}   models.Respuesta
// @Router       /api/respuestas [get]
func GetRespuestas(w http.ResponseWriter, r *http.Request) {
	list := make([]models.Respuesta, 0, len(RespuestaMap))
	for _, v := range RespuestaMap {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

// GetRespuestaByID godoc
// @Summary      Obtener respuesta por ID
// @Description  Retorna la respuesta del chatbot correspondiente al ID indicado
// @Tags         respuestas
// @Produce      json
// @Param        id   path      int  true  "ID de la respuesta"
// @Success      200  {object}  models.Respuesta
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/respuestas/{id} [get]
func GetRespuestaByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID inválido"})
		return
	}
	resp, ok := RespuestaMap[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Respuesta no encontrada"})
		return
	}
	json.NewEncoder(w).Encode(resp)
}
