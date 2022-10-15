package testutil

import (
	"context"
	"os"

	"go.dagger.io/dagger/internal/buildkitd"
	"go.dagger.io/dagger/sdk/go/dagger"
)

type QueryOptions struct {
	Variables map[string]any
	Operation string
}

func Query(query string, res any, opts *QueryOptions, clientOpts ...dagger.ClientOpt) error {
	ctx := context.Background()

	if opts == nil {
		opts = &QueryOptions{}
	}
	if opts.Variables == nil {
		opts.Variables = make(map[string]any)
	}

	c, err := dagger.Connect(ctx, clientOpts...)
	if err != nil {
		return err
	}
	defer c.Close()

	return c.Do(ctx,
		&dagger.Request{
			Query:     query,
			Variables: opts.Variables,
			OpName:    opts.Operation,
		},
		&dagger.Response{Data: &res},
	)
}

func SetupBuildkitd() error {
	host, err := buildkitd.StartGoModBuildkitd(context.Background())
	if err != nil {
		return err
	}
	os.Setenv("BUILDKIT_HOST", host)
	return nil
}