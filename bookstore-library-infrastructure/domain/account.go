package domain

type Account struct {
	Username  string `thrift:"username,1" frugal:"1,default,string" protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" validate:"required"`
	Password  string `thrift:"password,2" frugal:"2,default,string" protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name      string `thrift:"name,3" frugal:"3,default,string" protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" validate:"required"`
	Avatar    string `thrift:"avatar,4" frugal:"4,default,string" protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Telephone string `thrift:"telephone,5" frugal:"5,default,string" protobuf:"bytes,5,opt,name=telephone,proto3" json:"telephone,omitempty" validate:"required,matches(1[0-9]{10})"`
	Email     string `thrift:"email,6" frugal:"6,default,string" protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty" validate:"email"`
	Location  string `thrift:"location,7" frugal:"7,default,string" protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
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
