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
- ✅ **Tratamento de Erros** - Respostas HTTP padronizadas
- ✅ **Clean Architecture** - Separação clara de responsabilidades
- ✅ **Docker & Docker Compose** - Containerização para testes locais
- ✅ **Cloud Run Ready** - Deploy com um comando

---

## 🚀 Como Executar

### Pré-requisitos

- [Go 1.23+](https://golang.org/dl/)
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/)
- Conta no [WeatherAPI](https://www.weatherapi.com/) com API Key *(free tier disponível)*

### ✨ Execução com Docker Compose

🚀 Subir a aplicação
```bash
# 1. Configurar variáveis de ambiente
cp .env.example .env
# Edite .env e adicione sua WEATHER_API_KEY

# 2. Iniciar o container
docker-compose up --build -d

# 3. Verificar logs
docker-compose logs -f app
```

### ✅ Confirmação do Serviço

Quando tudo estiver funcionando, você verá:
```
Starting Weather App on port 8080
```

### 🔄 Comandos Úteis

```bash
# Status dos containers
docker-compose ps

# Parar serviços
docker-compose down

# Rebuild completo (limpar volumes)
docker-compose down -v
docker-compose up --build -d

# Ver logs em tempo real
docker-compose logs -f app
```

---

## 🧪 Como Testar

### 🌐 **REST API** - Porta 8080

#### ✅ Sucesso - CEP Válido
```http
GET http://localhost:8080/weather?cep=01310100
```

**Response (200 OK):**
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

#### ❌ Erro - CEP Inválido (formato)
```http
GET http://localhost:8080/weather?cep=123
```

**Response (422 Unprocessable Entity):**
```json
{
  "message": "invalid zipcode"
}
```

#### ❌ Erro - CEP Não Encontrado
```http
GET http://localhost:8080/weather?cep=99999999
```

**Response (404 Not Found):**
```json
{
  "message": "can not find zipcode"
}
```

#### 🏥 Health Check
```http
GET http://localhost:8080/
```

**Response (200 OK):**
```json
{
  "status": "ok"
}
```

> 📁 **Arquivos de teste:** Veja pasta `api/` para testes HTTP prontos

---

## 📊 Arquitetura

```
🐳 DOCKER ARCHITECTURE
┌──────────────────────────────────────┐
│         Container Services           │
│      Weather App :8080               │
│                                      │
│  ├─ ViaCEP API (https)               │
│  └─ WeatherAPI (https)               │
└──────────────────────────────────────┘
           ⬇️
┌──────────────────────────────────────┐
│        APPLICATION LAYERS            │
│  REST API (Chi Router)               │
├──────────────────────────────────────┤
│         Use Cases                    │
│  GetWeatherByZipCode                 │
├──────────────────────────────────────┤
│         Entities                     │
│  Weather, Location                   │
├──────────────────────────────────────┤
│       Infrastructure                 │
│  ViaCEP Service  │  WeatherAPI       │
└──────────────────────────────────────┘
```

---

## 📁 Estrutura do Projeto

```
weather-app/
├── cmd/
│   └── weatherapp/           # Aplicação principal
│       └── main.go           # Entry point
├── internal/
│   ├── entity/               # Entidades de domínio
│   │   └── weather.go
│   ├── usecase/              # Casos de uso
│   │   └── get_weather_by_zipcode.go
│   └── infra/                # Infraestrutura
│       ├── service/          # Serviços externos
│       │   ├── viacep.go     # Integração ViaCEP
│       │   └── weather.go    # Integração WeatherAPI
│       └── web/              # HTTP
│           ├── handler.go    # Handlers
│           └── server.go     # Router
├── api/                      # Testes HTTP
│   ├── health_check.http
│   ├── get_weather_success.http
│   ├── get_weather_invalid.http
│   └── get_weather_notfound.http
├── tests/                    # Testes automatizados
├── Dockerfile                # Build para Docker
├── docker-compose.yml        # Orquestração
├── .env.example              # Variáveis de ambiente
├── .gitignore
├── go.mod                    # Dependências
└── README.md
```

---

## 🛠️ Tecnologias Utilizadas

### Core
- **Go 1.23** - Linguagem principal
- **Chi Router** - HTTP router minimalista e poderoso
- **Standard Library** - Sem dependências pesadas

### APIs Externas
- **ViaCEP** - Busca de localização por CEP
- **WeatherAPI** - Dados meteorológicos em tempo real

### DevOps
- **Docker** - Containerização
- **Docker Compose** - Orquestração local

---

## 📐 Fórmulas de Conversão

As conversões de temperatura utilizadas no projeto:

$$\text{Fahrenheit} = \text{Celsius} \times 1.8 + 32$$

$$\text{Kelvin} = \text{Celsius} + 273$$

---

## ☁️ Deploy no Google Cloud Run

### 📝 Pré-requisitos

1. Conta no [Google Cloud](https://console.cloud.google.com/)
2. Projeto criado no GCP
3. [gcloud CLI](https://cloud.google.com/sdk) instalado e autenticado

### 🚀 Deploy Automático

```bash
# 1. Fazer login no GCP
gcloud auth login

# 2. Definir projeto
gcloud config set project YOUR_PROJECT_ID

# 3. Fazer build e push da imagem
gcloud run deploy weather-app \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=YOUR_API_KEY
```

### 📍 Acessar a Aplicação

Após deploy bem-sucedido, você receberá uma URL como:
```
https://weather-app-xxxxx-uc.a.run.app
```

Teste a aplicação:
```bash
# Health check
curl https://weather-app-xxxxx-uc.a.run.app/

# Buscar clima
curl "https://weather-app-xxxxx-uc.a.run.app/weather?cep=01310100"
```

---

## 🧪 Testes

### Executar Testes Locais

```bash
# Build a imagem
docker-compose build

# Rodar testes HTTP (usando VS Code REST Client ou Postman)
# Abra os arquivos em api/ e envie as requisições
```

### Teste com curl

```bash
# Health check
curl http://localhost:8080/

# CEP válido (sucesso)
curl "http://localhost:8080/weather?cep=01310100"

# CEP inválido (error 422)
curl "http://localhost:8080/weather?cep=123"

# CEP não encontrado (error 404)  
curl "http://localhost:8080/weather?cep=99999999"
```

---

## 🐛 Troubleshooting

### API Key Inválida
```
❌ Error: "error: unable to fetch weather data"
```
**Solução:** Verifique se `WEATHER_API_KEY` está definida corretamente em `.env`

### Container não inicia
```bash
# Verificar logs
docker-compose logs app

# Rebuild limpo
docker-compose down -v
docker-compose up --build -d
```

### Porta já em uso
```bash
# Mudar porta em docker-compose.yml
# Altere: "8080:8080" para "8081:8080"
docker-compose up --build -d
```

---

## 👤 Autor

- **GitHub:** [sarapmc](https://github.com/sarapmc)
- **Email (GitHub):** sarapmc@hotmail.com  
- **Email (Google Cloud):** sarapmcantao@gmail.com

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## 📚 Referências

- [Go Documentation](https://golang.org/doc/)
- [ViaCEP API](https://viacep.com.br/)
- [WeatherAPI Documentation](https://www.weatherapi.com/docs/)
- [Google Cloud Run](https://cloud.google.com/run/docs)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Chi Router](https://github.com/go-chi/chi)