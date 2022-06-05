# TODO application using GraphQL and gRPC by Go

## ER diagram

```mermaid
erDiagram

users {
  id number PK
}

tasks {
  id number PK
  user_id number FK
  name string
}

tags {
  id number PK
  name string
}

tasks_tags {
  id number PK
  task_id number FK
  tag_id number FK
}

users ||--o{ tasks : "1:n"
tasks ||--o{ tasks_tags : "1:n"
tags ||--|{ tasks_tags : "1:n"
```

## Links

- [gqlgen](https://gqlgen.com/)
- [ent](https://entgo.io/)
- [Atlas](https://atlasgo.io/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [gRPC](https://grpc.io/)