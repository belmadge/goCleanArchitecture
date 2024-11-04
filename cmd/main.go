package main

import (
	"log"
	"net"
	"net/http"

	"github.com/belmadge/goCleanArchitecture/pkg/api/graphql"
	"github.com/belmadge/goCleanArchitecture/pkg/api/graphql/generated"
	grpcService "github.com/belmadge/goCleanArchitecture/pkg/api/grpc"
	"github.com/belmadge/goCleanArchitecture/pkg/api/rest"
	"github.com/belmadge/goCleanArchitecture/pkg/database"
	"github.com/belmadge/goCleanArchitecture/pkg/services"
	pb "github.com/belmadge/goCleanArchitecture/proto"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	svc := &services.OrderService{DB: db}

	http.HandleFunc("/order", rest.ListOrdersHandler(svc))

	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{OrderService: svc},
	}))

	http.Handle("/graphql", gqlHandler)
	http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Erro ao iniciar gRPC: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &grpcService.OrderGRPCService{OrderService: svc})

	go grpcServer.Serve(lis)
	log.Println("Servidores REST, GraphQL e gRPC iniciados nas portas 8080 (REST/GraphQL) e 50051 (gRPC)")
	http.ListenAndServe(":8080", nil)
}
