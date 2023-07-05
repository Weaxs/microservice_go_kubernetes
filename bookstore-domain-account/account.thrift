namespace go account

struct Account {
    1: string username
    2: string password
    3: string name
    4: string avatar
    5: string telephone
    6: string email
    7: string location
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