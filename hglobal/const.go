package hglobal

import (
	"time"
)

// 所有常量文件读取位置
const (
	DEFAULT_CONFIG_FILE_PATH = "\\config\\config.yaml"
	TMPDIR                   = "./tmp_upload"
	ATOKEN_EFFECTIVE_TIME    = time.Hour * 2

	RTOKEN_EFFECTIVE_TIME = time.Hour * 24 * 30
	AUTH_ENUMS_ATOKEN     = "atoken"
	AUTH_ENUMS_RTOKEN     = "rtoken"
	DEFAULT_NODE_ID       = 1
	TOKEN_USER_ID         = "user_id"
	TOKEN_ROLE            = "role"
	TOKEN_USER_NAME       = "user_name"
	TOKEN_EMAIL           = "email"

	ROLE_NOT_LOGIN = -1
	ROLE_GUEST     = 0
	ROLE_USER      = 1
	ROLE_ADMIN     = 2
)

// 分布式锁
const ()
