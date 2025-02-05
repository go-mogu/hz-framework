package util

import (
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"gorm.io/gorm/schema"
	"io"
	"net"
	"net/url"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

// InAnySlice 判断某个字符串是否在字符串切片中
func InAnySlice[T string | int | int64 | float32 | float64](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// GenerateBaseSnowId 生成雪花算法ID
func GenerateBaseSnowId(num int, n *snowflake.Node) string {
	if n == nil {
		node, err := snowflake.NewNode(1)
		if err != nil {
			return ""
		}
		n = node
	}
	id := n.Generate()
	switch num {
	case 2:
		return id.Base2()
	case 32:
		return id.Base32()
	case 36:
		return id.Base36()
	case 58:
		return id.Base58()
	case 64:
		return id.Base64()
	default:
		return id.Base32()
	}
}

// GenerateUuid 生成随机字符串
func GenerateUuid(size int) string {
	return SubStr(uuid.New().String(), 0, size)
}

func SubStr(str string, start int, length ...int) (substr string) {
	strLength := len(str)
	// Simple border checks.
	if start < 0 {
		start = 0
	}
	if start >= strLength {
		start = strLength
	}
	end := strLength
	if len(length) > 0 {
		end = start + length[0]
		if end < start {
			end = strLength
		}
	}
	if end > strLength {
		end = strLength
	}
	return str[start:end]
}

// GeneratePasswordHash 生成密码hash值
func GeneratePasswordHash(password string, salt string) string {
	md5 := md5.New()
	io.WriteString(md5, password)
	str := fmt.Sprintf("%x", md5.Sum(nil))
	s := sha256.New()
	io.WriteString(s, password+salt)
	str = fmt.Sprintf("%x", s.Sum(nil))
	return str
}

// FormatToString 格式化转化成string
func FormatToString(originStr interface{}) string {
	str := ""
	switch originStr.(type) {
	case float64:
		str = strconv.FormatFloat(originStr.(float64), 'f', 10, 64)
	case float32:
		str = strconv.FormatFloat(originStr.(float64), 'f', 10, 32)
	case nil:
		str = ""
	case int, int32, int64:
		str = strconv.FormatInt(originStr.(int64), 10)
	default:
		str = originStr.(string)
	}
	return str
}

// IsPathExist 判断所给路径文件/文件夹是否存在
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// MakeMultiDir 调用os.MkdirAll递归创建文件夹
func MakeMultiDir(filePath string) error {
	if !IsPathExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// MakeFileOrPath 创建文件/文件夹
func MakeFileOrPath(path string) bool {
	create, err := os.Create(path)
	defer create.Close()
	if err != nil {
		return false
	}
	return true
}

// String2Int 将数组的string转int
func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))
	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}
	return res
}

// GetStructColumnName 获取结构体中的字段名称 _type: 1: 获取tag字段值 2：获取结构体字段值
func GetStructColumnName(s interface{}, _type int) ([]string, error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return []string{}, fmt.Errorf("interface is not a struct")
	}
	t := v.Type()
	var fields []string
	for i := 0; i < v.NumField(); i++ {
		var field string
		if _type == 1 {
			field = t.Field(i).Tag.Get("json")
			if field == "" {
				tagSetting := schema.ParseTagSetting(t.Field(i).Tag.Get("gorm"), ";")
				field = tagSetting["COLUMN"]
			}
		} else {
			field = t.Field(i).Name
		}
		fields = append(fields, field)
	}
	return fields, nil
}

// GetProjectModuleName 获取当前项目的module名称
func GetProjectModuleName() string {
	cmd := exec.Command("go", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(output), "\n")
}

func GetIpAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// 192.168.1.20:61085
	ip := strings.Split(localAddr.String(), ":")[0]

	return ip
}

func ParseURL(str string, component int) (map[string]string, error) {
	u, err := url.Parse(str)
	if err != nil {
		err = errors.New(fmt.Sprintf(`url.Parse failed for URL "%s"`, str))
		return nil, err
	}
	if component == -1 {
		component = 1 | 2 | 4 | 8 | 16 | 32 | 64 | 128
	}
	var components = make(map[string]string)
	if (component & 1) == 1 {
		components["scheme"] = u.Scheme
	}
	if (component & 2) == 2 {
		components["host"] = u.Hostname()
	}
	if (component & 4) == 4 {
		components["port"] = u.Port()
	}
	if (component & 8) == 8 {
		components["user"] = u.User.Username()
	}
	if (component & 16) == 16 {
		components["pass"], _ = u.User.Password()
	}
	if (component & 32) == 32 {
		components["path"] = u.Path
	}
	if (component & 64) == 64 {
		components["query"] = u.RawQuery
	}
	if (component & 128) == 128 {
		components["fragment"] = u.Fragment
	}
	return components, nil
}
