# 🌦️ Weather App - GO Desafio Deploy Google Cloud Run

[![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-Required-blue.svg)](https://docker.com)
[![WeatherAPI](https://img.shields.io/badge/WeatherAPI-Integration-orange.svg)](https://www.weatherapi.com/)
[![ViaCEP](https://img.shields.io/badge/ViaCEP-Integration-green.svg)](https://viacep.com.br/)
[![Cloud Run](https://img.shields.io/badge/Platform-Google%20Cloud%20Run-red.svg)](https://cloud.google.com/run)
[![Clean Architecture](https://img.shields.io/badge/Architecture-Clean-green.svg)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## 📋 Sobre o Projeto

Este projeto é resultado de um **desafio prático** de desenvolvimento em Go e deploy em Google Cloud Run. O sistema recebe um CEP (código de endereçamento postal brasileiro), identifica a cidade correspondente e retorna informações do clima atual em diferentes unidades de temperatura.

### 🎯 Objetivo do Desafio

Desenvolver um microsserviço em Go que:
- ✅ Receba um CEP válido (8 dígitos)
- ✅ Consulte a localização via API ViaCEP
- ✅ Obtenha dados meteorológicos via WeatherAPI
- ✅ Retorne temperatura em **Celsius, Fahrenheit e Kelvin**
- ✅ Responda com códigos HTTP apropriados
- ✅ Deploy automático no **Google Cloud Run** (free tier)

### 🏆 Funcionalidades Implementadas

- ✅ **Busca de CEP** - Integração com ViaCEP para validar e localizar endereços
- ✅ **Consulta Meteorológica** - Integração com WeatherAPI para dados climáticos
- ✅ **Conversão de Temperaturas** - Celsius → Fahrenheit, Kelvin
- ✅ **Validação de Entrada** - Formato correto de CEP (8 dígitos)
- ✅ **Tratamento de Erros** - Respostas HTTP padronizadas (200, 422, 404, 500)
- ✅ **Clean Architecture** - Separação clara de responsabilidades
- ✅ **Docker & Docker Compose** - Containerização para testes locais
- ✅ **Testes Automatizados** - Unit tests inclusos
- ✅ **Cloud Run Ready** - Deploy com um comando

---

## 🚀 Quick Start (5 minutos)

### Pré-requisitos

- [Go 1.23+](https://golang.org/dl/)
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/)
- Conta no [WeatherAPI](https://www.weatherapi.com/) (free tier)

### 1️⃣ Configurar Variáveis de Ambiente

```bash
cp .env.example .env
# Edite .env e adicione sua WEATHER_API_KEY de https://www.weatherapi.com/
```

### 2️⃣ Executar Localmente

```bash
# Subir com Docker Compose
docker-compose up --build -d

# Ver logs
docker-compose logs -f app

# Testar
curl http://localhost:8080/
curl "http://localhost:8080/weather?cep=01310100"

# Parar
docker-compose down
```

---

## 🧪 Endpoints & Testes

### 🌐 REST API

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/` | Health check |
| `GET` | `/weather?cep=XXXXXXXX` | Buscar clima por CEP |

### ✅ Exemplo de Sucesso

```bash
curl "http://localhost:8080/weather?cep=01310100"
```

**Response (HTTP 200):**
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

### ❌ Exemplo de Erro - CEP Inválido

```bash
curl "http://localhost:8080/weather?cep=123"
```

**Response (HTTP 422):**
```json
{
  "message": "invalid zipcode"
}
```

### ❌ Exemplo de Erro - CEP Não Encontrado

```bash
curl "http://localhost:8080/weather?cep=99999999"
```

**Response (HTTP 404):**
```json
{
  "message": "can not find zipcode"
}
```

### 🧪 Testes com arquivos HTTP

A pasta `api/` contém arquivos `.http` para testar:
- `health_check.http` - Health check
- `get_weather_success.http` - CEP válido
- `get_weather_invalid.http` - CEP inválido
- `get_weather_notfound.http` - CEP não existe

Use com VS Code REST Client (extensão: humao.rest-client)

---

## 📁 Estrutura do Projeto

```
.
├── cmd/weatherapp/              # Aplicação principal
│   └── main.go                  # Entry point
├── internal/
│   ├── entity/                  # Entidades de domínio
│   │   └── weather.go
│   ├── usecase/                 # Casos de uso (business logic)
│   │   └── get_weather_by_zipcode.go
│   └── infra/                   # Infraestrutura
│       ├── service/             # Serviços externos
│       │   ├── viacep.go        # API ViaCEP
│       │   ├── viacep_test.go   # Testes ViaCEP
│       │   └── weather.go       # API WeatherAPI
│       └── web/                 # HTTP
│           ├── handler.go       # HTTP handlers
│           ├── handler_test.go  # Testes handlers
│           └── server.go        # Router e server
├── api/                         # Testes HTTP (.http files)
├── Dockerfile                   # Build para Docker
├── docker-compose.yml           # Orquestração
├── go.mod                       # Dependências Go
├── go.sum                       # Lock file
├── .env.example                 # Template de variáveis
├── .gitignore
├── LICENSE                      # MIT
└── README.md                    # Este arquivo
```

---

## 🔧 Tecnologias Utilizadas

| Categoria | Tecnologia |
|-----------|-----------|
| **Linguagem** | Go 1.23 |
| **HTTP Router** | Chi v5 |
| **APIs Externas** | ViaCEP, WeatherAPI |
| **Containerização** | Docker, Docker Compose |
| **Cloud** | Google Cloud Run |
| **Arquitetura** | Clean Architecture |
| **Testes** | Go testing package |

---

## 📐 Fórmulas de Conversão

$$\text{Fahrenheit} = \text{Celsius} \times 1.8 + 32$$

$$\text{Kelvin} = \text{Celsius} + 273$$

---

## ☁️ Deploy no Google Cloud Run

### Pré-requisitos

1. Conta em [Google Cloud Console](https://console.cloud.google.com/)
2. Projeto criado no GCP
3. [gcloud CLI](https://cloud.google.com/sdk) instalado

### Passo-a-Passo

#### 1. Autenticar

```bash
gcloud auth login
gcloud config set project YOUR_PROJECT_ID
```

#### 2. Deploy

```bash
gcloud run deploy weather-app \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=your_api_key_here
```

#### 3. Testar

Após sucesso, você receberá uma URL como:
```
https://weather-app-xxxxx-uc.a.run.app
```

Teste:
```bash
curl https://weather-app-xxxxx-uc.a.run.app/
curl "https://weather-app-xxxxx-uc.a.run.app/weather?cep=01310100"
```

#### 4. Ver Logs

```bash
gcloud run services logs read weather-app --region us-central1 --limit 50
```

---

## 🧪 Testes Automatizados

### Executar Testes

```bash
# Todos os testes
go test ./...

# Com coverage
go test -cover ./...

# Verbose
go test -v ./...
```

### Testes Implementados

- ✅ Validação de CEP (formato correto)
- ✅ Busca de localização (ViaCEP)
- ✅ Health check endpoint
- ✅ Weather endpoint
- ✅ Tratamento de erros

---

## 🔄 Comandos Úteis

```bash
# Build da imagem
docker-compose build

# Iniciar containers
docker-compose up -d

# Ver status
docker-compose ps

# Ver logs
docker-compose logs -f app

# Parar containers
docker-compose down

# Remover volumes (reset completo)
docker-compose down -v
```

---

## 🐛 Troubleshooting

### API Key não funciona

```
Error: "error: unable to fetch weather data"
```

**Solução:**
1. Verifique se a chave está correta em `.env`
2. Testei com: `8c841b5ab4804b9d9ea14925261201`
3. Obtenha nova em: https://www.weatherapi.com/

### Container não inicia

```bash
# Ver erro
docker-compose logs app

# Rebuild limpo
docker-compose down -v
docker-compose up --build -d
```

### Porta 8080 em uso

Edite `docker-compose.yml`:
```yaml
ports:
  - "8081:8080"  # Mudou para 8081
```

### Module not found

```bash
# Limpar cache e baixar dependências novamente
go clean -modcache
go mod tidy
```

---

## 📊 Arquitetura

```
Entrada HTTP
    ↓
handler.GetWeatherByZipCode (web/handler.go)
    ↓
usecase.Execute (usecase/get_weather_by_zipcode.go)
    ↓
service.GetLocation (infra/service/viacep.go)
    ↓ [ViaCEP API]
    ↓
service.GetWeather (infra/service/weather.go)
    ↓ [WeatherAPI]
    ↓
Resposta JSON (entity.Weather)
    ↓
Cliente HTTP
```

---

## 👤 Autor

- **GitHub:** [sarapmc](https://github.com/SaraPMC)
- **Email (GitHub):** sarapmc@hotmail.com
- **Email (Google Cloud):** sarapmcantao@gmail.com
- **Repositório:** https://github.com/SaraPMC/GO-desafio-deploy-google-cloud-run

---

## 📄 Licença

Este projeto está sob a licença **MIT**. Veja [LICENSE](LICENSE) para detalhes.

---

## 📚 Referências

- [Go Documentation](https://golang.org/doc/)
- [ViaCEP API](https://viacep.com.br/)
- [WeatherAPI Documentation](https://www.weatherapi.com/docs/)
- [Google Cloud Run](https://cloud.google.com/run/docs)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Chi Router](https://github.com/go-chi/chi)