// Package container provides ...
package api

import (
    "../models"
    "github.com/interactiv/pimple"
)

type AbstractAPI struct {

    AccessToken models.AccessToken
}

func (abstractAPI *AbstractAPI) NewAPI(p *pimple.Pimple) {
    abstractAPI.AccessToken = p.Get("access_token").(models.AccessToken)
}
