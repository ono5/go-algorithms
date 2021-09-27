package proxy

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(2342423)
}

func TestUserListProxy(t *testing.T) {
	mockedDatabase := UserList{}

	// ランダムなUserIDを作成する
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		mockedDatabase = append(
			mockedDatabase,
			User{ID: n})
	}

	// Proxy設定(mock database, cache, size)
	proxy := UserListProxy{
		MockedDatabase: &mockedDatabase,
		StackCache:     UserList{},
		StackSize:      2,
	}

	// いくつかIDを取り出す
	knowIDs := [3]int32{
		mockedDatabase[3].ID,
		mockedDatabase[4].ID,
		mockedDatabase[5].ID,
	}

	// Proxyを経由してUserIDを取得
	t.Run("FindUser -Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knowIDs[0])
		if err != nil {
			t.Error(err)
		}
		if user.ID != knowIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it must be one")
		}
		if proxy.LastSearchUsedCache == true {
			t.Error("No user can be returned from an empty cache")
		}
	})

	// ↑のテストでProxyのパラメータは書き換えられる
	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		proxy.FindUser(knowIDs[0])

		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an object that is stored on it")
		}
		if !proxy.LastSearchUsedCache {
			t.Error("The user should have been returned from the cache")
		}
	})

	t.Run("FindUser - overflowing the satck", func(t *testing.T) {
		user1, err := proxy.FindUser(knowIDs[0])
		if err != nil {
			t.Error(err)
		}
		user2, err := proxy.FindUser(knowIDs[1])
		if proxy.LastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}
		user3, err := proxy.FindUser(knowIDs[2])
		if proxy.LastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}
		for i := 0; i < len(proxy.StackCache); i++ {
			if proxy.StackCache[i].ID == user1.ID {
				t.Error("User that should be gone was found")
			}
		}
		if len(proxy.StackCache) != 2 {
			t.Error("After inserting 3 users the cache should not grow more thant to two")
		}
		for _, v := range proxy.StackCache {
			if v != user2 && v != user3 {
				t.Error("A non expected user was found on the cache")
			}
		}
	})
}
