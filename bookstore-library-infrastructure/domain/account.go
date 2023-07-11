package domain

type Account struct {
	Id        int64  `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Username  string `thrift:"username,2" frugal:"2,default,string" json:"username"`
	Password  string `thrift:"password,3" frugal:"3,default,string" json:"password"`
	Name      string `thrift:"name,4" frugal:"4,default,string" json:"name"`
	Avatar    string `thrift:"avatar,5" frugal:"5,default,string" json:"avatar"`
	Telephone string `thrift:"telephone,6" frugal:"6,default,string" json:"telephone"`
	Email     string `thrift:"email,7" frugal:"7,default,string" json:"email"`
	Location  string `thrift:"location,8" frugal:"8,default,string" json:"location"`
}

type GetAccountRequest struct {
	Username string `thrift:"username,1" frugal:"1,default,string" json:"username"`
}

type GetAccountResponse struct {
	Account *Account `thrift:"account,1" frugal:"1,default,Account" json:"account"`
}

type ChangeAccountRequest struct {
	Account *Account `thrift:"account,1" frugal:"1,default,Account" json:"account"`
}

type ChangeAccountResponse struct {
}

type GetAccountArgs struct {
	Username *GetAccountRequest `thrift:"usernmae,1" frugal:"1,default,GetAccountRequest" json:"username"`
}
type GetAccountResult struct {
	Success *GetAccountResponse `thrift:"success,0,optional" frugal:"0,optional,GetAccountResponse" json:"success,omitempty"`
}
type ChangeAccountArgs struct {
	Account *ChangeAccountRequest `thrift:"account,1" frugal:"1,default,ChangeAccountRequest" json:"account"`
}
type ChangeAccountResult struct {
	Success *ChangeAccountResponse `thrift:"success,0,optional" frugal:"0,optional,ChangeAccountResponse" json:"success,omitempty"`
}
