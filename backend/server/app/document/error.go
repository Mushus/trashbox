package document

import "golang.org/x/xerrors"

var (
	// ErrDocumentNotFound ドキュメントが存在しなかったエラー
	ErrDocumentNotFound = xerrors.New("document not found")
)
