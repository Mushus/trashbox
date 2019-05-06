//go:generate mockgen -source ../server/app/asset/repository.go -destination ./asset/repository.go
//go:generate mockgen -source ../server/app/cache/asset/repository.go -destination ./cache/asset/repository.go
//go:generate mockgen -source ../server/app/document/repository.go -destination ./document/repository.go
//go:generate mockgen -source ../server/app/user/repository.go -destination ./user/repository.go

package testutil
