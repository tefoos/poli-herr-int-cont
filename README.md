# int-cont

API REST en Go para manipular estructuras de datos (pila, cola y lista enlazada) con Redis como capa de persistencia, orquestados mediante Docker Compose.

## Requisitos

- Docker
- Docker Compose

## Ejecutar el proyecto

```bash
docker compose up --build
```

## Limpiar la persistencia de Redis

```bash
docker compose exec redis redis-cli FLUSHALL
```

## Ejecutar los tests

```bash
chmod +x test.sh
./test.sh
```

## Endpoints

### Stack
- POST /stack/{name}/push
- POST /stack/{name}/pop
- GET /stack/{name}

### Queue
- POST /queue/{name}/enqueue
- POST /queue/{name}/dequeue
- GET /queue/{name}

### Linked List
- POST /list/{name}/append
- POST /list/{name}/remove
- GET /list/{name}
test webhook
test webhook
test webhook
