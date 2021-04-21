package customer

import (
	"context"
	"github.com/amjadjibon/grpc_todo/pkg/postgres"
	"github.com/amjadjibon/grpc_todo/pkg/utils"
	"github.com/go-pg/pg/v10"
	"github.com/golang/protobuf/ptypes"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	UnimplementedCustomerServiceServer
	DB *pg.DB
}

func (s *Service) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	req.Customer.Id = uuid.NewV4().String()
	time, _ := ptypes.Timestamp(req.Customer.CreatedAt)

	if err := req.Validate(); err != nil {
		return nil, grpc.Errorf(codes.Internal, "validation error: %s", err)
	}

	request := postgres.Customer {
		Id:          req.Customer.Id,
		FirstName:   req.Customer.FirstName,
		LastName:    req.Customer.LastName,
		PhoneNumber: req.Customer.PhoneNumber,
		Email:       req.Customer.Email,
		Password:    utils.CreateHash(req.Customer.Password),
		CreatedAt:   time,
		UpdatedAt:   time,
	}

	_, err := s.DB.Model(&request).Insert()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}
	var response postgres.Customer
	err = s.DB.Model(&response).Where("phone_number = ?", req.Customer.PhoneNumber).First()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}

	createdAt, _ := ptypes.TimestampProto(response.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(response.UpdatedAt)
	res := Customer {
		FirstName:   response.FirstName,
		LastName:    response.LastName,
		PhoneNumber: response.PhoneNumber,
		Email:       response.Email,
		Password:    response.Password,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	return &CreateResponse{Customer: &res}, nil
}

func (s *Service) ReadAll(ctx context.Context, req *ReadAllRequest) (*ReadAllResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, grpc.Errorf(codes.Internal, "validation error: %s", err)
	}
	var customers []postgres.Customer

	err := s.DB.Model(&customers).Select()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "db err: %s", err)
	}

	var cus []*Customer

	res := ReadAllResponse{
		Customer: cus,
	}

	res.Validate()

	return &res, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func (s *Service) mustEmbedUnimplementedCustomerServiceServer() {}
