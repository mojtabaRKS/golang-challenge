package command

import (
	"context"
	"fmt"
	"log"
	"sikabiz/internal/api/rest"
	user_handler "sikabiz/internal/api/rest/handlers/user"
	"sikabiz/internal/config"
	"sikabiz/internal/infra/mysql"
	"sikabiz/internal/infra/redis"
	"sikabiz/internal/repositories"
	user_service "sikabiz/internal/services/user"

	"sync"

	"github.com/spf13/cobra"
)

type Server struct{}

func (cmd Server) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "run server",
		Run: func(_ *cobra.Command, _ []string) {
			cmd.main(cfg, ctx)
		},
	}
}

func (cmd *Server) main(cfg *config.Config, ctx context.Context) {
	var wg sync.WaitGroup

	db, err := mysql.NewClient(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to connect to mysql database %s", err.Error())
		return
	}
	gormDB, err := mysql.NewGormWithInstance(db, cfg.AppDebug)
	if err != nil {
		log.Fatalf("failed to connect to mysql database: err : %s", err.Error())
		return
	}

	_, err = redis.NewClient(ctx, *cfg)
	if err != nil {
		log.Fatalf("failed to connect to redis, %s", err.Error())
	}

	err = mysql.Migrate(db)
	if err != nil {
		log.Fatalf("mysql migration failed: %s", err.Error())
	}

	server := rest.New()

	// intialize repositories
	userRepo := repositories.NewUserRepository(gormDB)

	// intialize services
	userService := user_service.New(userRepo)

	// intialize handlers
	userHandler := user_handler.New(userService)

	server.SetupAPIRoutes(userHandler)
	if err := server.Serve(ctx, fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
