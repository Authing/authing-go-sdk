package enum

type SortByEnum string

const (
	SortByCreatedAtDesc SortByEnum = "CREATEDAT_DESC"
	SortByCreatedAtAsc  SortByEnum = "CREATEDAT_ASC"
	SortByUpdatedAtDesc SortByEnum = "UPDATEDAT_DESC"
	SortByUpdatedAtAsc  SortByEnum = "UPDATEDAT_ASC"
)
