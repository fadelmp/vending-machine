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

	mapper := mapper.NewVendingMapper()
	repo := repository.NewVendingRepository(db, redis)

	comparator := comparator.NewVendingComparator(repo)
	usecase := usecase.NewVendingUsecase(repo, mapper, comparator)
	handler := handler.NewVendingHandler(usecase)

	return handler
}
