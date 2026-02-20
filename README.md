# fullstack app

This is a fullstack payment dashboard application built using:

- Backend: Golang (clean architecture style — handler / usecase / repository)
- Frontend: Nuxt (dashboard UI)
- API contract: OpenAPI (generated with oapi-codegen)
- Database: SQLite (auto-created on first run)

The backend provides authentication and payment dashboard APIs.  
The frontend consumes the API to display and manage payment data.

---

## list of tools version of your machine:

```bash
go version go1.25.5 darwin/arm64
node v24.13.1