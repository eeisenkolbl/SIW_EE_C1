package models

type Departamento struct {
	ID               int      `json:"id"`
	Nombre           string   `json:"nombre"`
	Precio           int      `json:"precio"`
	PrecioFormateado string   `json:"precioFormateado"`
	Superficie       string   `json:"superficie"`
	Dormitorios      int      `json:"dormitorios"`
	Banos            int      `json:"banos"`
	Caracteristicas  []string `json:"caracteristicas"`
	Imagen           string   `json:"imagen"`
}

type Servicio struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Icono       string `json:"icono"`
}

// Respuesta represents a single chatbot FAQ response identified by ID.
type Respuesta struct {
	ID        int    `json:"id"`
	Categoria string `json:"categoria"`
	Texto     string `json:"texto"`
}

// FAQ maps a set of keywords to a response ID in the decision tree.
type FAQ struct {
	Keywords    []string
	RespuestaID int
}

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	ID       int    `json:"id"`
	Response string `json:"response"`
}
