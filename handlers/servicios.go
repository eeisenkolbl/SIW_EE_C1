package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eeisenkolbl/SIW_EE_C1/models"
)

var servicios = []models.Servicio{
	{
		ID:          1,
		Nombre:      "Mantenimiento 24hs",
		Descripcion: "Servicio de mantenimiento disponible las 24 horas para cualquier emergencia en tu departamento.",
		Icono:       "🔧",
	},
	{
		ID:          2,
		Nombre:      "Seguridad",
		Descripcion: "Vigilancia permanente con cámaras de circuito cerrado y personal de seguridad en el edificio.",
		Icono:       "🛡️",
	},
	{
		ID:          3,
		Nombre:      "Áreas Comunes",
		Descripcion: "Piscina, gimnasio, salón de eventos y lavandería comunitaria disponibles para todos los inquilinos.",
		Icono:       "🏊",
	},
	{
		ID:          4,
		Nombre:      "Administración de Pagos",
		Descripcion: "Gestión centralizada de pagos con múltiples opciones: efectivo, transferencia y billeteras electrónicas.",
		Icono:       "💳",
	},
}

func GetServicios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(servicios)
}
