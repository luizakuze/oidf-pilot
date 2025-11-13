## OIDF Pilot

### Sobre o repositório 
São estudos sobre a implementação [Lighthouse](https://go-oidfed.github.io/lighthouse/)

- Ela pode atuar como Trust Anchor (TA), Autoridade Intermediária, Resolver de metadados e Emissor de Trust Marks (selos).

- Publica sua Entity Configuration (quem você é, suas chaves públicas e endpoints).

- Emite/consome Entity Statements (JWTs assinados que formam a cadeia de confiança).

- Resolve metadados de entidades aplicando políticas (ex.: exigir certos claims, filtrar atributos).

### Referências
- https://go-oidfed.github.io/lighthouse/deployment/caddy/
- https://github.com/GEANT/edugain-oidf-pilot/tree/main

### Obtendo a Entity Configuration

#### Payload
> Subir a composição antes
```bash
curl -s http://localhost:7672/.well-known/openid-federation \
| awk -F. '{print $2}' \     
| tr '_-' '/+' \            
| base64 -d 2>/dev/null \   
| jq .                      
```

#### Saída
```bash
{
  "authority_hints": ["http://localhost:7672"],
  "exp": 176...,
  "iat": 176...,
  "iss": "http://localhost:7672",
  "sub": "http://localhost:7672",
  "jwks": {
    "keys": [
      {
        "kty": "EC",
        "crv": "P-256",
        "use": "sig",
        "kid": "....",
        "x": "....",
        "y": "...."
      }
    ]
  },
  "metadata": {
    "federation_entity": {
      "display_name": "Minha IA Local",
      "federation_fetch_endpoint": "http://localhost:7672/fetch",
      "federation_list_endpoint": "http://localhost:7672/list",
      "federation_resolve_endpoint": "http://localhost:7672/resolve",
      "federation_trust_mark_endpoint": "http://localhost:7672/trustmark",
      "federation_trust_mark_list_endpoint": "http://localhost:7672/trustmark/list"
    }
  },
  "organization_name": "Minha Org"
}
```