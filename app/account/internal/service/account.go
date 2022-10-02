package service

import (
	"context"

	pb "yaosiqi525/cloud-platform/app/account/api/v1"
	"yaosiqi525/cloud-platform/app/account/internal/biz"
)

type AccountService struct {
	pb.UnimplementedAccountServer

	uc *biz.AccountUsecase
}

func NewAccountService(uc *biz.AccountUsecase) *AccountService {
	return &AccountService{uc: uc}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CommonReply, error) {
	_, err := s.uc.CreateAccount(ctx, &biz.Account{
		Account: req.AccountInfo.Account,
		Name:    req.AccountInfo.Name,
		Pwd:     req.AccountInfo.Pwd,
		Status:  1,
	})
	if err != nil {
		return &pb.CommonReply{
			Code: int32(pb.ErrorReason_UNKNOW),
			Msg:  err.Error(),
		}, nil
	}
	return &pb.CommonReply{
		Code: int32(pb.ErrorReason_SUCCESS),
		Msg:  "",
	}, nil
}
func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountReply, error) {
	return &pb.UpdateAccountReply{}, nil
}
func (s *AccountService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountReply, error) {
	return &pb.DeleteAccountReply{}, nil
}
func (s *AccountService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	return &pb.GetAccountReply{}, nil
}
func (s *AccountService) ListAccount(ctx context.Context, req *pb.ListAccountRequest) (*pb.ListAccountReply, error) {
	return &pb.ListAccountReply{}, nil
}
