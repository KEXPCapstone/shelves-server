package sessions

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis"
)

//RedisStore represents a session.Store backed by redis.
type RedisStore struct {
	//Redis client used to talk to redis server.
	Client *redis.Client
	//Used for key expiry time on redis.
	SessionDuration time.Duration
}

//NewRedisStore constructs a new RedisStore
func NewRedisStore(client *redis.Client, sessionDuration time.Duration) *RedisStore {
	if client == nil {
		panic("nil pointer to redis client passed in")
	}
	return &RedisStore{
		Client:          client,
		SessionDuration: sessionDuration,
	}
}

//Store implementation

//Save saves the provided `sessionState` and associated SessionID to the store.
//The `sessionState` parameter is typically a pointer to a struct containing
//all the data you want to associated with the given SessionID.
func (rs *RedisStore) Save(sid SessionID, sessionState interface{}) error {
	j, err := json.Marshal(sessionState)
	if err != nil {
		return errors.New("error marshalling json")
	}

	if err := rs.Client.Set(sid.getRedisKey(), j, rs.SessionDuration).Err(); err != nil {
		return errors.New("error saving to redis")
	}
	return nil
}

//Get populates `sessionState` with the data previously saved
//for the given SessionID
func (rs *RedisStore) Get(sid SessionID, sessionState interface{}) error {
	key := sid.getRedisKey()
	r, err := rs.Client.Get(key).Bytes()
	if err != nil {
		return ErrStateNotFound
	}
	if err := json.Unmarshal(r, sessionState); err != nil {
		return errors.New("error unmarshalling json")
	}
	expErr := rs.Client.Expire(sid.getRedisKey(), rs.SessionDuration).Err()
	if expErr == redis.Nil {
		return ErrStateNotFound
	} else if expErr != nil {
		return errors.New("error setting new expiry time to redis store")
	}
	return nil
}

//Delete deletes all state data associated with the SessionID from the store.
func (rs *RedisStore) Delete(sid SessionID) error {
	err := rs.Client.Del(sid.getRedisKey()).Err()
	if err == redis.Nil {
		return ErrStateNotFound
	} else if err != nil {
		return errors.New("error deleting key from redis")
	}
	return nil
}

//getRedisKey() returns the redis key to use for the SessionID
func (sid SessionID) getRedisKey() string {
	return "sid:" + sid.String()
}
