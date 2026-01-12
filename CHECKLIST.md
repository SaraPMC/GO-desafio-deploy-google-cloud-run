# ✅ Checklist - Próximos Passos

## 🚀 Antes de Começar

- [ ] **Clonar/Acessar o repositório**
  ```bash
  cd d:\Estudos\Go\_Pos\desafio-deploy-google-cloud-run
  ```

- [ ] **Criar conta em WeatherAPI** (se não tiver)
  - Acessar: https://www.weatherapi.com/
  - Registrar e obter API Key (grátis)

- [ ] **Configurar variáveis de ambiente**
  ```bash
  cp .env.example .env
  # Editar .env e colocar sua WEATHER_API_KEY
  ```

---

## 💻 Executar Localmente

- [ ] **Testar com Docker Compose**
  ```bash
  make docker-up
  ```

- [ ] **Verificar se está rodando**
  ```bash
  curl http://localhost:8080/
  # Deve retornar: {"status":"ok"}
  ```

- [ ] **Testar endpoints**
  - Sucesso: `curl "http://localhost:8080/weather?cep=01310100"`
  - Erro 422: `curl "http://localhost:8080/weather?cep=123"`
  - Erro 404: `curl "http://localhost:8080/weather?cep=99999999"`

- [ ] **Rodar testes**
  ```bash
  make test
  ```

- [ ] **Ver logs**
  ```bash
  make docker-logs
  ```

- [ ] **Parar containers**
  ```bash
  make docker-down
  ```

---

## 🌐 Testar via REST Client (VS Code)

- [ ] **Instalar extensão**
  - Extensão: "REST Client" (humao.rest-client)

- [ ] **Testar requisições**
  - Abrir arquivos em `api/` pasta
  - Clicar em "Send Request" em cada um

---

## ☁️ Preparar para Deploy (Google Cloud Run)

- [ ] **Criar conta Google Cloud** (se não tiver)
  - https://console.cloud.google.com/
  - Criar novo projeto

- [ ] **Instalar gcloud CLI**
  ```bash
  # Já está instalado na sua máquina, apenas:
  gcloud auth login
  ```

- [ ] **Configurar projeto padrão**
  ```bash
  gcloud config set project YOUR_PROJECT_ID
  ```

- [ ] **Verificar configuração**
  ```bash
  gcloud config list
  ```

---

## 🚀 Fazer Deploy

- [ ] **Deploy inicial**
  ```bash
  gcloud run deploy weather-app \
    --source . \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated \
    --set-env-vars WEATHER_API_KEY=sua_api_key_aqui
  ```

- [ ] **Anotar URL fornecida pelo GCP**
  ```
  Service deployed to: https://weather-app-xxxxx-uc.a.run.app
  ```

- [ ] **Testar aplicação no Cloud Run**
  ```bash
  curl https://weather-app-xxxxx-uc.a.run.app/
  curl "https://weather-app-xxxxx-uc.a.run.app/weather?cep=01310100"
  ```

- [ ] **Verificar logs no Cloud Run**
  ```bash
  gcloud run services logs read weather-app --region us-central1 --limit 50
  ```

---

## 📚 Documentação

Leia os arquivos nesta ordem:

1. ✅ **QUICKSTART.md** - Começar rápido
2. 📖 **README.md** - Documentação completa
3. 🚀 **DEPLOYMENT.md** - Guia de deploy detalhado
4. 🧪 **TESTING_GUIDE.md** - Guia de testes
5. 📦 **PROJECT_STRUCTURE.md** - Visão geral da estrutura

---

## 🎯 Checklist Final

- [ ] Aplicação roda localmente ✓
- [ ] Testes passam ✓
- [ ] Docker Compose funciona ✓
- [ ] Deploy no Cloud Run bem-sucedido ✓
- [ ] URL ativa e respondendo ✓
- [ ] Código commitado no GitHub ✓
- [ ] Documentação completa ✓

---

## 🆘 Problemas Comuns

### Docker não encontrado
```bash
# Instalar Docker: https://www.docker.com/
```

### gcloud command not found
```bash
# Feche e reabra o VS Code
# Ou execute:
gcloud --version
```

### API Key inválida
```bash
# Verify key em: https://www.weatherapi.com/
# Atualizar em .env e redeploy
```

### Porta 8080 em uso
```bash
# Mudar em docker-compose.yml:
# "8080:8080" → "8081:8080"
```

---

## 📞 Contato

- GitHub: sarapmc@hotmail.com
- Google Cloud: sarapmcantao@gmail.com

---

## 🎉 Quando Tudo Funcionar

1. ✅ Aplicação local rodando
2. ✅ Testes passando
3. ✅ Deploy no GCP bem-sucedido
4. ✅ URL ativa respondendo corretamente

**Parabéns! 🎊 Você completou o desafio!**
