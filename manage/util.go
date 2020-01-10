package manage

import (
	"log"
	"net/url"
	"strings"

	"gopkg.in/oauth2.v3/errors"
)

type (
	// ValidateURIHandler validates that redirectURI is contained in baseURI
	ValidateURIHandler func(baseURI, redirectURI string) error
)

// DefaultValidateURI validates that redirectURI is contained in baseURI
func DefaultValidateURI(domain string, redirectURI string) error {
	bases := strings.Split(domain, ",")
	for i := 0; i < len(bases); i++ {
		log.Println("DefaultValidateURI baseURI", bases[i], "redirectURI", redirectURI)
		base, err := url.Parse(bases[i])
		if err != nil {
			return err
		}
		redirect, err := url.Parse(redirectURI)
		if err != nil {
			return err
		}
		if strings.HasSuffix(redirect.Host, base.Host) {
			return nil
		}
	}
	return errors.ErrInvalidRedirectURI
}
