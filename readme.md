## Instância do OpenID Federation (Lighthouse TA + IA)

Este repositório contém um pequeno laboratório de OpenID Federation baseado no [Lighthouse](https://go-oidfed.github.io/lighthouse/).

O objetivo deste lab é emular os blocos centrais do piloto eduGAIN OpenID Federation:

- Um **Trust Anchor (TA)** rodando Lighthouse (`oidf-ta`, entity_id `http://ta:7672`)
- Uma **Intermediate Authority / Federation Entity (IA)** também rodando Lighthouse (`oidf-ia`, entity_id `http://ia:7673`)
- Uma **trust chain** funcional entre TA -> IA, usando os endpoints `/enroll`, `/list` e `/resolve`.

### Checklist

Planejamento contendo o que já foi desenvolvido e o que ainda falta: 

- [x] Um TA  
- [x] Uma IA  
- [ ] Um OP (OpenID Provider) 
- [ ] Um RP (Relying Party / Client)   
- [ ] Política de metadados (metadata policy)  
- [ ] Trust marks para OP e RP (no mínimo 1)  
- [ ] Documentação de como os OP/RP estão subordinados ao TA/IA


### 1. Estrutura de diretórios

```text
oidf-lab/
  docker-compose.yaml
  ta/
    config.yaml
    data/
      metadata-policy.json
      signing/   # ignorado pelo git
      storage/   # ignorado pelo git
  ia/
    config.yaml
    data/
      metadata-policy.json
      signing/   # ignorado pelo git
      storage/   # ignorado pelo git
```

### 2. Como executar o lab

A partir da raiz do repositório:

```bash
cd oidf-lab
docker compose up -d
```

Verifique se as duas instâncias do Lighthouse estão rodando:

```bash
docker compose logs ta ia --tail=30
```

Você deve ver o Fiber ouvindo nas portas `7672` (TA) e `7673` (IA).

### 3. Inspecionar as entity configurations

**Trust Anchor (TA):**

```bash
curl -s http://localhost:7672/.well-known/openid-federation \
  | cut -d'.' -f2 | base64 -d 2>/dev/null; echo
```

**Intermediate Authority (IA):**

```bash
curl -s http://localhost:7673/.well-known/openid-federation \
  | cut -d'.' -f2 | base64 -d 2>/dev/null; echo
```

Você deve ver:

- `iss` e `sub` iguais a `http://ta:7672` para o TA  
- `iss` e `sub` iguais a `http://ia:7673` para a IA  
- `authority_hints` na IA apontando para `http://ta:7672`

### 4. Fazer o enrollment da IA na TA

```bash
curl "http://localhost:7672/enroll?sub=http://ia:7673&entity_type=federation_entity"
```

Esse comando faz o TA buscar e validar a entity configuration da IA e armazenar uma **entity statement** TA -> IA no storage dele.

Listar as entidades conhecidas pelo TA:

```bash
curl "http://localhost:7672/list"
```

Saída esperada:

```json
["http://ia:7673"]
```

### 5. Resolver a trust chain (TA -> IA)

```bash
curl "http://localhost:7672/resolve?sub=http://ia:7673&trust_anchor=http://ta:7672"
```

Esse comando retorna um `resolve-response+jwt` contendo:

- Os metadados finais resolvidos para `http://ia:7673`; e  
- A trust chain ancorada no TA (`http://ta:7672`).