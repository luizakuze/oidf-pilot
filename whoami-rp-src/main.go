package main

import (
	"log"

	"github.com/lestrrat-go/jwx/v3/jwa"

	"github.com/go-oidfed/lib"
)

const fedSigningKeyName = "fed.signing.key"
const oidcSigningKeyName = "oidc.signing.key"

func main() {
	mustLoadConfig()
	initKeys(fedSigningKeyName, oidcSigningKeyName)
	for _, c := range conf.TrustMarks {
		if err := c.Verify(
			conf.EntityID, "",
			oidfed.NewTrustMarkSigner(getKey(fedSigningKeyName), jwa.ES512()),
		); err != nil {
			log.Fatal(err)
		}
	}
	if conf.UseResolveEndpoint {
		oidfed.DefaultMetadataResolver = oidfed.SmartRemoteMetadataResolver{}
	}
	initServer()
}
