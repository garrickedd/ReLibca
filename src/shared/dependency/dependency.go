package dependency

import (
	"github.com/garickedd/ReLibca/src/api/config"
	contractRepository "github.com/garickedd/ReLibca/src/domain/repository"
	infraRepository "github.com/garickedd/ReLibca/src/infrastructure/persistence/repository"
)

func GetUserRepository(cfg *config.Config) contractRepository.UserRepository {
	return infraRepository.NewUserRepository(cfg)
}
