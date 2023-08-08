/*
 * Copyright (c) Matthew Hartstonge <matt@mykro.co.nz>
 * SPDX-License-Identifier: MPL-2.0
 */

package provider

import (
	"fmt"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// reportedCreateErrors reports any FusionAuth errors found when creating a resource.
// Returns whether errors have been reported.
func reportedCreateErrors(diags diag.Diagnostics, faErrs *fusionauth.Errors, err error, resource string) bool {
	return reportedErrors(diags, faErrs, err, "create", resource)
}

// reportedReadErrors adds any FusionAuth errors found when reading a resource.
// Returns whether errors have been reported.
func reportedReadErrors(diags diag.Diagnostics, faErrs *fusionauth.Errors, err error, resource string) bool {
	return reportedErrors(diags, faErrs, err, "read", resource)
}

// reportedUpdateErrors adds any FusionAuth errors found when updating a resource.
// Returns whether errors have been reported.
func reportedUpdateErrors(diags diag.Diagnostics, faErrs *fusionauth.Errors, err error, resource string) bool {
	return reportedErrors(diags, faErrs, err, "update", resource)
}

// reportedDeleteErrors adds any FusionAuth errors found when deleting a resource.
// Returns whether errors have been reported.
func reportedDeleteErrors(diags diag.Diagnostics, faErrs *fusionauth.Errors, err error, resource string) bool {
	return reportedErrors(diags, faErrs, err, "delete", resource)
}

// reportedErrors finds returned fusionauth errors or underlying client errors and adds them to the passed in diagnostics.
// Returns whether errors have been reported.
func reportedErrors(diags diag.Diagnostics, faErrs *fusionauth.Errors, err error, method string, resource string) bool {
	switch {
	case faErrs.Present():
		diags.AddError(
			fmt.Sprintf("Error attempting to %s a %s", method, resource),
			fmt.Sprintf("FusionAuth reported errors while attempting to %s a new fusionauth_%s resource. Please fix the following errors: %s", method, resource, faErrs.Error()),
		)

	default:
		diags.AddError(
			"FusionAuth client error",
			fmt.Sprintf("Unable to %s a fusionauth_%s resource, client returned: %s", method, resource, err),
		)
	}

	return diags.HasError()
}
