## OIDF Pilot

### Referências
- https://go-oidfed.github.io/lighthouse/deployment/caddy/
- https://github.com/GEANT/edugain-oidf-pilot/tree/main

### Obtendo a Entity Configuration

#### Comando
> Subir a composição antes
```bash
curl -s http://localhost:7672/.well-known/openid-federation \
| awk -F. '{print $2}' \    # pega só o payload
| tr '_-' '/+' \            # converte base64url -> base64
| base64 -d 2>/dev/null \   # decodifica
| jq .                      # formata o JSON
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