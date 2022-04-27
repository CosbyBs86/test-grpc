package grpcimplementation

import (
	"fmt"
	"io"
	"time"
)

type ServerGrpc struct{}

func (s *ServerGrpc) mustEmbedUnimplementedGrpcServiceServer() {}

func (s *ServerGrpc) AppendItems(stream GrpcService_AppendItemsServer) error {
	cont := 0
	go func() {
		for {
			time.Sleep(time.Second)
			if cont == 0 {
				continue
			}
			fmt.Printf("%d ops/s\n", cont)
			cont = 0
		}
	}()
	start := time.Now()
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Streaming complated. Elapsed time: %5.2f seconds\n", time.Now().Sub(start).Seconds())

			return stream.SendAndClose(&EtsBaseResponse{
				ErrorCode: 0,
				BaseResponse: &BaseResponse{
					Success:    true,
					Message:    "",
					Parameters: []string{},
				}})
		}
		if err != nil {
			return err
		}
		cont++
	}

	return nil
}
