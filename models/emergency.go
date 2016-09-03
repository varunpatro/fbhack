package models

import "time"

type Emergency struct {
	CreatedAt   time.Time    `json:"createdAt"`
	PendingList []UserStatus `json:"pendingList"`
	SafeList    []UserStatus `json:"safeList"`
	UnsafeList  []UserStatus `json:"unsafeList"`
}
