# 🚀 Quick Start Guide

## 1️⃣ Preparação Rápida (2 min)

```bash
# Clone o repositório
git clone <seu-repo>
cd desafio-deploy-google-cloud-run

# Copie o arquivo de ambiente
cp .env.example .env

# Edite .env e adicione sua WEATHER_API_KEY
# (obtenha em https://www.weatherapi.com/)
```

## 2️⃣ Executar Localmente (Docker)

```bash
# Inicie os containers
make docker-up

# Teste a aplicação
curl "http://localhost:8080/weather?cep=01310100"

# Ver logs
make docker-logs

# Parar
make docker-down
```

## 3️⃣ Executar Tests

```bash
# Rodar todos os testes
make test

# Com coverage
make test-verbose

# Testes HTTP via VS Code REST Client
# Abra os arquivos em api/ e use "Send Request"
```

## 4️⃣ Deploy no Google Cloud Run

```bash
# Fazer login
gcloud auth login

# Deploy
gcloud run deploy weather-app \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=sua_api_key_aqui
```

## 📋 Comandos Disponíveis

```bash
make              # Ver todos os comandos
make build        # Build local
make run          # Executar localmente
make docker-up    # Docker Compose up
make docker-down  # Docker Compose down
make docker-logs  # Ver logs
make test         # Rodar testes
make clean        # Limpar tudo
make deploy       # Deploy GCP
```

## 🔗 Endpoints

| Método | URL | Descrição |
|--------|-----|----------|
| `GET` | `/` | Health check |
| `GET` | `/weather?cep=XXXXXXXX` | Buscar clima por CEP |

## 💡 Exemplos

### ✅ Sucesso
```bash
curl "http://localhost:8080/weather?cep=01310100"
# { "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5 }
```

### ❌ CEP Inválido
```bash
curl "http://localhost:8080/weather?cep=123"
# HTTP 422 - invalid zipcode
```

### ❌ CEP Não Encontrado
```bash
curl "http://localhost:8080/weather?cep=99999999"
# HTTP 404 - can not find zipcode
```

---

📚 Para mais informações, veja:
- [README.md](README.md) - Documentação completa
- [DEPLOYMENT.md](DEPLOYMENT.md) - Guia detalhado de deploy
