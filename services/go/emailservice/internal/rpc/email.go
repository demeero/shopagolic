package rpc

import (
	"context"

	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	emailpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/email/v1beta1"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Email struct {
	emailpb.UnimplementedEmailServiceServer
}

func NewEmail() *Email {
	return &Email{}
}

func (p *Email) SendOrderConfirmation(ctx context.Context, req *emailpb.SendOrderConfirmationRequest) (*emailpb.SendOrderConfirmationResponse, error) {
	if err := validation.Validate(req.Email, validation.Required, is.EmailFormat); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	zaplogger.FromCtx(ctx).Info("request to send order confirmation email has been received",
		zap.String("email", req.GetEmail()))
	return &emailpb.SendOrderConfirmationResponse{}, nil
}
