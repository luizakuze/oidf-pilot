# Instância OpenID Federation (Lighthouse TA + IA + RP)

Este repositório contém um composição OpenID Federation baseado em [**Lighthouse**](https://go-oidfed.github.io/lighthouse/) e [**whoami-rp**](https://github.com/go-oidfed/whoami-rp).

## Checklist  

| Item | Status |
|------|--------|
| Trust Anchor (TA) | ✅ |
| Intermediate Authority (IA) | ✅ |
| Relying Party (RP) | ✅ |
| OpenID Provider (OP) | ❌    |
| Metadata Policy | ❌ |
| Trust Marks (OP/RP) | ❌ |


## 1. Estrutura de Diretórios

```bash
oidf-lab/
  docker-compose.yaml
  ta/
    config.yaml
    data/
      metadata-policy.json
      signing/ 
      storage/   
  ia/
    config.yaml
    data/
      metadata-policy.json
      signing/
      storage/
  rp/
    config.yaml     # whoami-rp
    data/          
```


## 2. Subindo a composição

```bash
cd oidf-lab
docker compose up -d
```

Ver logs do TA e IA:

```bash
docker compose logs ta ia --tail=30
```

Você deve ver o Lighthouse ouvindo:

- TA -> porta **7672**
- IA -> porta **7673**

## 3. Inspecionar Entity Configurations

### Trust Anchor (TA)

```bash
curl -s http://localhost:7672/.well-known/openid-federation   | cut -d'.' -f2 | base64 -d 2>/dev/null | jq .
```

### Intermediate Authority (IA)

```bash
curl -s http://localhost:7673/.well-known/openid-federation   | cut -d'.' -f2 | base64 -d 2>/dev/null | jq .
```

Você deverá ver:

- TA → `iss = sub = http://ta:7672`
- IA → `iss = sub = http://ia:7673`
- IA contém `authority_hints: ["http://ta:7672"]`

## 4. Enroll da IA no TA

```bash
curl "http://localhost:7672/enroll?sub=http://ia:7673&entity_type=federation_entity"
```

Listar entidades conhecidas pelo TA:

```bash
curl http://localhost:7672/list
```

Saída esperada:

```json
["http://ia:7673"]
```

## 5. Resolver IA no TA (trust chain completa)

```bash
curl -s "http://localhost:7672/resolve?sub=http://ia:7673&trust_anchor=http://ta:7672"   | cut -d'.' -f2 | base64 -d 2>/dev/null | jq .
```

Saída: metadados finais + trust chain TA → IA.

# 6. Adicionando o Relying Party (RP - whoami-rp)

O `whoami-rp` foi construído localmente pois não existe imagem pública.

### Build da imagem

```bash
cd whoami-rp-src
docker build -t oidfed/whoami-rp .
```

### Subir junto com TA/IA

```bash
cd oidf-lab
docker compose up -d
```

Ver logs:

```bash
docker compose logs rp --tail=50
```

Esperado:

```
Serving on :7680
```


## 7. Testar Entity Configuration do RP

```bash
curl -s http://localhost:7680/.well-known/openid-federation   | cut -d'.' -f2 | base64 -d 2>/dev/null | jq .
```

Isso deve retornar:

- `iss = sub = http://rp:7680`
- chave pública do RP
- authority_hints -> IA


## 8. Enroll do RP na IA

```bash
curl "http://localhost:7673/enroll?sub=http://rp:7680&entity_type=openid_relying_party"
```

Listar no IA:

```bash
curl "http://localhost:7673/list"
```

Deve incluir:

```json
["http://rp:7680"]
```


## 9. Resolver o RP no TA (cadeia TA -> IA -> RP)

```bash
curl -s "http://localhost:7672/resolve?sub=http://rp:7680&trust_anchor=http://ta:7672"   | cut -d'.' -f2 | base64 -d 2>/dev/null | jq .
```

Essa saída deve mostrar:

### Metadados finais do RP:
- application_type
- client_name
- redirect_uris
- response_types
- grant_types
- jwks

### Trust chain completa:
```
TA -> IA -> RP
```


# 10. Comandos de Validação do Fluxo (Resumo)

### 1) Ver TA
```bash
curl -s http://localhost:7672/.well-known/openid-federation | cut -d'.' -f2 | base64 -d
```

### 2) Ver IA
```bash
curl -s http://localhost:7673/.well-known/openid-federation | cut -d'.' -f2 | base64 -d
```

### 3) Ver RP
```bash
curl -s http://localhost:7680/.well-known/openid-federation | cut -d'.' -f2 | base64 -d
```

### 4) Enroll IA -> TA  
```bash
curl "http://localhost:7672/enroll?sub=http://ia:7673&entity_type=federation_entity"
```

### 5) Enroll RP -> IA  
```bash
curl "http://localhost:7673/enroll?sub=http://rp:7680&entity_type=openid_relying_party"
```

### 6) Resolve IA
```bash
curl -s "http://localhost:7672/resolve?sub=http://ia:7673&trust_anchor=http://ta:7672"   | cut -d'.' -f2 | base64 -d | jq .
```

### 7) Resolve RP (fluxo completo)
```bash
curl -s "http://localhost:7672/resolve?sub=http://rp:7680&trust_anchor=http://ta:7672"   | cut -d'.' -f2 | base64 -d | jq .
```
