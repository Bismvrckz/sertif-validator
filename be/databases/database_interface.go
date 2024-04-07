package databases

import "context"

type (
	TkbaiInterface interface {
		ViewToeflDataAll(ctx context.Context, start, length string) (result []ToeflCertificate, err error)
		CountToeflDataAll(ctx context.Context) (result int64, err error)
	}
)
