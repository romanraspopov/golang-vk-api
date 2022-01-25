package vkapi

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Методы для работы с друзьями

type FriendsRequests struct {
	Count    int        `json:"count"`
	Requests []*Request `json:"items"`
}

type Friends struct {
	Count int     `json:"count"`
	Users []*User `json:"items"`
}

type Request struct {
	UserID        int     `json:"user_id"`
	MutualFriends *Mutual `json:"mutual"`
}

type Mutual struct {
	Count int   `json:"count"`
	Users []int `json:"users"`
}

type FriendDeleteResult struct {
	Success             int `json:"success"`             // удалось успешно удалить друга;
	Friend_deleted      int `json:"friend_deleted"`      // был удален друг;
	Out_request_deleted int `json:"out_request_deleted"` // отменена исходящая заявка;
	In_request_deleted  int `json:"in_request_deleted"`  // отклонена входящая заявка;
	Suggestion_deleted  int `json:"suggestion_deleted"`  // отклонена рекомендация друга.
}

func (client *VKClient) FriendsGet(uid int, count int) (int, []*User, error) {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(uid))
	params.Set("count", strconv.Itoa(count))
	params.Set("fields", userFields)

	resp, err := client.MakeRequest("friends.get", params)
	if err != nil {
		return 0, nil, err
	}

	var friends *Friends
	json.Unmarshal(resp.Response, &friends)
	return friends.Count, friends.Users, nil
}

func (client *VKClient) FriendsGetRequests(count int, out int) (int, []*Request, error) {
	params := url.Values{}
	params.Set("count", strconv.Itoa(count))
	params.Set("out", strconv.Itoa(out))
	params.Set("extended", "1")

	resp, err := client.MakeRequest("friends.getRequests", params)
	if err != nil {
		return 0, nil, err
	}

	var reqs *FriendsRequests
	json.Unmarshal(resp.Response, &reqs)
	return reqs.Count, reqs.Requests, nil
}

// FriendsAdd - одобряет или создаёт заявку на добавление в друзья.
// user_id (positive) - Идентификатор пользователя, которому необходимо отправить заявку, либо заявку от которого необходимо одобрить.
// text (string) - Текст сопроводительного сообщения для заявки на добавление в друзья. Максимальная длина сообщения — 500 символов.
// follow (checkbox) = 1, если необходимо отклонить входящую заявку (оставить пользователя в подписчиках).
//
// После успешного выполнения возвращает одно из следующих значений:
// 1 - заявка на добавление данного пользователя в друзья отправлена;
// 2 - заявка на добавление в друзья от данного пользователя одобрена;
// 4 - повторная отправка заявки.
func (client *VKClient) FriendsAdd(userID int, text string, follow int) (int, error) {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(userID))
	params.Set("follow", strconv.Itoa(follow))
	if text != "" {
		params.Set("text", text)
	}

	resp, err := client.MakeRequest("friends.add", params)
	if err != nil {
		return 0, err
	}

	var result int
	json.Unmarshal(resp.Response, &result)

	return result, nil
}

// FriendsDelete - удаляет пользователя из списка друзей или отклоняет заявку в друзья.
// user_id (positive) - Идентификатор пользователя, которого необходимо удалить из списка друзей, либо заявку от которого необходимо отклонить.
// Результат - возвращается объект с полями (типа "success":1):
// success - удалось успешно удалить друга;
// friend_deleted - был удален друг;
// out_request_deleted - отменена исходящая заявка;
// in_request_deleted - отклонена входящая заявка;
// suggestion_deleted - отклонена рекомендация друга.
func (client *VKClient) FriendsDelete(userID int) (bool, error) {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(userID))

	resp, err := client.MakeRequest("friends.delete", params)
	if err != nil {
		return false, err
	}

	var result FriendDeleteResult
	json.Unmarshal(resp.Response, &result)

	return IntToBool(result.Success), nil
}
