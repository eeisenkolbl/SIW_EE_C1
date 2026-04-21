package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eeisenkolbl/SIW_EE_C1/models"
)

var departamentos = []models.Departamento{
	{
		ID:               1,
		Nombre:           "Ed. SkyTower — 1 Dormitorio",
		Precio:           950,
		PrecioFormateado: "U$S 950/mes",
		Superficie:       "57 m²",
		Dormitorios:      1,
		Banos:            1,
		Caracteristicas:  []string{"Balcón", "Estacionamiento", "Pileta (piso 6)", "Gimnasio y spa", "Seguridad 24hs"},
		Imagen:           "https://cdn2.infocasas.com.uy/repo/img/8ccf59862cad490aee5b79734ed9eb6ed3ab7f0b.jpeg",
	},
	{
		ID:               2,
		Nombre:           "Villa Morra — 2 Dormitorios",
		Precio:           5500000,
		PrecioFormateado: "Gs. 5.500.000/mes",
		Superficie:       "80 m²",
		Dormitorios:      2,
		Banos:            2,
		Caracteristicas:  []string{"Balcón con parrilla", "Cocina amoblada", "Gimnasio equipado", "Pileta con solarium", "Estacionamiento"},
		Imagen:           "https://cdn2.infocasas.com.uy/repo/img/a94382b490395961c0bd6fea010b135f0159eca1.jpeg",
	},
	{
		ID:               3,
		Nombre:           "Forvm Molas López — 2 Dormitorios",
		Precio:           1200,
		PrecioFormateado: "U$S 1.200/mes",
		Superficie:       "80 m²",
		Dormitorios:      2,
		Banos:            2,
		Caracteristicas:  []string{"Amoblado", "Living con balcón", "Cocina equipada", "Pileta con solarium", "Lavandería en edificio"},
		Imagen:           "https://cdn1.infocasas.com.uy/repo/img/fd1aadeb66f8273d94707f39888825121b1589d8.jpg",
	},
	{
		ID:               4,
		Nombre:           "Carmelitas — 2 Dormitorios",
		Precio:           850,
		PrecioFormateado: "U$S 850/mes",
		Superficie:       "85 m²",
		Dormitorios:      2,
		Banos:            2,
		Caracteristicas:  []string{"A estrenar", "Balcón grande con parrilla", "Gimnasio", "Pileta con solarium", "Seguridad 24hs"},
		Imagen:           "https://cdn1.infocasas.com.uy/repo/img/7686c2de6221988b120f1ce0433b95a4f7337456.jpg",
	},
}

func GetDepartamentos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(departamentos)
}
