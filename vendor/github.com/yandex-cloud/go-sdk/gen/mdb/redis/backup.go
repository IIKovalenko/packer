// Code generated by sdkgen. DO NOT EDIT.

// nolint
package redis

import (
	"context"

	"google.golang.org/grpc"

	redis "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1alpha"
)

// BackupServiceClient is a redis.BackupServiceClient with
// lazy GRPC connection initialization.
type BackupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ redis.BackupServiceClient = &BackupServiceClient{}

// Get implements redis.BackupServiceClient
func (c *BackupServiceClient) Get(ctx context.Context, in *redis.GetBackupRequest, opts ...grpc.CallOption) (*redis.Backup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewBackupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements redis.BackupServiceClient
func (c *BackupServiceClient) List(ctx context.Context, in *redis.ListBackupsRequest, opts ...grpc.CallOption) (*redis.ListBackupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewBackupServiceClient(conn).List(ctx, in, opts...)
}
