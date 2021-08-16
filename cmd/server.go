package main

import (
	"database/sql"
	"log"
	"net"

	"achuala.in/velolimits/cfg"
	"achuala.in/velolimits/pbgen/velolimits"
	"achuala.in/velolimits/svc"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := cfg.NewConfig()

	db, err := sql.Open(cfg.String("dbdriver"), cfg.DatabaseConnectionString())
	if err != nil {
		log.Fatal("db_connect_failed: ", err)
	}
	defer db.Close()

	gorm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("orm_init_failed: ", err)
	}

	lis, err := net.Listen("tcp", cfg.String("listen"))
	if err != nil {
		log.Fatal("listen_failed", err)
	}
	s := grpc.NewServer()
	velolimits.RegisterVelocityLimitServer(s, svc.NewLimitRulesSvc(gorm))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
