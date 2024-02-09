package entity

import "context"

/*================================ ENTITY / MODELS ==============================*/
// interface database
type (
	ValidatorInterface interface {
		// Sertif Table
		ViewSertifTableByID(ctx context.Context, certificate_id, ip string) (Sertifikat, error)
	}
)
