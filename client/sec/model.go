package sec

// HomeParam 主页参数模型
type HomeParam struct {
	Code int64 `json:"code"`
	Data struct {
		Backend      string      `json:"backend"`
		CopyRight    string      `json:"copy_right"`
		HasBeianPerm bool        `json:"hasBeianPerm"`
		HasLibrary   bool        `json:"hasLibrary"`
		IsLocalUser  bool        `json:"isLocalUser"`
		Organization string      `json:"organization"`
		PlacardShow  interface{} `json:"placard_show"`
		Protocols    []string    `json:"protocols"`
		SysPlacard   interface{} `json:"sys_placard"`
		Username     string      `json:"username"`
	} `json:"data"`
	Msg string `json:"msg"`
}

// CurrentMember 当前用户返回结构
type CurrentMemberRte struct {
	Attributes interface{} `json:"attributes"`
	Count      interface{} `json:"count"`
	Msg        string      `json:"msg"`
	Obj        struct {
		LastLoginTime           string      `json:"lastLoginTime"`
		MemberAcademicNumber    string      `json:"memberAcademicNumber"`
		MemberCreateTime        int64       `json:"memberCreateTime"`
		MemberID                string      `json:"memberId"`
		MemberIDNumber          string      `json:"memberIdNumber"`
		MemberImage             interface{} `json:"memberImage"`
		MemberMailbox           interface{} `json:"memberMailbox"`
		MemberNickname          string      `json:"memberNickname"`
		MemberOtherBirthday     interface{} `json:"memberOtherBirthday"`
		MemberOtherClass        interface{} `json:"memberOtherClass"`
		MemberOtherDepartment   interface{} `json:"memberOtherDepartment"`
		MemberOtherGrade        interface{} `json:"memberOtherGrade"`
		MemberOtherMajor        interface{} `json:"memberOtherMajor"`
		MemberOtherNation       interface{} `json:"memberOtherNation"`
		MemberOtherNative       interface{} `json:"memberOtherNative"`
		MemberOtherSchoolNumber interface{} `json:"memberOtherSchoolNumber"`
		MemberPhone             interface{} `json:"memberPhone"`
		MemberPwd               interface{} `json:"memberPwd"`
		MemberSex               int64       `json:"memberSex"`
		MemberSign              interface{} `json:"memberSign"`
		MemberState             int64       `json:"memberState"`
		MemberUsername          string      `json:"memberUsername"`
		RoleCodeList            []string    `json:"roleCodeList"`
		RoleList                []struct {
			RoleCode    string      `json:"roleCode"`
			RoleComment interface{} `json:"roleComment"`
			RoleName    string      `json:"roleName"`
			RoleState   interface{} `json:"roleState"`
		} `json:"roleList"`
	} `json:"obj"`
	Success bool `json:"success"`
}

// GetStudentRte 查询学生返回结构
type GetStudentRte struct {
	Attributes interface{} `json:"attributes"`
	Count      interface{} `json:"count"`
	Msg        string      `json:"msg"`
	Obj        struct {
		StudentAdmissionTime    string      `json:"studentAdmissionTime"`
		StudentAdress           interface{} `json:"studentAdress"`
		StudentBirthday         string      `json:"studentBirthday"`
		StudentCategory         string      `json:"studentCategory"`
		StudentClassCode        string      `json:"studentClassCode"`
		StudentClassName        string      `json:"studentClassName"`
		StudentEductionalSystme string      `json:"studentEductionalSystme"`
		StudentFacultiesCode    string      `json:"studentFacultiesCode"`
		StudentFacultiesName    string      `json:"studentFacultiesName"`
		StudentGrade            string      `json:"studentGrade"`
		StudentID               string      `json:"studentId"`
		StudentIDNumber         string      `json:"studentIdNumber"`
		StudentMajor            string      `json:"studentMajor"`
		StudentMajorName        string      `json:"studentMajorName"`
		StudentName             string      `json:"studentName"`
		StudentNation           interface{} `json:"studentNation"`
		StudentPhone            string      `json:"studentPhone"`
		StudentPoliticalStatus  interface{} `json:"studentPoliticalStatus"`
		StudentRegisterState    string      `json:"studentRegisterState"`
		StudentSex              string      `json:"studentSex"`
	} `json:"obj"`
	Success bool `json:"success"`
}

// GetClassmatesDetailRte  获取同班同学信息响应结构

type GetClassmatesDetailRte struct {
	Attributes struct {
		CountMan                   uint64 `json:"countMan"`
		CountWoman                 uint64 `json:"countWoman"`
		InstructorClassClassNumber string `json:"instructorClassClassNumber"`
		StudentSex                 string `json:"studentSex"`
	} `json:"attributes"`
	Count   interface{} `json:"count"`
	Msg     string      `json:"msg"`
	Obj     interface{} `json:"obj"`
	Success bool        `json:"success"`
}

// GetClassmatesRte 获取同班同学列表响应结构
type GetClassmatesRte struct {
	Attributes interface{} `json:"attributes"`
	Count      int64       `json:"count"`
	Msg        string      `json:"msg"`
	Obj        []struct {
		StudentAdmissionTime    string      `json:"studentAdmissionTime"`
		StudentAdress           interface{} `json:"studentAdress"`
		StudentBirthday         string      `json:"studentBirthday"`
		StudentCategory         string      `json:"studentCategory"`
		StudentClassCode        string      `json:"studentClassCode"`
		StudentClassName        string      `json:"studentClassName"`
		StudentEductionalSystme string      `json:"studentEductionalSystme"`
		StudentFacultiesCode    string      `json:"studentFacultiesCode"`
		StudentFacultiesName    string      `json:"studentFacultiesName"`
		StudentGrade            string      `json:"studentGrade"`
		StudentID               string      `json:"studentId"`
		StudentIDNumber         string      `json:"studentIdNumber"`
		StudentMajor            string      `json:"studentMajor"`
		StudentMajorName        string      `json:"studentMajorName"`
		StudentName             string      `json:"studentName"`
		StudentNation           interface{} `json:"studentNation"`
		StudentPhone            string      `json:"studentPhone"`
		StudentPoliticalStatus  interface{} `json:"studentPoliticalStatus"`
		StudentRegisterState    string      `json:"studentRegisterState"`
		StudentSex              string      `json:"studentSex"`
	} `json:"obj"`
	Success bool `json:"success"`
}

// GetOneCardBalanceRte 获取一卡通余额返回结构
type GetOneCardBalanceRte struct {
	Attributes interface{} `json:"attributes"`
	Count      interface{} `json:"count"`
	Msg        string      `json:"msg"`
	Obj        struct {
		Balance        string `json:"balance"`
		LastMonthMoney string `json:"lastMonthMoney"`
		ThisMonthMoney string `json:"thisMonthMoney"`
	} `json:"obj"`
	Success bool `json:"success"`
}

// GetOneCardChargeRecordsRte 获取一卡通充值记录返回结构
type GetOneCardChargeRecordsRte struct {
	Attributes interface{} `json:"attributes"`
	Count      int64       `json:"count"`
	Msg        string      `json:"msg"`
	Obj        []struct {
		GeneraCardRechargeRecordID                 int64       `json:"generaCardRechargeRecordId"`
		GeneraCardRechargeRecordNumber             string      `json:"generaCardRechargeRecordNumber"`
		GeneraCardRechargeRecordTransactionAdress  string      `json:"generaCardRechargeRecordTransactionAdress"`
		GeneraCardRechargeRecordTransactionBalance string      `json:"generaCardRechargeRecordTransactionBalance"`
		GeneraCardRechargeRecordTransactionMoney   string      `json:"generaCardRechargeRecordTransactionMoney"`
		GeneraCardRechargeRecordTransactionTime    string      `json:"generaCardRechargeRecordTransactionTime"`
		GeneraCardRechargeRecordTransactionType    string      `json:"generaCardRechargeRecordTransactionType"`
		GeneraCardRechargeRecordWalletType         interface{} `json:"generaCardRechargeRecordWalletType"`
	} `json:"obj"`
	Success bool `json:"success"`
}

// GetOneCardConsumeRecordsRte 获取一卡通消费记录返回结构
type GetOneCardConsumeRecordsRte struct {
	Attributes interface{} `json:"attributes"`
	Count      int64       `json:"count"`
	Msg        string      `json:"msg"`
	Obj        []struct {
		GeneraCardConsumeRecordID                 int64       `json:"generaCardConsumeRecordId"`
		GeneraCardConsumeRecordNumber             string      `json:"generaCardConsumeRecordNumber"`
		GeneraCardConsumeRecordTransactionAdress  string      `json:"generaCardConsumeRecordTransactionAdress"`
		GeneraCardConsumeRecordTransactionBalance string      `json:"generaCardConsumeRecordTransactionBalance"`
		GeneraCardConsumeRecordTransactionMoney   string      `json:"generaCardConsumeRecordTransactionMoney"`
		GeneraCardConsumeRecordTransactionTime    string      `json:"generaCardConsumeRecordTransactionTime"`
		GeneraCardConsumeRecordTransactionType    string      `json:"generaCardConsumeRecordTransactionType"`
		GeneraCardConsumeRecordWalletType         interface{} `json:"generaCardConsumeRecordWalletType"`
	} `json:"obj"`
	Success bool `json:"success"`
}

// GetExamArrangementsRte 获取考试安排返回结构
type GetExamArrangementsRte struct {
	Attributes interface{} `json:"attributes"`
	Count      interface{} `json:"count"`
	Msg        string      `json:"msg"`
	Obj        []struct {
		ExaminationAdressCode  interface{} `json:"examinationAdressCode"`
		ExaminationAdressName  string      `json:"examinationAdressName"`
		ExaminationCourseCode  string      `json:"examinationCourseCode"`
		ExaminationCourseName  string      `json:"examinationCourseName"`
		ExaminationEndTime     string      `json:"examinationEndTime"`
		ExaminationSchoolYear  string      `json:"examinationSchoolYear"`
		ExaminationSeat        string      `json:"examinationSeat"`
		ExaminationStartTime   string      `json:"examinationStartTime"`
		ExaminationStudentID   string      `json:"examinationStudentId"`
		ExaminationStudentName string      `json:"examinationStudentName"`
		ExaminationTerm        string      `json:"examinationTerm"`
		ExaminationTime        string      `json:"examinationTime"`
	} `json:"obj"`
	Success bool `json:"success"`
}

// GetWeekCoursesRte 获取考试安排返回结构
type GetWeekCoursesRte struct {
	Attributes struct {
		CourseListTopTwo []struct {
			CourseAdressCode            string      `json:"courseAdressCode"`
			CourseCategoryName          interface{} `json:"courseCategoryName"`
			CourseClassCode             interface{} `json:"courseClassCode"`
			CourseClassName             interface{} `json:"courseClassName"`
			CourseClassRoomCode         interface{} `json:"courseClassRoomCode"`
			CourseClassRoomName         interface{} `json:"courseClassRoomName"`
			CourseCode                  string      `json:"courseCode"`
			CourseCredit                interface{} `json:"courseCredit"`
			CourseDate                  string      `json:"courseDate"`
			CourseDepartmentCourse      interface{} `json:"courseDepartmentCourse"`
			CourseExaminationMethodCode interface{} `json:"courseExaminationMethodCode"`
			CourseHours                 interface{} `json:"courseHours"`
			CourseID                    interface{} `json:"courseId"`
			CourseName                  string      `json:"courseName"`
			CoursePlan                  interface{} `json:"coursePlan"`
			CourseSchoolYear            string      `json:"courseSchoolYear"`
			CourseSection               string      `json:"courseSection"`
			CourseSectionWeek           interface{} `json:"courseSectionWeek"`
			CourseSingleDoubleWeek      string      `json:"courseSingleDoubleWeek"`
			CourseStudentID             interface{} `json:"courseStudentId"`
			CourseStudentName           interface{} `json:"courseStudentName"`
			CourseSubjectCourseNumber   interface{} `json:"courseSubjectCourseNumber"`
			CourseTeacherName           string      `json:"courseTeacherName"`
			CourseTeacherNumber         interface{} `json:"courseTeacherNumber"`
			CourseTeachingNumber        interface{} `json:"courseTeachingNumber"`
			CourseTerm                  string      `json:"courseTerm"`
			CourseTotolHours            interface{} `json:"courseTotolHours"`
			CourseWeek                  interface{} `json:"courseWeek"`
			CourseWeekNumber            interface{} `json:"courseWeekNumber"`
			CourseWeekly                string      `json:"courseWeekly"`
			TeacherName                 string      `json:"teacherName"`
		} `json:"courseListTopTwo"`
		CourseTime string `json:"courseTime"`
	} `json:"attributes"`
	Count interface{} `json:"count"`
	Msg   string      `json:"msg"`
	Obj   []struct {
		CourseAdressCode            string      `json:"courseAdressCode"`
		CourseCategoryName          interface{} `json:"courseCategoryName"`
		CourseClassCode             interface{} `json:"courseClassCode"`
		CourseClassName             interface{} `json:"courseClassName"`
		CourseClassRoomCode         interface{} `json:"courseClassRoomCode"`
		CourseClassRoomName         interface{} `json:"courseClassRoomName"`
		CourseCode                  string      `json:"courseCode"`
		CourseCredit                interface{} `json:"courseCredit"`
		CourseDate                  string      `json:"courseDate"`
		CourseDepartmentCourse      interface{} `json:"courseDepartmentCourse"`
		CourseExaminationMethodCode interface{} `json:"courseExaminationMethodCode"`
		CourseHours                 interface{} `json:"courseHours"`
		CourseID                    interface{} `json:"courseId"`
		CourseName                  string      `json:"courseName"`
		CoursePlan                  interface{} `json:"coursePlan"`
		CourseSchoolYear            string      `json:"courseSchoolYear"`
		CourseSection               string      `json:"courseSection"`
		CourseSectionWeek           interface{} `json:"courseSectionWeek"`
		CourseSingleDoubleWeek      string      `json:"courseSingleDoubleWeek"`
		CourseStudentID             interface{} `json:"courseStudentId"`
		CourseStudentName           interface{} `json:"courseStudentName"`
		CourseSubjectCourseNumber   interface{} `json:"courseSubjectCourseNumber"`
		CourseTeacherName           string      `json:"courseTeacherName"`
		CourseTeacherNumber         interface{} `json:"courseTeacherNumber"`
		CourseTeachingNumber        interface{} `json:"courseTeachingNumber"`
		CourseTerm                  string      `json:"courseTerm"`
		CourseTotolHours            interface{} `json:"courseTotolHours"`
		CourseWeek                  interface{} `json:"courseWeek"`
		CourseWeekNumber            interface{} `json:"courseWeekNumber"`
		CourseWeekly                string      `json:"courseWeekly"`
		TeacherName                 string      `json:"teacherName"`
	} `json:"obj"`
	Success bool `json:"success"`
}
