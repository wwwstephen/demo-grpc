```grpcurl -plaintext localhost:8089 list

Invoicer
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
```

```grpcurl -plaintext -d '{
  "amount": {
    "amount": 1000,
    "currency": "USD"
  },
  "from": "Alice",
  "to": "Bob"
}' localhost:8089 Invoicer/Create
{
  "pdf": "dGVzdA==",
  "docx": "dGVzdA=="
}
```
