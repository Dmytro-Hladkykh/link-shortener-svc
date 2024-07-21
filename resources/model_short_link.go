/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources
import "gitlab.com/tokend/go/xdr"
import (
	"time"
)

type ShortLink struct {
	Clicks *int32 `json:"clicks,omitempty"`
	CreatedAt *date-time `json:"created_at,omitempty"`
	Id *int64 `json:"id,omitempty"`
	OriginalUrl *string `json:"original_url,omitempty"`
	ShortCode *string `json:"short_code,omitempty"`
}
