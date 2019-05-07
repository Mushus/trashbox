//go:generate mockgen -destination ./asset/repository.go -self_package github.com/Mushus/trashbox/backend/server/app/asset -package=user github.com/Mushus/trashbox/backend/server/app/asset Repository
//go:generate mockgen -destination ./cache/asset/repository.go -self_package github.com/Mushus/trashbox/backend/server/app/cache/asset -package=user github.com/Mushus/trashbox/backend/server/app/cache/asset Repository
//go:generate mockgen -destination ./document/repository.go -self_package github.com/Mushus/trashbox/backend/server/app/document -package=user github.com/Mushus/trashbox/backend/server/app/document Repository
//go:generate mockgen -destination ./user/repository.go -self_package github.com/Mushus/trashbox/backend/testutil/user -package=user github.com/Mushus/trashbox/backend/server/app/user Repository

package testutil
