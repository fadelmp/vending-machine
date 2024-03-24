package injection

import (
	"vending-machine/comparator"
	"vending-machine/handler"
	"vending-machine/mapper"
	"vending-machine/repository"
	"vending-machine/usecase"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func VendingInjection(db *gorm.DB, redis *redis.Client) *handler.VendingHandler {

	baseMapper := mapper.NewBaseMapper()
	vendingMapper := mapper.NewVendingMapper()
	repo := repository.NewVendingRepository(db, redis)

	comparator := comparator.NewVendingComparator(repo)
	usecase := usecase.NewVendingUsecase(baseMapper, vendingMapper, repo, comparator)
	handler := handler.NewVendingHandler(usecase)

	return handler
}
