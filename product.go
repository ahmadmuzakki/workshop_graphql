package main

import(
	"github.com/graphql-go/graphql"
)

var ProductObject = graphql.NewObject(graphql.ObjectConfig{
	Name:"Product",
	Fields:graphql.Fields{
		"product_id":&graphql.Field{
			Name:"ProductID",
			Type:graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := p.Source.(ProductData)
				return product.ProductID,nil
			},
		},
		"product_name":&graphql.Field{
			Name:"ProductName",
			Type:graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := p.Source.(ProductData)
				return product.ProductName,nil
			},
		},
		"product_price":&graphql.Field{
			Name:"ProductPrice",
			Type:graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := p.Source.(ProductData)
				return product.ProductPrice,nil
			},
		},
		"shop":&graphql.Field{
			Type:ShopObject,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ShopData{
					ShopID:123,
					ShopName:"tokopedia",
				},nil
			},
		},
	},
})

// this struct is only to hold the product data
type ProductData struct {
	ProductID int
	ProductName string
	ProductPrice float64
}

func getProduct()ProductData{
	// this is where you should query from database
	return ProductData{
		ProductID:123,
		ProductName:"sepatu",
		ProductPrice:1000,
	}
}

func getProductByID(productID int)ProductData{
	// this is where you should query from database
	return ProductData{
		ProductID:productID,
		ProductName:"sepatu",
		ProductPrice:1000,
	}
}

func getProductList()[]ProductData{
	return []ProductData{
		{
			ProductID:123,
			ProductName:"sepatu",
			ProductPrice:1000,
		},{
			ProductID:222,
			ProductName:"baju",
			ProductPrice:5000,
		},
	}
}