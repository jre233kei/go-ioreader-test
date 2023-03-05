package usecase

import "io"

// go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=mock_$GOFILE
type SampleRepository interface {
	Read(r io.Reader) error
}
