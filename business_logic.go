package abci-demo

// 1检查-初始化平台
func (app *ABCIDemoApplication) checkInitPlatform (req *Request) error {
	_, _, exists := app.state.Get(KeyPlatform())
	if exists {
		return ErrPlatformIsInit
	}

	initPlatform := req.GetInitPlatform()
	if initPlatform == nil {
		return ErrGetValueIsNull
	}

	if initPlatform.UserName == "" {
		return ErrUserNameIsNull
	}
	if len(initPlatform.UserPublicKey) != 32 {
		return ErrUserPublicKey
	}

	return nil
}

// 1初始化平台
func (app *ABCIDemoApplication) initPlatform(req *Request) (*Response, error) {
	err := app.checkInitPlatform(req)
	if err != nil {
		return nil, err
	}

	initPlatform := req.GetInitPlatform()
	if initPlatform == nil {
		return nil, ErrGetValueIsNull
	}

	user := &User{}
	user.Username = initPlatform.UserName

	save, err := MarshalMessage(user)
	if err != nil {
		return nil, err
	}
	app.state.Set(KeyUser(user.Username), save)
	app.state.Set(KeyPlatform(), nil)

	return &Response{Value: &Response_InitPlatform{&ResponseInitPlatform{InstructionId: req.InstructionId}}}, nil
}

// 2检查--创建用户
func (app *ABCIDemoApplication) checkUserInfoOnChain(req *Request) error {
	addUser := req.GetUserInfoOnChain()
	if addUser == nil {
		return ErrGetValueIsNull
	}

	// 用户信息都是必填字段，不可为空
	if addUser.Username == "" {
		return ErrUserNameIsNull
	}
	_, _, exists := app.state.Get(KeyUser(addUser.Username))
	if exists {
		return ErrUserExist
	}

	return nil
}

// 2创建用户
func (app *ABCIDemoApplication) userInfoOnChain(req *Request) (*Response, error) {
	addUser := req.GetUserInfoOnChain()
	if addUser == nil {
		return nil, ErrGetValueIsNull
	}
	err := app.checkUserInfoOnChain(req)
	if err != nil {
		return nil, err
	}
	user := &User{}
	user.Username = addUser.Username
	user.Age = addUser.Age
	user.Workstation = addUser.Workstation
	user.Educate = addUser.Educate
	save, err := MarshalMessage(user)
	if err != nil {
		return nil, err
	}
	app.state.Set(KeyUser(user.Username), save)
	return &Response{Value: &Response_UserInfoOnChain{&ResponseUserInfoOnChain{InstructionId: req.InstructionId}}}, nil
}





func (app *ABCIDemoApplication) saveReceipt(instructionId int64, receipt *Receipt) error {
	save, err := MarshalMessage(receipt)
	if err != nil {
		return err
	}
	app.state.Set(KeyReq(instructionId), save)
	return err
}

func (app *ABCIDemoApplication) checkInstructionId(instructionId int64) error {
	_, _, exists := app.state.Get(KeyReq(instructionId))
	if !exists {
		return nil
	}
	return ErrDupInstructionId
}


