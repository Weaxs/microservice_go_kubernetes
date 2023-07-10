package main

import (
	"context"
	"errors"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/asaskevich/govalidator"
)

var (
	advertiseService = &AdvertiseService{
		repo: &AdvertisementRepo{},
	}
	productService = &ProductService{
		repo:              &ProductRepo{},
		specificationRepo: &SpecificationRepo{},
	}
	stockpileService = &StockpileService{
		repo:           &StockpileRepo{},
		productService: productService,
	}
)

type AdvertiseService struct {
	repo *AdvertisementRepo
}

type ProductService struct {
	repo              *ProductRepo
	specificationRepo *SpecificationRepo
}

type StockpileService struct {
	repo           *StockpileRepo
	productService *ProductService
}

func (s *AdvertiseService) GetAll(ctx context.Context) (advertises []*domain.Advertisement, err error) {
	pos, err := s.repo.FindAll(ctx)
	if err != nil {
		return
	}
	for _, po := range pos {
		advertises = append(advertises, &domain.Advertisement{
			Image:     po.Image,
			ProductId: po.ProductId,
		})
	}
	return
}

func (s *ProductService) GetAll(ctx context.Context) (products []*domain.Product, err error) {
	pos, err := s.repo.FindAll(ctx)
	if err != nil {
		return
	}
	for _, po := range pos {
		specPos, _ := s.specificationRepo.FindByProductId(ctx, po.Id)
		var specifications []*domain.Specification
		if len(specPos) > 0 {
			for _, specPo := range specPos {
				specifications = append(specifications, &domain.Specification{
					Item:      specPo.Item,
					Value:     specPo.Value,
					ProductId: specPo.ProductId,
				})
			}
		}

		products = append(products, &domain.Product{
			Id:             po.Id,
			Title:          po.Title,
			Price:          po.Price,
			Rate:           po.Rate,
			Cover:          po.Cover,
			Description:    po.Description,
			Detail:         po.Detail,
			Specifications: specifications,
		})
	}
	return
}

func (s *ProductService) GetById(ctx context.Context, id int64) (product *domain.Product, err error) {
	po, err := s.repo.FindById(ctx, id)
	if err != nil {
		return
	}
	if po == nil {
		err = errors.New("product does not exist")
		return
	}

	specPos, _ := s.specificationRepo.FindByProductId(ctx, po.Id)
	var specifications []*domain.Specification
	if len(specPos) > 0 {
		for _, specPo := range specPos {
			specifications = append(specifications, &domain.Specification{
				Item:      specPo.Item,
				Value:     specPo.Value,
				ProductId: specPo.ProductId,
			})
		}
	}
	product = &domain.Product{}
	product.Id = po.Id
	product.Title = po.Title
	product.Price = po.Price
	product.Rate = po.Rate
	product.Description = po.Description
	product.Cover = po.Cover
	product.Detail = po.Detail
	product.Specifications = specifications
	return
}

func (s *ProductService) Create(ctx context.Context, product *domain.Product) (err error) {
	if _, err = govalidator.ValidateStruct(product); err != nil {
		return err
	}

	productPo := &ProductPo{
		Title:       product.Title,
		Price:       product.Price,
		Rate:        product.Rate,
		Description: product.Description,
		Cover:       product.Cover,
		Detail:      product.Detail,
	}
	err = s.repo.Create(ctx, productPo)
	if err != nil {
		return err
	}

	if len(product.Specifications) > 0 {
		var pos []*SpecificationPo
		for _, specification := range product.Specifications {
			pos = append(pos, &SpecificationPo{
				Item:      specification.Item,
				Value:     specification.Value,
				ProductId: productPo.Id,
			})
		}
		err = s.specificationRepo.Create(ctx, pos)
		if err != nil {
			return
		}
	}
	return
}

func (s *ProductService) Update(ctx context.Context, product *domain.Product) (err error) {
	if _, err = govalidator.ValidateStruct(product); err != nil {
		return err
	}

	po := &ProductPo{
		Id:          product.Id,
		Title:       product.Title,
		Price:       product.Price,
		Rate:        product.Rate,
		Description: product.Description,
		Cover:       product.Cover,
		Detail:      product.Detail,
	}
	err = s.repo.Update(ctx, po)
	if err != nil {
		return
	}

	return
}

func (s *ProductService) RemoveById(ctx context.Context, id int64) (err error) {
	err = s.repo.DelById(ctx, id)
	if err != nil {
		return
	}
	err = s.specificationRepo.DelByProductId(ctx, id)
	if err != nil {
		return
	}
	return
}

func (s *StockpileService) QueryByProduct(ctx context.Context, productId int64) (stockpile *domain.Stockpile, err error) {
	po, err := s.repo.FindByProductId(ctx, productId)
	if err != nil {
		return
	}

	product, _ := s.productService.GetById(ctx, productId)

	stockpile = &domain.Stockpile{}
	stockpile.Amount = po.Amount
	stockpile.Frozen = po.Frozen
	if product != nil {
		stockpile.Product = product
	}
	return
}

func (s *StockpileService) Update(ctx context.Context, productId, amount int64) (err error) {
	stockpile, err := s.QueryByProduct(ctx, productId)
	if err != nil {
		return
	}
	po := &StockpilePo{
		ProductId: productId,
		Amount:    amount,
		Frozen:    stockpile.Frozen,
	}
	err = s.repo.Update(ctx, po)
	if err != nil {
		return err
	}
	return
}

func (s *StockpileService) SetDeliveredStatus(ctx context.Context, productId, amount int64, status domain.DeliveredStatus) (err error) {
	stockpile, err := s.QueryByProduct(ctx, productId)
	if err != nil {
		return
	}
	if status == domain.DeliveredStatus_DECREASE {
		stockpile.Frozen -= amount

	} else if status == domain.DeliveredStatus_INCREASE {
		stockpile.Frozen += amount
	} else if status == domain.DeliveredStatus_FROZEN {
		stockpile.Amount -= amount
		stockpile.Frozen += amount
	} else if status == domain.DeliveredStatus_THAWED {
		stockpile.Amount += amount
		stockpile.Frozen -= amount
	}
	po := &StockpilePo{
		ProductId: productId,
		Amount:    stockpile.Amount,
		Frozen:    stockpile.Frozen,
	}
	err = s.repo.Update(ctx, po)
	if err != nil {
		return err
	}
	return
}
