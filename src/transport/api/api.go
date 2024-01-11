package api

func (al APIStandardLauncher) APISetUp() {
	am := al.SetUp()
	//유저 관련 api
	am.POST("/user/signup", al.SignUp)
	am.POST("/user/signin", al.SignIn)
	am.POST("/user/emailchk", al.EmailCheck)
	am.GET("/user/auto", al.AutoLogin)
	am.GET("/user/refresh", al.Refresh)
	am.GET("/user/mypage", al.Mypage)
	am.POST("/user/updatemypage", al.UpdateMypage)
	am.POST("/user/changepassword", al.ChangePassword)

	//게시글 관련 api
}
