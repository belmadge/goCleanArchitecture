package main

import (
	"log"
	"net"
	"net/http"

	api "github.com/belmadge/goCleanArchitecture/api/http"
	"github.com/belmadge/goCleanArchitecture/internal/app/service"
	"github.com/belmadge/goCleanArchitecture/internal/db"
	"github.com/belmadge/goCleanArchitecture/internal/graphql"
	grpcService "github.com/belmadge/goCleanArchitecture/internal/grpc"
	"github.com/belmadge/goCleanArchitecture/pkg/repository"
	pb "github.com/belmadge/goCleanArchitecture/proto"
	"google.golang.org/grpc"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(repo)
	handler := api.NewOrderHandler(repo)

	go func() {
		http.HandleFunc("/order", handler.ListOrders)
		log.Println("HTTP Server started at :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterOrderServiceServer(grpcServer, grpcService.NewOrderServiceServer(orderService))
		log.Println("GRPC Server started at :50051")
		log.Fatal(grpcServer.Serve(lis))
	}()

	go func() {
        schema, err := graphql.NewSchema(orderService)
        if err != nil {
            log.Fatalf("Failed to create GraphQL schema: %v", err)
        }

        h := handler.New(&handler.Config{
            Schema: &schema,
            Pretty: true,
        })

        http.Handle("/graphql", h)
        log.Println("GraphQL Server started at :8081")
        log.Fatal(http.ListenAndServe(":8081", nil))
    }()

	select {}
}
