# SIW_EE_C1 — Media Hub

Proyecto para Sistemas de Interfaz Web — Ernesto Eisenkolbl.

**URL pública:** [media.bashkoko.com](https://media.bashkoko.com)

---

## Descripción

**Media Hub** es un dashboard web + chatbot de IA para monitorear y controlar un pipeline de medios (películas, series, descargas). Permite al usuario ver qué hay en la biblioteca, solicitar nuevos contenidos, revisar el estado de descargas y conversar con un asistente inteligente que tiene contexto en tiempo real del sistema.

---

## Arquitectura

```
┌──────────────────────────────────────────┐
│          media.bashkoko.com              │
│         (Authelia + nginx proxy)         │
└──────────────────┬───────────────────────┘
                   │ HTTPS
                   ▼
┌──────────────────────────────────────────┐
│         media-hub (binario Go)           │
│                                          │
│  ┌─────────────┐   ┌──────────────────┐  │
│  │ React SPA   │   │   API REST       │  │
│  │ (embebida)  │   │  /api/*  (chi)   │  │
│  └─────────────┘   └──────┬───────────┘  │
│                           │              │
│  ┌────────────────────────┼───────────┐  │
│  │         Clientes arr               │  │
│  │  Sonarr · Radarr · Jellyseerr      │  │
│  │  Jellyfin · qBittorrent            │  │
│  └────────────────────────────────────┘  │
│                                          │
│  ┌──────────┐   ┌──────────────────────┐ │
│  │ SQLite   │   │  Ollama (LLM local)  │ │
│  │ (sesiones│   │  chat con contexto   │ │
│  │  y chat) │   │  del pipeline        │ │
│  └──────────┘   └──────────────────────┘ │
└──────────────────────────────────────────┘
```

### Stack tecnológico

| Capa | Tecnología |
|---|---|
| Frontend | React 18, Vite, Tailwind CSS |
| Backend | Go, chi router, `embed.FS` |
| Base de datos | SQLite (sesiones + historial de chat) |
| IA / Chatbot | Ollama (LLM local) |
| Autenticación | Authelia (SSO por headers de proxy) |
| Reverse proxy | nginx |

### Cómo funciona

1. **Build:** `make build` compila el binario Go e incrusta el frontend React compilado (`frontend/dist/`) dentro del binario usando `embed.FS`. Se distribuye como **un solo ejecutable**.
2. **Autenticación:** El nginx upstream valida la sesión con Authelia. Los headers `Remote-User`, `Remote-Groups` y `Remote-Name` llegan al backend, que los convierte en una sesión SQLite.
3. **Dashboard:** El frontend llama a `/api/dashboard/*`. El backend consulta Sonarr, Radarr, Jellyseerr, Jellyfin y qBittorrent en paralelo y agrega los resultados.
4. **Chat:** El usuario escribe en el chatbot → `/api/chat` → el backend inyecta contexto en vivo del pipeline (cola de descargas, biblioteca, requests pendientes) → Ollama genera la respuesta → streaming de vuelta al navegador.
5. **Admin:** Usuarios con grupo `admins` pueden ver logs de sesiones y el historial completo de conversaciones de todos los usuarios.

---

## Rutas API principales

| Método | Ruta | Descripción |
|---|---|---|
| GET | `/api/health` | Estado del servicio |
| GET | `/api/me` | Usuario autenticado actual |
| GET | `/api/dashboard/library` | Biblioteca Jellyfin |
| GET | `/api/dashboard/requests` | Solicitudes pendientes (Jellyseerr) |
| GET | `/api/dashboard/queue` | Cola de descargas |
| GET | `/api/dashboard/pipeline` | Estado del pipeline completo |
| POST | `/api/dashboard/requests` | Crear nueva solicitud |
| POST | `/api/chat` | Enviar mensaje al chatbot |
| DELETE | `/api/chat/session` | Limpiar sesión de chat |
| GET | `/api/admin/sessions` | (admin) Estadísticas de usuarios |
| GET | `/api/admin/messages` | (admin) Historial de chat |

---

## Servicios integrados

| Servicio | Función |
|---|---|
| **Sonarr** | Gestión de series de TV |
| **Radarr** | Gestión de películas |
| **Jellyseerr** | Solicitudes de contenido de los usuarios |
| **Jellyfin** | Servidor de streaming / biblioteca |
| **qBittorrent** | Cliente de descargas |
| **Ollama** | LLM local para el chatbot |

---

## Paneles de la UI

- **Library:** Explora la biblioteca, ve qué hay disponible, busca y solicita nuevas películas/series.
- **Chat:** Chatbot con IA que conoce el estado actual del pipeline (descargas activas, requests, contenido reciente).
- **Admin** *(solo admins):* Monitor de sesiones activas, historial de conversaciones de todos los usuarios.
