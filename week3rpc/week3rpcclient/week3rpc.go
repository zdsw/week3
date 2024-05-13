// Code generated by goctl. DO NOT EDIT.
// Source: week3rpc.proto

package week3rpcclient

import (
	"context"

	"zg4-1/src/rk/week3/week3rpc/week3rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ManySearchRequest        = week3rpc.ManySearchRequest
	ManySearchResponse       = week3rpc.ManySearchResponse
	PaginationSearchRequest  = week3rpc.PaginationSearchRequest
	PaginationSearchResponse = week3rpc.PaginationSearchResponse
	RollbackRequest          = week3rpc.RollbackRequest
	RollbackResponse         = week3rpc.RollbackResponse
	SearchRequest            = week3rpc.SearchRequest
	SearchResponse           = week3rpc.SearchResponse
	SynchronousRequest       = week3rpc.SynchronousRequest
	SynchronousResponse      = week3rpc.SynchronousResponse
	UploadRequest            = week3rpc.UploadRequest
	UploadResponse           = week3rpc.UploadResponse

	Week3rpc interface {
		// 单条件搜索接口
		Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
		// 多条件搜索接口
		ManySearch(ctx context.Context, in *ManySearchRequest, opts ...grpc.CallOption) (*ManySearchResponse, error)
		// 分页搜索接口
		PaginationSearch(ctx context.Context, in *PaginationSearchRequest, opts ...grpc.CallOption) (*PaginationSearchResponse, error)
		// es同步
		Synchronous(ctx context.Context, in *SynchronousRequest, opts ...grpc.CallOption) (*SynchronousResponse, error)
		// 上传
		Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
		// 回滚
		Rollback(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*RollbackResponse, error)
	}

	defaultWeek3rpc struct {
		cli zrpc.Client
	}
)

func NewWeek3rpc(cli zrpc.Client) Week3rpc {
	return &defaultWeek3rpc{
		cli: cli,
	}
}

// 单条件搜索接口
func (m *defaultWeek3rpc) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	client := week3rpc.NewWeek3RpcClient(m.cli.Conn())
	return client.Search(ctx, in, opts...)
}

// 多条件搜索接口
func (m *defaultWeek3rpc) ManySearch(ctx context.Context, in *ManySearchRequest, opts ...grpc.CallOption) (*ManySearchResponse, error) {
	client := week3rpc.NewWeek3RpcClient(m.cli.Conn())
	return client.ManySearch(ctx, in, opts...)
}

// 分页搜索接口
func (m *defaultWeek3rpc) PaginationSearch(ctx context.Context, in *PaginationSearchRequest, opts ...grpc.CallOption) (*PaginationSearchResponse, error) {
	client := week3rpc.NewWeek3RpcClient(m.cli.Conn())
	return client.PaginationSearch(ctx, in, opts...)
}

// es同步
func (m *defaultWeek3rpc) Synchronous(ctx context.Context, in *SynchronousRequest, opts ...grpc.CallOption) (*SynchronousResponse, error) {
	client := week3rpc.NewWeek3RpcClient(m.cli.Conn())
	return client.Synchronous(ctx, in, opts...)
}

// 上传
func (m *defaultWeek3rpc) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	client := week3rpc.NewWeek3RpcClient(m.cli.Conn())
	return client.Upload(ctx, in, opts...)
}

// 回滚
func (m *defaultWeek3rpc) Rollback(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*RollbackResponse, error) {
	client := week3rpc.NewWeek3RpcClient(m.cli.Conn())
	return client.Rollback(ctx, in, opts...)
}
