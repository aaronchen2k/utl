package repo

import (
	"fmt"
	"github.com/utlai/utl/internal/server/biz/domain"
	"github.com/utlai/utl/internal/server/biz/redis"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"strings"
	"time"
)

type TokenRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewTokenRepo() *TokenRepo {
	return &TokenRepo{}
}

func (r *TokenRepo) GetRedisSession(conn *redisUtils.RedisCluster, token string) (*domain.UserCredentials, error) {
	sKey := domain.ZXW_SESSION_TOKEN_PREFIX + token
	if !conn.Exists(sKey) {
		return nil, domain.ERR_TOKEN_INVALID
	}
	pp := new(domain.UserCredentials)
	if err := r.loadRedisHashToStruct(conn, sKey, pp); err != nil {
		return nil, err
	}
	return pp, nil
}
func (r *TokenRepo) GetWebSession(conn *redisUtils.RedisCluster, token string) (*domain.UserCredentials, error) {
	sKey := domain.ZXW_SESSION_TOKEN_PREFIX + token
	if !conn.Exists(sKey) {
		return nil, domain.ERR_TOKEN_INVALID
	}
	pp := new(domain.UserCredentials)
	if err := r.loadRedisHashToStruct(conn, sKey, pp); err != nil {
		return nil, err
	}
	return pp, nil
}

func (r *TokenRepo) loadRedisHashToStruct(conn *redisUtils.RedisCluster, sKey string, pst interface{}) error {
	vals, err := redis.Values(conn.HGetAll(sKey))
	if err != nil {
		return err
	}
	err = redis.ScanStruct(vals, pst)
	if err != nil {
		return err
	}
	return nil
}

func (r *TokenRepo) IsUserTokenOver(userId string) bool {
	conn := redisUtils.GetRedisClusterClient()
	defer conn.Close()
	if r.getUserTokenCount(conn, userId) >= r.getUserTokenMaxCount(conn) {
		return true
	}
	return false
}

func (r *TokenRepo) getUserTokenCount(conn *redisUtils.RedisCluster, userId string) int {
	count, err := redis.Int(conn.Scard(domain.ZXW_SESSION_USER_PREFIX + userId))
	if err != nil {
		fmt.Println(fmt.Sprintf("getUserTokenCount error :%+v", err))
		return 0
	}
	return count
}

func (r *TokenRepo) getUserTokenMaxCount(conn *redisUtils.RedisCluster) int {
	count, err := redis.Int(conn.GetKey(domain.ZXW_SESSION_USER_MAX_TOKEN_PREFIX))
	if err != nil {
		return domain.ZXW_SESSION_USER_MAX_TOKEN_DEFAULT
	}
	return count
}

func (r *TokenRepo) UserTokenExpired(token string) {
	conn := redisUtils.GetRedisClusterClient()
	defer conn.Close()

	uKey := domain.ZXW_SESSION_BIND_USER_PREFIX + token
	sKeys, err := redis.Strings(conn.Members(uKey))
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.Members key %s error :%+v", uKey, err))
		return
	}
	for _, v := range sKeys {
		if !strings.Contains(v, domain.ZXW_SESSION_USER_PREFIX) {
			continue
		}
		_, err := conn.Do("SREM", v, token)
		if err != nil {
			fmt.Println(fmt.Sprintf("conn.SREM key %s token %s  error :%+v", v, token, err))
			return
		}
	}
	if _, err := conn.Del(uKey); err != nil {
		fmt.Println(fmt.Sprintf("conn.Del key %s error :%+v", uKey, err))
	}
	return
}

func (r *TokenRepo) GetUserScope(userType string) uint64 {
	switch userType {
	case "admin":
		return domain.AdminScope
	}
	return domain.NoneScope
}

func (r *TokenRepo) CacheToRedis(conn *redisUtils.RedisCluster, cred domain.UserCredentials, token string) error {
	sKey := domain.ZXW_SESSION_TOKEN_PREFIX + token

	if _, err := conn.HMSet(sKey,
		"user_id", cred.UserId,
		"login_type", cred.LoginType,
		"auth_type", cred.AuthType,
		"creation_data", cred.CreationDate,
		"expires_in", cred.ExpiresIn,
		"scope", cred.Scope,
	); err != nil {
		fmt.Println(fmt.Sprintf("conn.CacheToRedis error :%+v", err))
		return err
	}
	return nil
}

func (r *TokenRepo) SyncUserTokenCache(conn *redisUtils.RedisCluster, cred domain.UserCredentials, token string) error {
	sKey := domain.ZXW_SESSION_USER_PREFIX + token
	if _, err := conn.Sadd(sKey, token); err != nil {
		fmt.Println(fmt.Sprintf("conn.SyncUserTokenCache1 error :%+v", err))
		return err
	}
	sKey2 := domain.ZXW_SESSION_BIND_USER_PREFIX + token
	_, err := conn.Sadd(sKey2, sKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.SyncUserTokenCache2 error :%+v", err))
		return err
	}
	return nil
}

func (r *TokenRepo) UpdateUserTokenCacheExpire(conn *redisUtils.RedisCluster, rs domain.UserCredentials, token string) error {
	if _, err := conn.Expire(domain.ZXW_SESSION_TOKEN_PREFIX+token, int(r.GetTokenExpire(rs).Seconds())); err != nil {
		fmt.Println(fmt.Sprintf("conn.UpdateUserTokenCacheExpire error :%+v", err))
		return err
	}
	return nil
}

func (r *TokenRepo) GetTokenExpire(rs domain.UserCredentials) time.Duration {
	timeout := domain.RedisSessionTimeoutApp
	if rs.LoginType == domain.LoginTypeWeb {
		timeout = domain.RedisSessionTimeoutWeb
	} else if rs.LoginType == domain.LoginTypeWx {
		timeout = domain.RedisSessionTimeoutWx
	} else if rs.LoginType == domain.LoginTypeAlipay {
		timeout = domain.RedisSessionTimeoutWx
	}
	return timeout
}

func (r *TokenRepo) DelUserTokenCache(conn *redisUtils.RedisCluster, rs domain.UserCredentials, token string) error {
	sKey := domain.ZXW_SESSION_USER_PREFIX + rs.UserId
	_, err := conn.Do("SREM", sKey, token)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.DelUserTokenCache1 error :%+v", err))
		return err
	}
	err = r.DelTokenCache(conn, rs, token)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.DelUserTokenCache2 error :%+v", err))
		return err
	}

	return nil
}
func (r *TokenRepo) DelTokenCache(conn *redisUtils.RedisCluster, rs domain.UserCredentials, token string) error {
	sKey2 := domain.ZXW_SESSION_BIND_USER_PREFIX + token
	_, err := conn.Del(sKey2)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.DelUserTokenCache2 error :%+v", err))
		return err
	}
	sKey3 := domain.ZXW_SESSION_TOKEN_PREFIX + token
	_, err = conn.Del(sKey3)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.DelUserTokenCache3 error :%+v", err))
		return err
	}

	return nil
}

func (r *TokenRepo) CleanUserTokenCache(conn *redisUtils.RedisCluster, rs domain.UserCredentials) error {
	sKey := domain.ZXW_SESSION_USER_PREFIX + rs.UserId
	allTokens, err := redis.Strings(conn.Members(sKey))
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.CleanUserTokenCache1 error :%+v", err))
		return err
	}
	_, err = conn.Del(sKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.CleanUserTokenCache2 error :%+v", err))
		return err
	}

	for _, token := range allTokens {
		err = r.DelTokenCache(conn, rs, token)
		if err != nil {
			fmt.Println(fmt.Sprintf("conn.DelUserTokenCache2 error :%+v", err))
			return err
		}
	}
	return nil
}
