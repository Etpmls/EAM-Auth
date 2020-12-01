// https://github.com/grpc/grpc-go/blob/15a78f19307d5faf10cfdd9d4e664c65a387cbd1/examples/helloworld/greeter_client/main.go
package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Etpmls/EM-Auth/src/application"
	"github.com/Etpmls/EM-Auth/src/application/protobuf"
	em "github.com/Etpmls/Etpmls-Micro"
	em_library "github.com/Etpmls/Etpmls-Micro/library"
	em_utils "github.com/Etpmls/Etpmls-Micro/utils"
	"github.com/afex/hystrix-go/hystrix"
	"time"
)

// Get User avatar
// 获取用户头像
func (this *client) User_GetAvatar(owner_id uint32, owner_type string) (string) {
	cl := em.NewClient()
	output := make(chan string, 1)
	errs := cl.Go("common", func() error {

		conn, err := cl.ConnectService(application.Service_AttachmentService)
		if err != nil {
			return err
		}
		defer conn.Close()
		c := protobuf.NewAttachmentClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.GetOne(ctx, &protobuf.AttachmentGetOne{
			Service:       em_library.Config.ServiceDiscovery.Service.Name,
			OwnerId:       owner_id,
			OwnerType:     owner_type,
		})
		if err != nil {
			em.LogError.Output(em_utils.MessageWithLineNum(err.Error()))
			return err
		}

		type res struct {
			Path string	`json:"path"`
		}
		var rsps res
		err = json.Unmarshal([]byte(r.GetData()), &rsps)
		if err != nil {
			return err
		}
		output <- rsps.Path
		return nil

	}, func(err error) error {
		output <- ""
		// do this when services are down
		return nil
	})

	select {
	case out := <-output:
		// success
		return out
	case <-errs:
		// failure
		return ""
	}
}

func (this *client) User_CreateAvatar(ctx context.Context, path string, owner_id uint32, owner_type string) error {
	cl := em.NewClient()
	err := cl.Do("common", func() error {
		// 1.Get token By Request
		token, err := em.NewAuth().GetTokenFromCtx(ctx)
		if err != nil {
			return err
		}

		// 2.Connect Service
		conn, err := cl.ConnectService(application.Service_AttachmentService)
		if err != nil {
			return err
		}
		defer conn.Close()
		c := protobuf.NewAttachmentClient(conn)
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// 3.Set Token
		cl.SetValueToClientHeader(&ctx, map[string]string{"token": token})

		r, err := c.Create(ctx, &protobuf.AttachmentCreate{
			Service:       em_library.Config.ServiceDiscovery.Service.Name,
			Path:		   path,
			OwnerId:       owner_id,
			OwnerType:     owner_type,
		})
		if err != nil {
			em.LogError.Output(em_utils.MessageWithLineNum(err.Error()))
			return err
		}

		if r.GetStatus() == em.SUCCESS_Status {
			return nil
		} else {
			em.LogError.Output(em_utils.MessageWithLineNum("Create failed!"))
			return errors.New("Create failed!")
		}

	},nil)

	return err
}

func (this *client) Setting_DiskCleanUp() error {
	cl := em.NewClient()
	err := hystrix.Do("common", func() error {

		conn, err := cl.ConnectService(application.Service_AttachmentService)
		if err != nil {
			return err
		}
		defer conn.Close()
		c := protobuf.NewAttachmentClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.DiskCleanUp(ctx, &protobuf.AttachmentDiskCleanUp{
			Service:       em_library.Config.ServiceDiscovery.Service.Name,
		})
		if err != nil {
			em.LogError.Output(em_utils.MessageWithLineNum(err.Error()))
			return err
		}

		if r.GetStatus() == em.SUCCESS_Status {
			return nil
		} else {
			em.LogError.Output(em_utils.MessageWithLineNum("Delete Failed!"))
			return errors.New("Delete Failed!")
		}

	}, nil)

	return err
}