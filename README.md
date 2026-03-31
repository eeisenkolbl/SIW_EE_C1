# Residencial Guaraní — Landing Page

Proyecto para la materia **Sistemas de Interfaz Web** por Ernesto Eisenkolbl.

Landing page con chatbot de FAQs para un residencial de departamentos en Asunción, Paraguay.

---

## Cómo ejecutar

### Opción 1 — Docker (nginx)

Requiere [Docker](https://docs.docker.com/get-docker/) instalado.

```bash
docker compose up
```

Abrir en el navegador: `http://localhost:8085`

Para detener: `Ctrl+C`

---

### Opción 2 — Node.js

Requiere [Node.js](https://nodejs.org/) instalado.

```bash
npx serve static
```

Abrir en el navegador: `http://localhost:3000`

---

## Estructura

```
static/
  index.html        # Landing page principal
  css/
    styles.css      # Estilos (paleta tierra suave)
  js/
    main.js         # Navegación y render de cards
    chatbot.js      # Chatbot de FAQs (cliente)
docker-compose.yml  # Servidor nginx
```

---

## Tecnologías

- HTML5 / CSS3 / JavaScript vanilla
- nginx:alpine (Docker)
