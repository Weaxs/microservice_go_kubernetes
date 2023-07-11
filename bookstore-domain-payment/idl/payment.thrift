namespace go payment

include "warehouse.thrift"

typedef string Timestamp

enum PaymentState {
     WAITING
     CANCEL
     PAYED
     TIMEOUT
     NOT_SUPPORT
}

struct Item {
  i64 amount
  i64 productId
}

struct Purchase {
  bool delivery
  string name
  string telephone
  string location
}

struct Payment {
  Timestamp createTime
  string payId
  double totalPrice
  i64 expires
  string paymentLink
  PaymentState payState
}

// 支付结算单
struct Settlement {
  list<Item> items
  Purchase purchase
  map<i64, warehouse.Product> productMap
}

struct ExecuteSettlementRequest {
  Settlement settlement
}
struct ExecuteSettlementResponse {
  Payment payment
}

struct UpdatePaymentStateRequest {
  string payId
  PaymentState state
}

struct UpdatePaymentStateAlias {
  string payId
  i64 accountId
  PaymentState state
}

service PaymentApi {
  ExecuteSettlementResponse executeSettlement(ExecuteSettlementRequest request)
  warehouse.Empty updatePaymentState(UpdatePaymentStateRequest request)
  warehouse.Empty updatePaymentStateAlias(UpdatePaymentStateAlias request)
}