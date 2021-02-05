package dao

import (
	"chatroom/models"
	"chatroom/util"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/prometheus/common/log"
)

const (
	getUserByTelSql     = "SELECT USER_ID FROM USER WHERE TEL = ?"
	getUserByEmailSql   = "SELECT USER_ID FROM USER WHERE EMAIL = ?"
	getUserByAccountSql = "SELECT USER_ID FROM USER WHERE ACCOUNT = ?"
	getUserByUserIdSql  = "SELECT id, name, tel, email, account, pwd, user_id, nick_name, salt, status, is_delete FROM USER WHERE USER_ID = ?"

	checkUserRegSql = "SELECT CASE WHEN account = ? THEN '账号已存在' WHEN tel = ? THEN '手机号已存在' WHEN email = ? THEN '邮箱已存在' WHEN nick_name = ? THEN '昵称已存在' ELSE '' END AS errMsg " +
		"FROM `user` u " +
		"WHERE (account = ? and account != '') OR (tel = ? and tel != '')  OR (email = ? and email != '')  OR (nick_name = ? and nick_name != '') "

	createUserSql = "INSERT INTO chatroom.`user` (tel, email, account, pwd, user_id , nick_name, salt) VALUES (?, ?, ?, ?, ? , ?, ?)"

	origin_ras_pub_key = "-----BEGIN PUBLIC KEY-----\n" +
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDN37iIMNod9I3X/7Vfh4YSwNHm\n" +
		"aFAr0Qsc5ovEzjRAPTOINMz5moB+r7c1eeqmUmbhsTRMlFuzHqRk9z+OJOhCfdN3\n" +
		"8Y0Rj+pQgc5qX6WF3eFrJ3FbMrfDoVLC/ctmefDQ4FNZSHrDOdOsLAkNROJmrJhi\n" +
		"ZF+OiGncgZsXJfhpGQIDAQAB\n" +
		"-----END PUBLIC KEY-----\n"

	origin_ras_privete_key = "-----BEGIN PRIVATE KEY-----\n" +
		"MIICXQIBAAKBgQDN37iIMNod9I3X/7Vfh4YSwNHmaFAr0Qsc5ovEzjRAPTOINMz5\n" +
		"moB+r7c1eeqmUmbhsTRMlFuzHqRk9z+OJOhCfdN38Y0Rj+pQgc5qX6WF3eFrJ3Fb\n" +
		"MrfDoVLC/ctmefDQ4FNZSHrDOdOsLAkNROJmrJhiZF+OiGncgZsXJfhpGQIDAQAB\n" +
		"AoGBAMfvAujQeKMqwy4H2X6iwOQpei9HEsTayO4SP56rmzbfuNIIZR/qmetufoBi\n" +
		"nC1WTS/VxjKwybVUhta+/2vuD9f1GCEcmkCr6kA7i05Jd6gFwSbfkRy33xlhjuNC\n" +
		"WJ1PREHfLb35L9rd+XWeZhIzSM7gGxbIqVDVznIe3EtxhmhhAkEA4QV0xWg/m+XE\n" +
		"p/uGyEUEdR5auRNWou340EYxAW+0BrzUAXscPQ6iqO4JGE9zkRhM3z1XkPN9uMZ4\n" +
		"Duj8QBRIrQJBAOo3cVbAAmmd1WT8xPRrevGeRjCJjbE2v/pfYpMYsiMCXV/hRI3v\n" +
		"3Rj11xwDF66EKUrNOOfe71+nWb5cTy4PE50CQQDTvnShhnXE17P0dsXgEsIdC5FH\n" +
		"cyEldFWcd1CKD3kSlgHR2u05r1n1KPk5/Rm8wWck8u5Boj797xTuwuML0YqJAkB7\n" +
		"fxhn4X5kKjDmutEu/60n0Yi49w6bLn8ziS/018S16P1LHQCExsER9C6kOo02G8Ga\n" +
		"C3PB7y7QhPExCoNFifWRAkBYMws0qBKPblu2hteDiFH3nW5vdL3LyhfWNMNnxN/z\n" +
		"mzxsD5PZfOXxO80kiNxfMOKY4hA/y0qJ6D9u4wEdUaqX\n" +
		"-----END PRIVATE KEY-----\n"

	origin_ras_pub_key_pwd = "-----BEGIN PUBLIC KEY-----\n" +
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDN37iIMNod9I3X/7Vfh4YSwNHm\n" +
		"aFAr0Qsc5ovEzjRAPTOINMz5moB+r7c1eeqmUmbhsTRMlFuzHqRk9z+OJOhCfdN3\n" +
		"8Y0Rj+pQgc5qX6WF3eFrJ3FbMrfDoVLC/ctmefDQ4FNZSHrDOdOsLAkNROJmrJhi\n" +
		"ZF+OiGncgZsXJfhpGQIDAQAB\n" +
		"-----END PUBLIC KEY-----\n"

	origin_ras_privete_key_pwd = "-----BEGIN PUBLIC KEY-----\n" +
		"MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALffzE7ZaS7pu41ML5/kUzbYtZ9I5ldX\n" +
		"nCPYd4jCM1sCl5WmJvfJugYOGLrqiaKi8Ceu4y8/YyFgDQFsu1SE2lcCAwEAAQ==\n" +
		"-----END PUBLIC KEY-----\n"
)

func CheckTel(tel string) (res string, err error) {
	o := orm.NewOrm()
	if err = o.Raw(getUserByTelSql, tel).QueryRow(&res); err != nil {
		if err != orm.ErrNoRows {
			return
		}
		log.Error("checkTel failed params(tel %s),error(%v)", res, err)
		return
	}
	return
}

func CheckEmail(email string) (res string, err error) {
	o := orm.NewOrm()
	if err = o.Raw(getUserByEmailSql, email).QueryRow(&res); err != nil {
		if err != orm.ErrNoRows {
			return
		}
		log.Error("checkEmail failed params(email %s),error(%v)", res, err)
		return
	}
	return
}

func CheckAccount(account string) (res string, err error) {
	o := orm.NewOrm()
	if err = o.Raw(getUserByAccountSql, account).QueryRow(&res); err != nil {
		if err != orm.ErrNoRows {
			return
		}
		log.Error("checkEmail failed params(email %s),error(%v)", res, err)
		return
	}
	return
}

func GetUserByUserId(userId string) (user *models.User, err error) {
	o := orm.NewOrm()
	user = new(models.User)
	if err = o.Raw(getUserByUserIdSql, userId).QueryRow(user); err != nil {
		if err != orm.ErrNoRows {
			return
		}
		log.Error("getUserByUserId failed params(userId %s),error(%v)", userId, err)
		return
	}
	return
}

// 校验用户密码
func CheckUserPwd(pwd64 string, userPwd string, userSalt string) (err error) {
	// 1、先用私钥解码
	decryptPwd, err := GetFrontPwd(pwd64)
	if err != nil {
		return err
	}

	// 2、再加盐匹配原密码返回结果
	rasPwdStr, err := CreateUserPwd(decryptPwd, userSalt)
	if rasPwdStr == userPwd {
		return nil
	} else {
		return errors.New("账号或密码不符合要求")
	}
}

// 获取前端密码
func GetFrontPwd(pwd64 string) ([]byte, error) {
	pwd, err := base64.StdEncoding.DecodeString(pwd64)
	if err != nil {
		return nil, err
	}
	decryptPwd, err := util.RsaDecrypt(pwd, []byte(origin_ras_privete_key))
	if err != nil {
		return nil, err
	}
	return decryptPwd, nil
}

// 创建用户密码
func CreateUserPwd(userPwd []byte, userSalt string) (string, error) {
	md5Pwd := md5.Sum([]byte(string(userPwd) + userSalt))
	hexPwdStr := hex.EncodeToString(md5Pwd[:])
	rasPwd, err := util.RsaEncrypt([]byte(hexPwdStr), []byte(origin_ras_pub_key_pwd))
	if err != nil {
		return "", err
	}
	rasPwdStr := base64.StdEncoding.EncodeToString(rasPwd)
	return rasPwdStr, nil
}

func CreateNewPwdByFront(pwd64 string) (string, string, error) {
	// 1. 解码
	decryptPwd, err := GetFrontPwd(pwd64)
	if err != nil {
		return "", "", err
	}

	// 2. 生成新密码
	userSalt := util.GetSaltByTime()
	newPwd, err := CreateUserPwd(decryptPwd, userSalt)
	return userSalt, newPwd, err
}

// 校验账号、手机号、邮箱是否存在
func CheckReg(account, tel, email, nickName string) string {
	o := orm.NewOrm()
	var errMsg string
	if err := o.Raw(checkUserRegSql, account, tel, email, nickName, account, tel, email, nickName).QueryRow(&errMsg); err != nil {
		return ""
	}
	return errMsg
}

// 创建用户插入到数据库当中
func CreateUser(user *models.User) error {
	user.UserID = util.GetUuid32()
	user.Status = "1"
	user.IsDelete = "0"

	o := orm.NewOrm()
	if _, err := o.Insert(user); err != nil {
		return err
	}
	return nil
}

func GetFrontUserPubKey() string {
	return origin_ras_pub_key
}