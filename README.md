# Proyecto de estudio: gRPC con Go

Este repositorio es únicamente para fines de estudio y experimentación con gRPC en Go.

- **Objetivo:** Aprender a definir servicios gRPC, generar código con Protocol Buffers y exponer un servidor gRPC en Go.
- **No para producción:** Código educativo — no optimizado ni seguro para entornos de producción.

Requisitos
- Go 1.26 o superior
- Protoc y plugins de Go si desea regenerar los archivos `.pb.go`

Uso rápido
1. Ejecutar migraciones automáticas y levantar el servidor:

```bash
go run ./cmd/server
```

2. El servidor escucha en `:50051` por defecto.

Notas
- La migración crea la tabla `categories` en `./db.db` si no existe.
- Los protobufs están en el directorio `proto/` y el código generado en `internal/delivery/grpc/pb`.

Contribuciones
- Este proyecto está pensado para uso personal y educativo; siéntase libre de experimentar y adaptar el código para aprender.

Hecho con ❤️ para aprender gRPC.