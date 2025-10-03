# Global-Scale-Auth-and-Policy-Service

JWT-secured REST API in Go with **Keycloak** for authentication and **OPA/Rego** for authorization. Containerized with Docker and shipped via GitHub Actions → GHCR.

---

## Features
- **AuthN:** Validates `Bearer` tokens (JWT) against Keycloak **JWKS**.
- **AuthZ:** **Policy-as-code** with **OPA/Rego** (deny-by-default).
- **Clean middleware chain:** JWT (401 on fail) → OPA (403 on deny) → handler.
- **Containerized:** multi-stage Dockerfile.
- **CI/CD:** builds, tests, and pushes image to **GHCR**.

---

## Architecture
Client ──HTTP──> Gin Router

├── JWT Middleware (validate token via JWKS) ── 401 on fail

├── OPA Middleware (evaluate Rego policy) ── 403 on deny

└── /hello Handler (returns JSON on success)

