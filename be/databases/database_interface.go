package databases

import "context"

type (
	TkbaiInterface interface {
		ViewToeflDataAll(ctx context.Context, start, length string) (result []ToeflCertificate, err error)
		CountToeflDataAll(ctx context.Context) (result int64, err error)
		ViewToeflDataByIDAndName(ctx context.Context, certificateId, certificateHolder string) (result ToeflCertificate, err error)
		CreateCertificateBulk(ctx context.Context, certificates []ToeflCertificate) (rowsAffected int64, err error)
		CreateToeflCertificate(ctx context.Context, certificate ToeflCertificate) (rowsAffected int64, err error)
	}
)
