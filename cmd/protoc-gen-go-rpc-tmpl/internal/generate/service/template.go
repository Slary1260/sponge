package service

import (
	"math/rand"
	"text/template"
	"time"
)

func init() {
	var err error
	serviceLogicTmpl, err = template.New("serviceLogicTmpl").Parse(serviceLogicTmplRaw)
	if err != nil {
		panic(err)
	}
	serviceLogicTestTmpl, err = template.New("serviceLogicTestTmpl").Parse(serviceLogicTestTmplRaw)
	if err != nil {
		panic(err)
	}
	rpcErrCodeTmpl, err = template.New("rpcErrCode").Parse(rpcErrCodeTmplRaw)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
}

var (
	serviceLogicTmpl    *template.Template
	serviceLogicTmplRaw = `// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"

	"google.golang.org/grpc"

	//"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	//"github.com/zhufuyi/sponge/pkg/logger"

	serverNameExampleV1 "moduleNameExample/api/serverNameExample/v1"

	//"moduleNameExample/internal/cache"
	//"moduleNameExample/internal/dao"
	//"moduleNameExample/internal/ecode"
	//"moduleNameExample/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		{{- range .PbServices}}
		serverNameExampleV1.Register{{.Name}}Server(server, New{{.Name}}Server())
		{{- end}}
	})
}

{{- range .PbServices}}

var _ serverNameExampleV1.{{.Name}}Server = (*{{.LowerName}})(nil)

type {{.LowerName}} struct {
	serverNameExampleV1.Unimplemented{{.Name}}Server

	// example:
	//		iDao dao.{{.Name}}Dao
}

// New{{.Name}}Server create a server
func New{{.Name}}Server() serverNameExampleV1.{{.Name}}Server {
	return &{{.LowerName}}{
		// example:
		//		iDao: dao.New{{.Name}}Dao(
		//			model.GetDB(),
		//			cache.New{{.Name}}Cache(model.GetCacheType()),
		//		),
	}
}

{{- range .Methods}}
{{if eq .InvokeType 1}}
{{.Comment}}
func (s *{{.LowerServiceName}}) {{.MethodName}}(stream serverNameExampleV1.{{.ServiceName}}_{{.MethodName}}Server) error {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    ctx := interceptor.WrapServerCtx(stream.Context())
	//	    for {
	//	        req, err := stream.Recv()
	//	        if err != nil {
	//	    	    if err == io.EOF {
	//	    	        return stream.SendAndClose(&serverNameExampleV1.{{.Reply}}{
		    	            {{- range .ReplyFields}}
	//	    	            {{.Name}}: reply.{{.Name}},
		    	            {{- end}}
	//	    	        })
	//	    	    }
	//	    	    return err
	//	        }
	//
	//	        err = req.Validate()
	//	        if err != nil {
	//		        logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		        return ecode.StatusInvalidParams.Err()
	//	        }
	//
	// 	    reply, err := s.iDao.{{.MethodName}}(ctx, &model.{{.ServiceName}}{
				    {{- range .RequestFields}}
	//     	    {{.Name}}: req.{{.Name}},
				    {{- end}}
	//         })
	// 	    if err != nil {
	//			    logger.Warn("{{.MethodName}} error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			    return ecode.StatusInternalServerError.Err()
	//		    }
	//	    }
}
{{else if eq .InvokeType 2}}
{{.Comment}}
func (s *{{.LowerServiceName}}) {{.MethodName}}(req *serverNameExampleV1.{{.Request}}, stream serverNameExampleV1.{{.ServiceName}}_{{.MethodName}}Server) error {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    ctx := interceptor.WrapServerCtx(stream.Context())
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return ecode.StatusInvalidParams.Err()
	//	    }
	//
	//	    for i := 0; i < 3; i++ {
	// 	    reply, err := s.iDao.{{.MethodName}}(ctx, &model.{{.ServiceName}}{
				    {{- range .RequestFields}}
	//     	    {{.Name}}: req.{{.Name}},
				    {{- end}}
	//         })
	// 	    if err != nil {
	//			    logger.Warn("{{.MethodName}} error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			    return ecode.StatusInternalServerError.Err()
	//		    }
	//
	//	        err = stream.Send(&serverNameExampleV1.{{.Reply}}{
				    {{- range .ReplyFields}}
	//	            {{.Name}}: reply.{{.Name}},
				    {{- end}}
	//	        })
	//	        if err != nil {
	//			    logger.Warn("stream.Send error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//	    	    return err
	//	        }
	//	    }
	//	    return nil
}
{{else if eq .InvokeType 3}}
{{.Comment}}
func (s *{{.LowerServiceName}}) {{.MethodName}}(stream serverNameExampleV1.{{.ServiceName}}_{{.MethodName}}Server) error {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    ctx := interceptor.WrapServerCtx(stream.Context())
	//	    for {
	//	        req, err := stream.Recv()
	//	        if err != nil {
	//	    	    if err == io.EOF {
	//	    	        return nil
	//	    	    }
	//	    	    return err
	//	        }
	//
	//	        err = req.Validate()
	//	        if err != nil {
	//		        logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		        return ecode.StatusInvalidParams.Err()
	//	        }
	//
	// 	    reply, err := s.iDao.{{.MethodName}}(ctx, &model.{{.ServiceName}}{
				    {{- range .RequestFields}}
	//     	    {{.Name}}: req.{{.Name}},
				    {{- end}}
	//         })
	// 	    if err != nil {
	//			    logger.Warn("{{.MethodName}} error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			    return ecode.StatusInternalServerError.Err()
	//		    }
	//
	//	    	err = stream.Send(&serverNameExampleV1.{{.Reply}}{
				    {{- range .ReplyFields}}
	//			    {{.Name}}: reply.{{.Name}},
				    {{- end}}
	//	    	})
	// 	    if err != nil {
	//			    logger.Warn("stream.Send error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			    return ecode.StatusInternalServerError.Err()
	//		    }
	//	    }
}
{{else}}
{{.Comment}}
func (s *{{.LowerServiceName}}) {{.MethodName}}(ctx context.Context, req *serverNameExampleV1.{{.Request}}) (*serverNameExampleV1.{{.Reply}}, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
    // 	ctx = interceptor.WrapServerCtx(ctx)
    //
	// 	reply, err := s.iDao.{{.MethodName}}(ctx, &model.{{.ServiceName}}{
				{{- range .RequestFields}}
	//     	{{.Name}}: req.{{.Name}},
				{{- end}}
	//     })
	// 	if err != nil {
	//			logger.Warn("{{.MethodName}} error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &serverNameExampleV1.{{.Reply}}{
				{{- range .ReplyFields}}
	//     	{{.Name}}: reply.{{.Name}},
				{{- end}}
	//     }, nil
}
{{end}}
{{- end}}

// ---------- Do not delete or move this split line, this is the merge code marker ----------

{{- end}}
`

	serviceLogicTestTmpl    *template.Template
	serviceLogicTestTmplRaw = `// Code generated by https://github.com/zhufuyi/sponge
{{- range .PbServices}}
// Test_service_{{.LowerName}}_methods is used to test the {{.LowerName}} api
// Test_service_{{.LowerName}}_benchmark is used to performance test the {{.LowerName}} api
{{- end}}

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"

	serverNameExampleV1 "moduleNameExample/api/serverNameExample/v1"
	"moduleNameExample/configs"
	"moduleNameExample/internal/config"
)

{{- range .PbServices}}

// Test service {{.LowerName}} api via grpc client
func Test_service_{{.LowerName}}_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := serverNameExampleV1.New{{.Name}}Client(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{
{{- range .Methods}}
{{if eq .InvokeType 1}}
		{
			name: "{{.MethodName}}",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &serverNameExampleV1.{{.Request}}{
					{{- range .RequestFields}}
					{{.Name}}: {{.GoTypeZero}}, {{if .Comment}} {{.Comment}}{{end}}
					{{- end}}
				}

				stream, err := cli.{{.MethodName}}(context.Background())
				if err != nil {
					return nil, err
				}
				for i:=0; i<3; i++ {
					err = stream.Send(req)
					if err != nil {
						return nil, err
					}
				}
				return stream.CloseAndRecv()
			},
			wantErr: false,
		},
{{else if eq .InvokeType 2}}
		{
			name: "{{.MethodName}}",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &serverNameExampleV1.{{.Request}}{
					{{- range .RequestFields}}
					{{.Name}}: {{.GoTypeZero}}, {{if .Comment}} {{.Comment}}{{end}}
					{{- end}}
				}

				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				stream, err := cli.{{.MethodName}}(ctx, req)
				if err != nil {
					return nil, err
				}
				result := &serverNameExampleV1.{{.Reply}}{}
				for {
					select {
					case <-ctx.Done():
						return nil, stream.CloseSend()
					default:
						reply, err := stream.Recv()
						if err == ioEOF {
							return result, nil
						}
						if err != nil {
							return nil, err
						}
						result = reply
					}
				}
			},
			wantErr: false,
		},
{{else if eq .InvokeType 3}}
		{
			name: "{{.MethodName}}",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &serverNameExampleV1.{{.Request}}{
					{{- range .RequestFields}}
					{{.Name}}: {{.GoTypeZero}}, {{if .Comment}} {{.Comment}}{{end}}
					{{- end}}
				}

				stream, err := cli.{{.MethodName}}(context.Background())
				if err != nil {
					return nil, err
				}
				reply := &serverNameExampleV1.{{.Reply}}{}
				for i:=0; i<3; i++ {
					err = stream.Send(req)
					if err != nil {
						return nil, err
					}
					reply, err = stream.Recv()
					if err == ioEOF {
						return &serverNameExampleV1.{{.Reply}}{
							{{- range .ReplyFields}}
							{{.Name}}: reply.{{.Name}},
							{{- end}}
						}, nil
					}
					if err != nil {
						return nil, err
					}
				}
				return reply, stream.CloseSend()
			},
			wantErr: false,
		},
{{else}}
		{
			name: "{{.MethodName}}",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &serverNameExampleV1.{{.Request}}{
					{{- range .RequestFields}}
					{{.Name}}: {{.GoTypeZero}}, {{if .Comment}} {{.Comment}}{{end}}
					{{- end}}
				}

				return cli.{{.MethodName}}(ctx, req)
			},
			wantErr: false,
		},
{{end}}
{{- end}}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fn()
			if (err != nil) != tt.wantErr {
				t.Errorf("test '%s' error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			data, _ := json.MarshalIndent(got, "", "    ")
			fmt.Println(string(data))
		})
	}
}

// performance test service {{.LowerName}} api, copy the report to
// the browser to view when the pressure test is finished.
func Test_service_{{.LowerName}}_benchmark(t *testing.T) {
	err := config.Init(configs.Path("serverNameExample.yml"))
	if err != nil {
		panic(err)
	}
	if len(config.Get().GrpcClient) == 0 {
		t.Error("grpcClient is not set in serverNameExample.yml")
		return
	}
	host := fmt.Sprintf("%s:%d", config.Get().GrpcClient[0].Host, config.Get().GrpcClient[0].Port)
	protoFile := configs.Path("../api/serverNameExample/v1/{{.ProtoName}}")
	// If third-party dependencies are missing during the press test,
	// copy them to the project's third_party directory.
	dependentProtoFilePath := []string{
		configs.Path("../third_party"), // third_party directory
		configs.Path(".."),             // Previous level of third_party
	}

	tests := []struct {
		name    string
		fn      func() error
		wantErr bool
	}{
{{- range .Methods}}
{{if eq .InvokeType 1}}
		{
			name: "{{.MethodName}}",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &serverNameExampleV1.{{.Request}}{
					{{- range .RequestFields}}
					{{.Name}}: {{.GoTypeZero}}, {{if .Comment}} {{.Comment}}{{end}}
					{{- end}}
				}
				total := 100 // total number of requests

				options := []benchmark.Option{
					// runner.WithStreamCallCount(10), // steam count, need to import "github.com/bojand/ghz/runner"
				}

				b, err := benchmark.New(host, protoFile, "{{.MethodName}}", message, dependentProtoFilePath, total, options...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
{{else if eq .InvokeType 2}}
		{
			name: "{{.MethodName}}",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &serverNameExampleV1.{{.Request}}{
					{{- range .RequestFields}}
					{{.Name}}: {{.GoTypeZero}}, {{if .Comment}} {{.Comment}}{{end}}
					{{- end}}
				}
				total := 100 // total number of requests

				options := []benchmark.Option{
					// runner.WithStreamCallCount(10), // steam count, need to import "github.com/bojand/ghz/runner"
				}

				b, err := benchmark.New(host, protoFile, "{{.MethodName}}", message, dependentProtoFilePath, total, options...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
{{else if eq .InvokeType 3}}
		{
			name: "{{.MethodName}}",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &serverNameExampleV1.{{.Request}}{
					{{- range .RequestFields}}
					{{.Name}}: {{.GoTypeZero}}, {{if .Comment}} {{.Comment}}{{end}}
					{{- end}}
				}
				total := 100 // total number of requests

				options := []benchmark.Option{
					// runner.WithStreamCallCount(10), // steam count, need to import "github.com/bojand/ghz/runner"
				}

				b, err := benchmark.New(host, protoFile, "{{.MethodName}}", message, dependentProtoFilePath, total, options...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
{{else}}
		{
			name: "{{.MethodName}}",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &serverNameExampleV1.{{.Request}}{
					{{- range .RequestFields}}
					{{.Name}}: {{.GoTypeZero}}, {{if .Comment}} {{.Comment}}{{end}}
					{{- end}}
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "{{.MethodName}}", message, dependentProtoFilePath, total)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
{{end}}
{{- end}}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			if (err != nil) != tt.wantErr {
				t.Errorf("test '%s' error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
		})
	}
}

{{- end}}
`

	rpcErrCodeTmpl    *template.Template
	rpcErrCodeTmplRaw = `// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

{{- range .PbServices}}

// {{.LowerName}} business-level rpc error codes.
// the _{{.LowerName}}NO value range is 1~100, if the same error code is used, it will cause panic.
var (
	_{{.LowerName}}NO       = {{.RandNumber}}
	_{{.LowerName}}Name     = "{{.LowerName}}"
	_{{.LowerName}}BaseCode = errcode.RCode(_{{.LowerName}}NO)
// --blank line--
{{- range $i, $v := .Methods}}
	Status{{.MethodName}}{{.ServiceName}}   = errcode.NewRPCStatus(_{{.LowerServiceName}}BaseCode+{{$v.AddOne $i}}, "failed to {{.MethodName}} "+_{{.LowerServiceName}}Name)
{{- end}}

	// error codes are globally unique, adding 1 to the previous error code
)

// ---------- Do not delete or move this split line, this is the merge code marker ----------

{{- end}}
`
)
