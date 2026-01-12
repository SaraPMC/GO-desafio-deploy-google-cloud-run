# 🧪 CEPs para Testes

Use os CEPs abaixo para testar a aplicação:

## ✅ CEPs Válidos (Cidades Principais)

| CEP | Cidade | Estado | Uso |
|-----|--------|--------|-----|
| `01310100` | São Paulo | SP | Av. Paulista |
| `20040020` | Rio de Janeiro | RJ | Centro |
| `30130100` | Belo Horizonte | MG | Centro |
| `50010000` | Recife | PE | Centro |
| `60060000` | Fortaleza | CE | Centro |
| `70040010` | Brasília | DF | Esplanada dos Ministérios |
| `80010000` | Curitiba | PR | Centro |
| `90010000` | Porto Alegre | RS | Centro |

## ❌ CEPs Inválidos (Para Testes de Erro)

| Entrada | Motivo | Status HTTP | Mensagem |
|---------|--------|-------------|----------|
| `123` | Menos de 8 dígitos | 422 | `invalid zipcode` |
| `ABCDEFGH` | Contém letras | 422 | `invalid zipcode` |
| `123456789` | Mais de 8 dígitos | 422 | `invalid zipcode` |
| `00000000` | CEP não existe | 404 | `can not find zipcode` |
| `99999999` | CEP fictício | 404 | `can not find zipcode` |
| (vazio) | Sem parâmetro | 422 | `invalid zipcode` |

## 🧪 Exemplos de Teste

### Test 1: Health Check
```bash
curl http://localhost:8080/
```

### Test 2: CEP Válido
```bash
curl "http://localhost:8080/weather?cep=01310100"
# Response: { "temp_C": 25.5, "temp_F": 77.9, "temp_K": 298.5 }
```

### Test 3: CEP Inválido
```bash
curl "http://localhost:8080/weather?cep=123"
# Response: HTTP 422 - { "message": "invalid zipcode" }
```

### Test 4: CEP Não Encontrado
```bash
curl "http://localhost:8080/weather?cep=99999999"
# Response: HTTP 404 - { "message": "can not find zipcode" }
```

### Test 5: Sem CEP
```bash
curl "http://localhost:8080/weather"
# Response: HTTP 422 - { "message": "invalid zipcode" }
```

## 🔧 Teste com VS Code REST Client

Todos os testes estão prontos em `api/`:

1. Instale extensão **REST Client** (humao.rest-client)
2. Abra qualquer arquivo em `api/`
3. Clique em **"Send Request"**

---

## 📍 Gerador de CEP

Precisa de outros CEPs? Use:
- https://viacep.com.br/ (Buscar por endereço)
- https://www.4devs.com.br/gerador_cep (Gerar CEPs aleatórios)

## 📚 Mais Informações

Para detalhes sobre formato e validação, veja:
- [README.md](README.md) - Documentação principal
- [QUICKSTART.md](QUICKSTART.md) - Guia rápido
