// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package miniogw

import (
	"context"

	"uplink"
)

type uplinkProjectKey struct{}

// WithUplinkProject injects project into ctx under a specific key. Use
// GetUplinkProject to retrieve project from ctx.
func WithUplinkProject(ctx context.Context, project *uplink.Project) context.Context {
	return context.WithValue(ctx, uplinkProjectKey{}, project)
}

// GetUplinkProject retrieves libuplink's *Project from ctx and reports whether
// it was successful.
func GetUplinkProject(ctx context.Context) (*uplink.Project, bool) {
	project, ok := ctx.Value(uplinkProjectKey{}).(*uplink.Project)
	return project, ok
}
