// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

// RequestEncryptKey
// 请求加密密钥用于敏感数据加密，16位字符，前后端需保持一致
// 安全起见，生产环境运行时请注意修改
var RequestEncryptKey = []byte("f080a463654b2279")

// 配置数据类型
const (
	ConfigTypeString      = "string"
	ConfigTypeInt         = "int"
	ConfigTypeInt8        = "int8"
	ConfigTypeInt16       = "int16"
	ConfigTypeInt32       = "int32"
	ConfigTypeInt64       = "int64"
	ConfigTypeUint        = "uint"
	ConfigTypeUint8       = "uint8"
	ConfigTypeUint16      = "uint16"
	ConfigTypeUint32      = "uint32"
	ConfigTypeUint64      = "uint64"
	ConfigTypeFloat32     = "float32"
	ConfigTypeFloat64     = "float64"
	ConfigTypeBool        = "bool"
	ConfigTypeDate        = "date"
	ConfigTypeDateTime    = "datetime"
	ConfigTypeSliceString = "[]string"
	ConfigTypeSliceInt    = "[]int"
	ConfigTypeSliceInt64  = "[]int64"
)

var ConfigTypes = []string{ConfigTypeString,
	ConfigTypeInt, ConfigTypeInt8, ConfigTypeInt16, ConfigTypeInt32, ConfigTypeInt64,
	ConfigTypeUint, ConfigTypeUint8, ConfigTypeUint16, ConfigTypeUint32, ConfigTypeUint64,
	ConfigTypeFloat32, ConfigTypeFloat64,
	ConfigTypeBool,
	ConfigTypeDate, ConfigTypeDateTime,
	ConfigTypeSliceString, ConfigTypeSliceInt, ConfigTypeSliceInt64,
}

// ConfigMaskDemoField 演示环境下需要隐藏的配置
var ConfigMaskDemoField = map[string]struct{}{
	// 邮箱
	"smtpUser": {}, "smtpPass": {},

	// 云存储
	"uploadUCloudPublicKey": {}, "uploadUCloudPrivateKey": {}, "uploadCosSecretId": {}, "uploadCosSecretKey": {},
	"uploadOssSecretId": {}, "uploadOssSecretKey": {}, "uploadQiNiuAccessKey": {}, "uploadQiNiuSecretKey": {},

	// 地图
	"geoAmapWebKey": {},

	// 短信
	"smsAliYunAccessKeyID": {}, "smsAliYunAccessKeySecret": {}, "smsTencentSecretId": {}, "smsTencentSecretKey": {},

	// 支付
	"payWxPayMchId": {}, "payWxPaySerialNo": {}, "payWxPayAPIv3Key": {}, "payWxPayPrivateKey": {}, "payQQPayMchId": {}, "payQQPayApiKey": {},

	// 微信
	"officialAccountAppSecret": {}, "officialAccountToken": {}, "officialAccountEncodingAESKey": {}, "openPlatformAppSecret": {},
	"openPlatformToken": {}, "openPlatformEncodingAESKey": {},
}

const (
	ConfigGroupTgBot = "tgBot"
)
