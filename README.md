**API de e-commerce desarrollada en Go usando:**

- Go (Golang)
= PostgreSQL
- Docker
- sqlc
- Goose (migraciones)
- Chi (router)

**Requisitos previos**

Antes de iniciar, asegúrate de tener instalado:
- Docker Desktop
- Golang (versión 1.21+ recomendada)
- TablePlus (opcional, para visualizar la base de datos)
- Postman (opcional, para probar los endpoints)

**Instalación del proyecto**
Clonar el repositorio

	git clone https://github.com/AMOZINGAS/go-api-ecommerce.git

Entrar al directorio:

	cd go-api-ecommerce

Instalar dependencias de Go

	go mod tidy

Este comando descargará todas las dependencias necesarias del proyecto.

**Configurar variables de entorno**

Editar el archivo .env y configurar:

- Usuario
- Contraseña
- Nombre de la base de datos
- Puerto
- Ruta de conexión (DSN)

Estos mismos datos deberán usarse en TablePlus si deseas visualizar la base de datos.

Levantar la base de datos con Docker

Asegúrate de que Docker Desktop esté corriendo.

Luego ejecuta:

	docker compose up -d

Esto iniciará el contenedor de PostgreSQL.

Ejecutar migraciones

En una nueva terminal:

	goose up

Esto creará las tablas en la base de datos.

Generar código con sqlc

	sqlc generate

Este comando genera el código Go basado en los archivos SQL.

Ejecutar la aplicación

	go run ./cmd/*.go

El servidor iniciará en:

http://localhost:8080

Endpoints disponibles

GET /health -> Verifica que el servidor esté activo
	
GET /products -> Lista los productos

POST /orders-> agrega una orden de productos

sintaxis
  
	{
		"customer_Id": 3,  
		"items": [
			{"product_Id": 1, "quantity": 2},
			{"product_Id": 2, "quantity": 1}
		]
	}
  
Probar la API

Puedes usar:
- Postman
- Insomnia
- Curl
- Navegador (para endpoints GET)

Notas importantes

Asegúrate de que Docker esté activo antes de correr la aplicación.

Si tienes problemas con dependencias, ejecuta:

	go mod tidy

Si modificas consultas SQL, vuelve a ejecutar:

	sqlc generate
