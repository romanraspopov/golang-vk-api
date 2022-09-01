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

// FriendsGet возвращает список идентификаторов друзей пользователя
// или расширенную информацию о друзьях пользователя (при использовании параметра fields).
//
// Если вы используете социальный граф пользователя ВКонтакте в своём приложении,
// обратите внимание на п. 4.4. Правил платформы (https://dev.vk.com/rules)
//
// user_id (integer) - Идентификатор пользователя, для которого необходимо получить список друзей.
// Если параметр не задан, то считается, что он равен идентификатору текущего пользователя
// (справедливо для вызова с передачей access_token).
//
// order (string) - Порядок, в котором нужно вернуть список друзей. Допустимые значения:
// hints — сортировать по рейтингу, аналогично тому, как друзья сортируются в разделе Мои друзья.
// Это значение доступно только для Standalone-приложений с ключом доступа, полученным по схеме Implicit Flow.
// random — возвращает друзей в случайном порядке.
// name — сортировать по имени. Данный тип сортировки работает медленно,
// так как сервер будет получать всех друзей а не только указанное количество count.
// (работает только при переданном параметре fields).
// По умолчанию список сортируется в порядке возрастания идентификаторов пользователей.
//
// count (positive) - Количество друзей, которое нужно вернуть. 0 - вернуть всех друзей.
// При использовании параметра fields возвращается не более 5000 друзей.
//
// offset (positive) - Смещение, необходимое для выборки определенного подмножества друзей.
//
// fields (string) - Список дополнительных полей, которые необходимо вернуть (см. var userFields в user.go).
//
// name_case (string) - Падеж для склонения имени и фамилии пользователя. Возможные значения:
// именительный – nom;
// родительный – gen;
// дательный – dat;
// винительный – acc;
// творительный – ins;
// предложный – abl.
// По умолчанию nom.
//
// Результат. После успешного выполнения возвращает список идентификаторов (id) друзей пользователя,
// если параметр fields не использовался.
// При использовании параметра fields возвращает список объектов пользователей, но не более 5000.
func (client *VKClient) FriendsGet(uid int, count int, offset int, fields string) (int, []*User, error) {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(uid))
	params.Set("count", strconv.Itoa(count))
	params.Set("offset", strconv.Itoa(offset))
	if fields != "" {
		params.Set("fields", fields)
	}

	resp, err := client.MakeRequest("friends.get", params)
	if err != nil {
		return 0, nil, err
	}

	var friends *Friends
	json.Unmarshal(resp.Response, &friends)
	return friends.Count, friends.Users, nil
}

// FriendsGetRequests возвращает информацию о полученных или отправленных заявках на добавление в друзья для текущего пользователя.
//
// offset (positive) - Смещение, необходимое для выборки определенного подмножества заявок на добавление в друзья.
//
// count (positive) - Максимальное количество заявок на добавление в друзья,
// которые необходимо получить (не более 1000). По умолчанию — 100.
//
// extended (checkbox) - Определяет, требуется ли возвращать в ответе сообщения от пользователей,
// подавших заявку на добавление в друзья. И отправителя рекомендации при suggested = 1.
//
// need_mutual (checkbox) - Определяет, требуется ли возвращать в ответе список общих друзей, если они есть.
// Обратите внимание, что при использовании need_mutual будет возвращено не более 2 заявок.
//
// out (checkbox) - 0 — возвращать полученные заявки в друзья (по умолчанию), 1 — возвращать отправленные пользователем заявки.
//
// sort (positive) - 0 — сортировать по дате добавления, 1 — сортировать по количеству общих друзей.
// (Если out = 1, этот параметр не учитывается).
//
// Результат. Если не установлен параметр need_mutual, то в случае успеха
// возвращает отсортированный в антихронологическом порядке по времени подачи заявки
// список идентификаторов (id) пользователей (кому или от кого пришла заявка).
// Если установлен параметр need_mutual, то в случае успеха
// возвращает отсортированный в антихронологическом порядке по времени подачи заявки массив объектов,
// содержащих информацию о заявках на добавление в друзья. Каждый из объектов содержит поле uid,
// являющийся идентификатором пользователя. При наличии общих друзей, в объекте будет содержаться поле mutual,
// в котором будет находиться список идентификаторов общих друзей.
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
// text (string) - Текст сопроводительного сообщения для заявки на добавление в друзья. Максимальная длина сообщения - 500 символов.
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
