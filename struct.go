package main

type userInfo struct {
	IsDeleted bool
	Name      string
	ImeiCode  string
	UseLeft   string
	Distance  string
	Email     string
	LastTime  string
}

const (
	smtpServer  = "smtp.qq.com:25"
	userName    = ""
	Vcode       = ""
	contentType = "Content-Type: text/html; charset=UTF-8"
	mailFrom    = ""
	infoText    = "姓名: (name)  剩余次数: (time)  跑步距离: (runDistance)"
	infoSubject = "提醒:今日份(date)阳光长跑已自动完成"
	infoBody    = `
    <html>
    <body>
    <h3>
	 (text)
    </h3>
    </body>
    </html>
    `
	apiRoot    = "http://client3.aipao.me/api"
	appVersion = "2.15"
	UserAgent  = "Dalvik/2.1.0 (Linux; U; Android 9.0.0; COL-AL10 Build/HUAWEICOL-LL10)"
	longtitude = "31.93178" // 自行更改
	latitute   = "118.8865121" // 自行更改
)

type codeInfo struct {
	Code     string
	ImeiCode string
	Data     struct {
		Name       string
		Time       string
		CreateTime string
	}
}

type Data struct {
	Name       string
	Time       string
	CreateTime string
}

type returnInfo struct {
	Success bool
	Data    struct {
		Token    string
		UserId   int
		IMEICode string
		// AndroidVer float32
		// AppleVer   float32
		// WinVer     float32
	}
}

type returnRun struct {
	Success bool
	Data    struct {
		StartTime string
		RunId     string
		// FUserId   int
		// FieldId   int
		Routes string
		// LifeValue int
		// Powers    int
		// LenValue  float32
		Point struct {
			// PointNo string
			// Lat     float32
			// Lng     float32
			// Minor   int
		}
		// FiledName string
		// Area      string
		// SenseType string
		// ImgUrl    string
		// Major     int
	}
}

type returnEnd struct {
	Success bool
	Data    string
}
