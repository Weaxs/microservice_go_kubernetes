namespace go warehouse

enum DeliveredStatus {
  DECREASE
  INCREASE
  FROZEN
  THAWED
}

struct Advertisement {
    1: string image
    2: i64 productId
}

struct Specification {
    1: string item
    2: string value
    3: i64 productId
}

struct Product {
    1:  string title
    2: double price
    3: double rate
    4: string description
    5: string cover
    6: string detail
    7: list<Specification> specifications
}

struct Stockpile {
    1: i64 amount
    2: i64 frozen
    3: Product product
}

struct Empty {
}

struct GetAllAdvertisementsResponse {
    1: list<Advertisement> advertisements
}

struct GetAllProductResponse {
    1: list<Product> products
}

struct GetProductRequest {
    1: i64 id
}

struct RemoveProductRequest {
    1: i64 id
}

struct GetProductResponse {
    1: Product product
}

struct ChangeProductRequest {
    1: Product product
}

struct UpdateStockpileRequest {
    1: i64 productId
    2: i64 amount
}

struct QueryStockpileRequest {
    1: i64 productId
}

struct QueryStockpileResponse {
    1: Stockpile stockpile
}

struct SetDeliveredStatusRequest {
    1: i64 productId
    2: i64 amount
    3: DeliveredStatus status
}

service WarehouseApi {
    // 获取所有的广告信息
      GetAllAdvertisementsResponse GetAllAdvertisements(Empty request)
    
      // 获取仓库中所有的货物信息
      GetAllProductResponse GetAllProducts(Empty request)
      // 获取仓库中指定的货物信息
      GetProductResponse GetProduct(GetProductRequest request)
      // 创建新的产品
      Empty CreateProduct(ChangeProductRequest request)
      // 更新产品信息
      Empty UpdateProduc(ChangeProductRequest request)
      // 移出产品信息
      Empty RemoveProduct(RemoveProductRequest request)
    
      // 将指定的产品库存调整为指定数额
      Empty UpdateStockpile (UpdateStockpileRequest request)
      QueryStockpileResponse QueryStockpile(QueryStockpileRequest request)
      Empty SetDeliveredStatus(SetDeliveredStatusRequest request)

}
