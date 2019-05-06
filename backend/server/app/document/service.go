package document

import (
	"github.com/Mushus/trashbox/backend/server/app/repository"
	"golang.org/x/xerrors"
)

type Service struct {
	repository Repository
}

func ProvideService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s Service) RestoreDocument(title string) (*Document, error) {
	prop, err := s.repository.Get(title)
	if err != nil {
		if xerrors.Is(err, repository.ErrDocumentNotFound) {
			return nil, ErrDocumentNotFound
		}
		return nil, xerrors.Errorf("failed to get Document: %w", err)
	}

	return NewDocument(prop), nil
}

func (s Service) SaveDocument(d *Document) error {
	prop := d.ToProp()

	if err := s.repository.Put(prop); err != nil {
		return xerrors.Errorf("failed to save error: %w", err)
	}
	return nil
}
