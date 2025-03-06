package gapi

import (
	"fmt"

	db "github.com/MohamedMosalm/banque/db/sqlc"
	token "github.com/MohamedMosalm/banque/maker"
	"github.com/MohamedMosalm/banque/pb"
	"github.com/MohamedMosalm/banque/util"
)

type Server struct {
	pb.UnimplementedBanqueServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
