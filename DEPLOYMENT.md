# 🚀 Guia de Deployment - Google Cloud Run

## Pré-requisitos

### 1. Instalar Google Cloud SDK
- [Download gcloud CLI](https://cloud.google.com/sdk)
- Verificar instalação: `gcloud --version`

### 2. Criar Conta e Projeto no GCP
- Acessar [Google Cloud Console](https://console.cloud.google.com/)
- Criar um novo projeto
- Habilitar Cloud Run API

### 3. WeatherAPI Key
- Registrar em [WeatherAPI.com](https://www.weatherapi.com/)
- Obter API Key do dashboard

## Passo a Passo para Deploy

### 1️⃣ Autenticação

```bash
# Login na sua conta Google
gcloud auth login

# Definir projeto padrão
gcloud config set project YOUR_PROJECT_ID

# Verificar configuração
gcloud config list
```

### 2️⃣ Preparar Variáveis de Ambiente

```bash
# Criar arquivo .env com sua API Key
echo "WEATHER_API_KEY=sua_api_key_aqui" > .env
```

### 3️⃣ Deploy Automático

```bash
# Deploy com source (recomendado para primeiro deploy)
gcloud run deploy weather-app \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=sua_api_key_aqui \
  --memory 256Mi \
  --cpu 1 \
  --timeout 300s
```

### 4️⃣ Testar a Aplicação

Após deploy bem-sucedido, você receberá uma URL:

```bash
# Health check
curl https://weather-app-xxxxx-uc.a.run.app/

# Teste com CEP válido
curl "https://weather-app-xxxxx-uc.a.run.app/weather?cep=01310100"

# Teste com CEP inválido
curl "https://weather-app-xxxxx-uc.a.run.app/weather?cep=123"
```

---

## 📋 Verificar Status do Deployment

```bash
# Listar todas as Cloud Run services
gcloud run services list

# Ver detalhes de um serviço
gcloud run services describe weather-app --region us-central1

# Ver logs em tempo real
gcloud run services logs read weather-app --region us-central1 --limit 50 --follow
```

---

## 🔄 Redeploy (Atualizações)

Após fazer mudanças no código:

```bash
gcloud run deploy weather-app \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=sua_api_key_aqui
```

---

## 💰 Custos (Free Tier)

✅ **Google Cloud Run Free Tier inclui:**
- 2 milhões de requisições/mês
- 360.000 GB-segundos/mês
- 180.000 vCPU-segundos/mês

> ℹ️ Esse projeto dificilmente ultrapassará os limites do free tier

---

## 🐛 Troubleshooting

### Erro: "Cloud Run API not enabled"
```bash
gcloud services enable run.googleapis.com
```

### Erro: "Permission denied"
```bash
# Verifique as permissões da sua conta
gcloud projects get-iam-policy YOUR_PROJECT_ID
```

### Erro: "Build failed"
```bash
# Verificar logs do build
gcloud builds log --limit=100
```

### Erro: 500 - Internal Server Error
```bash
# Verificar se WEATHER_API_KEY está configurada
gcloud run services describe weather-app --region us-central1 --format="value(spec.template.spec.containers[0].env)"
```

---

## 🔐 Segurança

### Melhorias Recomendadas

1. **Secret Manager** - Armazenar API Keys com segurança
```bash
# Criar secret
echo -n "sua_api_key_aqui" | gcloud secrets create weather-api-key --data-file=-

# Usar em deployment
gcloud run deploy weather-app \
  --source . \
  --platform managed \
  --region us-central1 \
  --set-env-vars WEATHER_API_KEY=projects/PROJECT_ID/secrets/weather-api-key/versions/latest
```

2. **Authentication** - Restringir acesso
```bash
# Remover acesso público (alterar ao deploy)
gcloud run services add-iam-policy-binding weather-app \
  --member=serviceAccount:your-account@appspot.gserviceaccount.com \
  --role=roles/run.invoker \
  --region us-central1
```

---

## 📊 Monitoramento

### Cloud Monitoring

```bash
# Ver métricas no console
# https://console.cloud.google.com/monitoring

# Criar alerta para altos erros
gcloud alpha monitoring policies create \
  --notification-channels=CHANNEL_ID \
  --display-name="Weather App - High Error Rate"
```

---

## 🔗 Referências Úteis

- [Cloud Run Documentation](https://cloud.google.com/run/docs)
- [gcloud run deploy](https://cloud.google.com/sdk/gcloud/reference/run/deploy)
- [Cloud Run Best Practices](https://cloud.google.com/run/docs/quickstarts/build-and-deploy)
- [Free Tier Pricing](https://cloud.google.com/free)

---

## 📞 Suporte

- Email de GitHub: sarapmc@hotmail.com
- Email de Google Cloud: sarapmcantao@gmail.com
