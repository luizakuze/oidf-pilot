## OIDF Pilot

### Referências
- https://go-oidfed.github.io/lighthouse/deployment/caddy/
- 

### Comandos de teste

```
curl -s http://localhost:7672/.well-known/openid-federation \
| awk -F. '{print $2}' \
| tr '_-' '/+' \
| base64 -d 2>/dev/null \
| jq .
```