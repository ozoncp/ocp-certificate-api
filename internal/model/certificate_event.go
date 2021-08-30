package model

// CertificateEvent - is an interface that describes information about certificate id
type CertificateEvent struct {
	Action    string
	ID        uint64
	Timestamp int64
}
