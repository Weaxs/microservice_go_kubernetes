namespace go account

struct Account {
    1: i64 id
    2: string username
    3: string password
    4: string name
    5: string avatar
    6: string telephone
    7: string email
    8: string location
}

struct GetAccountRequest {
    1: string username
}

struct GetAccountResponse {
    1: Account account
}

struct ChangeAccountRequest {
    1: Account account
}

struct ChangeAccountResponse {
}


service AccountApi {
    GetAccountResponse getAccount(1: GetAccountRequest usernmae)
    ChangeAccountResponse createAccount(1: ChangeAccountRequest account)
    ChangeAccountResponse updateAccount(1: ChangeAccountRequest account)
}