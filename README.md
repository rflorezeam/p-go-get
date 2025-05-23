# Microservicio de Lectura de Libros

## Desarrollador
Ricardo Florez

## Descripción
Este microservicio es responsable de listar todos los libros disponibles en el sistema. Forma parte de una arquitectura de microservicios para la gestión de una biblioteca digital.

## Características
- Implementado en Go 1.21
- Arquitectura limpia (Clean Architecture)
- Endpoints RESTful
- Integración con MongoDB
- Despliegue en Kubernetes

## Estructura del Proyecto
```
.
├── config/         # Configuración de la base de datos
├── models/         # Modelos de datos
├── repositories/   # Capa de acceso a datos
├── services/      # Lógica de negocio
├── k8s/           # Configuración de Kubernetes
└── tests/         # Pruebas unitarias
```

## API Endpoint
- **GET** `/libros`
  - Puerto: 30082 (NodePort)
  - Retorna la lista completa de libros

### Ejemplo de Respuesta
```json
[
    {
        "id": "5f7b5e1b9d3e2a1b4c7d8e9f",
        "titulo": "El Quijote",
        "autor": "Miguel de Cervantes"
    },
    {
        "id": "5f7b5e1b9d3e2a1b4c7d8e9g",
        "titulo": "Cien años de soledad",
        "autor": "Gabriel García Márquez"
    }
]
```

## Configuración Kubernetes
- Deployment con 3 réplicas
- Service tipo NodePort (30082)
- Conexión a MongoDB mediante Service Discovery

## Variables de Entorno
- MONGODB_URI: URI de conexión a MongoDB

## Despliegue
```bash
# Construir la imagen
docker build -t libro-read:latest .

# Desplegar en Kubernetes
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

## Pruebas
```bash
# Ejecutar pruebas unitarias
go test ./...

# Probar el endpoint
curl http://localhost:30082/libros
```

## Monitoreo
El servicio puede ser monitoreado mediante:
- Logs de Kubernetes
- Métricas de contenedor
- Estado del Service y Deployment

## Rendimiento
- Implementa paginación para grandes conjuntos de datos
- Caché de resultados frecuentes
- Optimización de consultas MongoDB 