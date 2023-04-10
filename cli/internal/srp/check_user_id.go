package srp

import "github.com/finallly/srp_protocol/cli/internal/db"

func CheckUserID(userID string) bool {
	count := db.GetCount(userID)
	return count == 1
}
