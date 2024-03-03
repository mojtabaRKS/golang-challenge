package rest

import "sikabiz/internal/api/rest/handlers/user"

func (s *Server) SetupAPIRoutes(userHandler *user.UserHandler) {
	r := s.engine

	r.POST("register", userHandler.Register)
}
