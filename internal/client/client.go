/*
 * Copyright (c) Matthew Hartstonge <matt@mykro.co.nz>
 * SPDX-License-Identifier: MPL-2.0
 */

package client

import (
	// Standard Library Imports
	"net/http"
	"net/url"
	"time"

	// External Imports
	"github.com/FusionAuth/go-client/pkg/fusionauth"
)

// New returns a new FusionAuth API client.
func New(baseURL *url.URL, apiToken string, tenantID string) *Client {
	// client configuration for data sources and resources
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	return &Client{
		API:      fusionauth.NewClient(httpClient, baseURL, apiToken),
		TenantID: tenantID,
	}
}

// Client defines the encapsulated FusionAuth Client passed to terraform data
// and resources.
type Client struct {
	// API stores the underlying fusionauth api client
	API *fusionauth.FusionAuthClient
	// TenantID stores the user configured tenant id for requests that require a
	// tenant.
	TenantID string
}
