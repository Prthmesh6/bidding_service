package main

import (
	"context"
	"fmt"

	"github.com/Prthmesh6/bidding_service/models"
	"github.com/Prthmesh6/bidding_service/repository"
	"github.com/Prthmesh6/bidding_service/service"
)

func main() {

	buyerRepo, sellerRepo := repository.CreateBuyerRepo(), repository.CreateSellerRepo()
	sellerSvc := service.CreateSellerMgr(sellerRepo)

	buyerSvc := service.CreatebuyerMgr(buyerRepo)

	sellerSvc.Add(context.Background(), models.User{
		Id: 1, Name: "Prathmesh",
	})

	sellerSvc.Add(context.Background(), models.User{
		Id: 2, Name: "Patil",
	})

	sellerSvc.Add(context.Background(), models.User{
		Id: 3, Name: "Suraj",
	})

	sellerSvc.Add(context.Background(), models.User{
		Id: 4, Name: "Kadam",
	})

	buyerSvc.Add(context.Background(), models.User{
		Id: 1, Name: "JungleeGames",
	})

	buyerSvc.Add(context.Background(), models.User{
		Id: 2, Name: "FlipKart",
	})

	buyerSvc.Add(context.Background(), models.User{
		Id: 3, Name: "MakeMyTrip",
	})

	buyerSvc.Add(context.Background(), models.User{
		Id: 4, Name: "Google",
	})

	fmt.Println(buyerSvc.GetUsers(context.Background()))

	fmt.Println(sellerSvc.GetUsers(context.Background()))

}
