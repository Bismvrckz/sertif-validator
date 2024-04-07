package entity

import "context"

/*================================ ENTITY / MODELS ==============================*/

type (
	ValidatorInterface interface {
		ViewTkbaiCertByID(ctx context.Context, certificateId, ip string) (ToeflCertificate, error)
		ViewToeflDataAll(ctx context.Context, start, length, ip string) (toeflCertificate []ToeflCertificate, err error)
		CreateCertificate(ctx context.Context, certificates []ToeflCertificate, ip string) (rowsAffected int64, err error)
		CountToeflDataAll(ctx context.Context, ip string) (totalRows int64, err error)
	}
)
