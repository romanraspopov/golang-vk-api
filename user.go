package vkapi

import (
	"encoding/json"
	"net/url"
)

// Методы для работы с данными пользователей

// last seen device
const (
	_ = iota
	PlatformMobile
	PlatformIPhone
	PlatfromIPad
	PlatformAndroid
	PlatformWPhone
	PlatformWindows
	PlatformWeb
)

var (
	userFields = "nickname,screen_name,sex,bdate,city,country," +
		"photo,photo_medium,photo_big,has_mobile,contacts," +
		"education,online,relation,last_seen,activity," +
		"can_write_private_message,can_see_all_posts,can_post,universities"
	userFieldsMin  = "id,first_name,last_name,deactivated,is_closed,can_access_closed"
	userFieldsFull = "id,first_name,last_name,deactivated,is_closed,can_access_closed," +
		"about,activities,bdate,blacklisted,blacklisted_by_me,books,can_post,can_see_all_posts," +
		"can_see_audio,can_send_friend_request,can_write_private_message,career,city,common_count," +
		"connections,contacts,counters,country,domain,education,exports," +
		"first_name_nom,first_name_gen,first_name_dat,first_name_acc,first_name_ins,first_name_abl," +
		"followers_count,friend_status,games,has_mobile,has_photo,home_town,interests,is_favorite,is_friend,is_hidden_from_feed," +
		"is_no_index,last_name_nom,last_name_gen,last_name_dat,last_name_acc,last_name_ins,last_name_abl,last_seen,lists,maiden_name," +
		"military,movies,music,nickname,occupation,online,personal,photo_50,photo_100,photo_200_orig,photo_200,photo_400_orig," +
		"photo_id,photo_max,photo_max_orig,quotes,relatives,relation,schools,screen_name,sex,site,status,timezone," +
		"trending,tv,universities,verified,wall_default"
)

// User содержит все (ли?) личные данные пользователя ВК
type UserNorm struct {
	//type User struct {
	UID                     int          `json:"id"`
	FirstName               string       `json:"first_name"`
	LastName                string       `json:"last_name"`
	Sex                     int          `json:"sex"`
	Nickname                string       `json:"nickname"`
	ScreenName              string       `json:"screen_name"`
	BDate                   string       `json:"bdate"`
	City                    *UserCity    `json:"city"`
	Country                 *UserCountry `json:"country"`
	Photo                   string       `json:"photo"`
	PhotoMedium             string       `json:"photo_medium"`
	PhotoBig                string       `json:"photo_big"`
	Photo50                 string       `json:"photo_50"`
	Photo100                string       `json:"photo_100"`
	HasMobile               int          `json:"has_mobile"`
	Online                  int          `json:"online"`
	OnlineInfo              *OnlineInfo  `json:"online_info"`
	CanPost                 int          `json:"can_post"`
	CanSeeAllPosts          int          `json:"can_see_all_posts"`
	CanWritePrivateMessages int          `json:"can_write_private_message"`
	Status                  string       `json:"status"`
	LastSeen                *LastSeen    `json:"last_seen"`
	Hidden                  int          `json:"hidden"`
	Deactivated             string       `json:"deactivated"`
	Relation                int          `json:"relation"`
}

// User содержит ВСЕ личные данные пользователя ВК
type User struct {
	UID                    int             `json:"id"`                        // Идентификатор пользователя.
	FirstName              string          `json:"first_name"`                // Имя.
	LastName               string          `json:"last_name"`                 // Фамилия.
	Deactivated            string          `json:"deactivated"`               // Поле возвращается, если страница пользователя удалена или заблокирована, содержит значение deleted или banned. В этом случае опциональные поля не возвращаются.
	IsClosed               bool            `json:"is_closed"`                 // Скрыт ли профиль пользователя настройками приватности.
	CanAccessClosed        bool            `json:"can_access_closed"`         // Может ли текущий пользователь видеть профиль при is_closed = 1 (например, он есть в друзьях).
	About                  string          `json:"about"`                     // Содержимое поля «О себе» из профиля.
	Activities             string          `json:"activities"`                // Содержимое поля «Деятельность» из профиля.
	Bdate                  string          `json:"bdate"`                     // Дата рождения. Возвращается в формате D.M.YYYY или D.M (если год рождения скрыт). Если дата рождения скрыта целиком, поле отсутствует в ответе.
	Blacklisted            int             `json:"blacklisted"`               // Информация о том, находится ли текущий пользователь в черном списке. Возможные значения: 1 — находится; 0 — не находится.
	BlacklistedByMe        int             `json:"blacklisted_by_me"`         // Информация о том, находится ли пользователь в черном списке у текущего пользователя. Возможные значения: 1 — находится; 0 — не находится.
	Books                  string          `json:"books"`                     // Содержимое поля «Любимые книги» из профиля пользователя.
	CanPost                int             `json:"can_post"`                  // Информация о том, может ли текущий пользователь оставлять записи на стене. Возможные значения: 1 — может; 0 — не может.
	CanSeeAllPosts         int             `json:"can_see_all_posts"`         // Информация о том, может ли текущий пользователь видеть чужие записи на стене. Возможные значения: 1 — может; 0 — не может.
	CanSeeAudio            int             `json:"can_see_audio"`             // Информация о том, может ли текущий пользователь видеть аудиозаписи. Возможные значения: 1 — может; 0 — не может.
	CanSendFriendRequest   int             `json:"can_send_friend_request"`   // Информация о том, будет ли отправлено уведомление пользователю о заявке в друзья от текущего пользователя. Возможные значения: 1 — уведомление будет отправлено; 0 — уведомление не будет отправлено.
	CanWritePrivateMessage int             `json:"can_write_private_message"` // Информация о том, может ли текущий пользователь отправить личное сообщение. Возможные значения: 1 — может; 0 — не может.
	Career                 []Career        `json:"career"`                    // Информация о карьере пользователя. Объект, содержащий следующие поля: group_id (integer) — идентификатор сообщества (если доступно, иначе company); company (string) — название компании (если доступно, иначе group_id); country_id (integer) — идентификатор страны; city_id (integer) — идентификатор города (если доступно, иначе city_name); city_name (string) — название города (если доступно, иначе city_id); from (integer) — год начала работы; until (integer) — год окончания работы; position (string) — должность.
	City                   *UserCity       `json:"city"`                      // Информация о городе, указанном на странице пользователя в разделе «Контакты». Возвращаются следующие поля: id (integer) — идентификатор города, который можно использовать для получения его названия с помощью метода database.getCitiesById; title (string) — название города.
	CommonCount            int             `json:"common_count"`              // Количество общих друзей с текущим пользователем.
	Connections            string          `json:"connections"`               // Возвращает данные об указанных в профиле сервисах пользователя, таких как: skype, livejournal. Для каждого сервиса возвращается отдельное поле с типом string, содержащее никнейм пользователя. Например, "skype": "username".
	Contacts               *UserContacts   `json:"contacts"`                  // Информация о телефонных номерах пользователя. Если данные указаны и не скрыты настройками приватности, возвращаются следующие поля: mobile_phone (string) — номер мобильного телефона пользователя (только для Standalone-приложений); home_phone (string) — дополнительный номер телефона пользователя.
	Counters               *UserCounters   `json:"counters"`                  // Количество различных объектов у пользователя. Поле возвращается только в методе users.get при запросе информации об одном пользователе, с передачей пользовательского access_token. Объект, содержащий следующие поля: albums (integer) — количество фотоальбомов; videos (integer) — количество видеозаписей; audios (integer) — количество аудиозаписей; photos (integer) — количество фотографий; notes (integer) — количество заметок; friends (integer) — количество друзей; groups (integer) — количество сообществ; online_friends (integer) — количество друзей онлайн; mutual_friends (integer) — количество общих друзей; user_videos (integer) — количество видеозаписей с пользователем; followers (integer) — количество подписчиков; pages (integer) — количество объектов в блоке «Интересные страницы».
	Country                *UserCountry    `json:"country"`                   // Информация о стране, указанной на странице пользователя в разделе «Контакты». Возвращаются следующие поля: id (integer) — идентификатор страны, который можно использовать для получения ее названия с помощью метода database.getCountriesById; title (string) — название страны.
	Domain                 string          `json:"domain"`                    // Короткий адрес страницы. Возвращается строка, содержащая короткий адрес страницы (например, andrew). Если он не назначен, возвращается "id"+user_id, например, id35828305.
	Education              *UserEducation  `json:"education"`                 // Информация о высшем учебном заведении пользователя. Возвращаются поля: university (integer) — идентификатор университета; university_name (string) — название университета; faculty (integer) — идентификатор факультета; faculty_name (string)— название факультета; graduation (integer) — год окончания.
	Exports                string          `json:"exports"`                   // Внешние сервисы, в которые настроен экспорт из ВК ( livejournal).
	FirstNameCaseNom       string          `json:"first_name_nom"`            // Имя в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; acc — винительный; ins — творительный; abl — предложный. В запросе можно передать несколько значений
	FirstNameCaseGen       string          `json:"first_name_gen"`            // Имя в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; acc — винительный; ins — творительный; abl — предложный. В запросе можно передать несколько значений
	FirstNameCaseDat       string          `json:"first_name_dat"`            // Имя в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; acc — винительный; ins — творительный; abl — предложный. В запросе можно передать несколько значений
	FirstNameCaseAcc       string          `json:"first_name_acc"`            // Имя в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; acc — винительный; ins — творительный; abl — предложный. В запросе можно передать несколько значений
	FirstNameCaseIns       string          `json:"first_name_ins"`            // Имя в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; acc — винительный; ins — творительный; abl — предложный. В запросе можно передать несколько значений
	FirstNameCaseAbl       string          `json:"first_name_abl"`            // Имя в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; acc — винительный; ins — творительный; abl — предложный. В запросе можно передать несколько значений
	FollowersCount         int             `json:"followers_count"`           // Количество подписчиков пользователя.
	FriendStatus           int             `json:"friend_status"`             // Статус дружбы с пользователем. Возможные значения: 0 — не является другом, 1 — отправлена заявка/подписка пользователю, 2 — имеется входящая заявка/подписка от пользователя, 3 — является другом.
	Games                  string          `json:"games"`                     // Содержимое поля «Любимые игры» из профиля.
	HasMobile              int             `json:"has_mobile"`                // Информация о том, известен ли номер мобильного телефона пользователя. Возвращаемые значения: 1 — известен, 0 — не известен.
	HasPhoto               int             `json:"has_photo"`                 // Информация о том, установил ли пользователь фотографию для профиля. Возвращаемые значения: 1 — установил, 0 — не установил.
	HomeTown               string          `json:"home_town"`                 // Название родного города.
	Interests              string          `json:"interests"`                 // Содержимое поля «Интересы» из профиля.
	IsFavorite             int             `json:"is_favorite"`               // Информация о том, есть ли пользователь в закладках у текущего пользователя. Возможные значения: 1 — есть; 0 — нет.
	IsFriend               int             `json:"is_friend"`                 // Информация о том, является ли пользователь другом текущего пользователя. Возможные значения: 1 — да; 0 — нет.
	IsHiddenFromFeed       int             `json:"is_hidden_from_feed"`       // Информация о том, скрыт ли пользователь из ленты новостей текущего пользователя. Возможные значения: 1 — да; 0 — нет.
	IsNoIndex              int             `json:"is_no_index"`               // Индексируется ли профиль поисковыми сайтами. Возможные значения: 1 — профиль скрыт от поисковых сайтов; 0 — профиль доступен поисковым сайтам. (В настройках приватности: https://vk.com/settings?act=privacy, в пункте «Кому в интернете видна моя страница», выбрано значение «Всем»).
	LastNameCaseNom        string          `json:"last_name_nom"`             // Фамилия в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; ac — винительный; ins — творительный; abl — предложный.
	LastNameCaseGen        string          `json:"last_name_gen"`             // Фамилия в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; ac — винительный; ins — творительный; abl — предложный.
	LastNameCaseDat        string          `json:"last_name_dat"`             // Фамилия в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; ac — винительный; ins — творительный; abl — предложный.
	LastNameCaseAcc        string          `json:"last_name_acc"`             // Фамилия в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; ac — винительный; ins — творительный; abl — предложный.
	LastNameCaseIns        string          `json:"last_name_ins"`             // Фамилия в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; ac — винительный; ins — творительный; abl — предложный.
	LastNameCaseAbl        string          `json:"last_name_abl"`             // Фамилия в заданном падеже. Возможные значения для {case}: nom — именительный; gen — родительный; dat — дательный; ac — винительный; ins — творительный; abl — предложный.
	LastSeen               *LastSeen       `json:"last_seen"`                 // Время последнего посещения. Объект, содержащий следующие поля: time (integer) — время последнего посещения в формате Unixtime. platform (integer) — тип платформы. Возможные значения: 1 — мобильная версия; 2 — приложение для iPhone; 3 — приложение для iPad; 4 — приложение для Android; 5 — приложение для Windows Phone; 6 — приложение для Windows 10; 7 — полная версия сайта.
	Lists                  string          `json:"lists"`                     // Разделенные запятой идентификаторы списков друзей, в которых состоит пользователь. Поле доступно только для метода friends.get.
	MaidenName             string          `json:"maiden_name"`               // Девичья фамилия.
	Military               []Military      `json:"military"`                  // Информация о военной службе пользователя. Объект, содержащий следующие поля: unit (string) — номер части; unit_id (integer) — идентификатор части в базе данных; country_id (integer) — идентификатор страны, в которой находится часть; from (integer) — год начала службы; until (integer) — год окончания службы.
	Movies                 string          `json:"movies"`                    // Содержимое поля «Любимые фильмы» из профиля пользователя.
	Music                  string          `json:"music"`                     // Содержимое поля «Любимая музыка» из профиля пользователя.
	Nickname               string          `json:"nickname"`                  // Никнейм (отчество) пользователя.
	Occupation             *UserOccupation `json:"occupation"`                // Информация о текущем роде занятия пользователя. Объект, содержащий следующие поля: type (string) — тип. Возможные значения: work — работа; school — среднее образование; university — высшее образование. id (integer) — идентификатор школы, вуза, сообщества компании (в которой пользователь работает); name (string) — название школы, вуза или места работы;
	Online                 int             `json:"online"`                    // Информация о том, находится ли пользователь сейчас на сайте. Если пользователь использует мобильное приложение либо мобильную версию, возвращается дополнительное поле online_mobile, содержащее 1. При этом, если используется именно приложение, дополнительно возвращается поле online_app, содержащее его идентификатор.
	Personal               *Personal       `json:"personal"`                  // Информация о полях из раздела «Жизненная позиция». political (integer) — политические предпочтения. Возможные значения: 1 — коммунистические; 2 — социалистические; 3 — умеренные; 4 — либеральные; 5 — консервативные; 6 — монархические; 7 — ультраконсервативные; 8 — индифферентные; 9 — либертарианские. langs (array) — языки. religion (string) — мировоззрение. inspired_by (string) — источники вдохновения. people_main (integer) — главное в людях. Возможные значения: 1 — ум и креативность; 2 — доброта и честность; 3 — красота и здоровье; 4 — власть и богатство; 5 — смелость и упорство; 6 — юмор и жизнелюбие. life_main (integer) — главное в жизни. Возможные значения: 1 — семья и дети; 2 — карьера и деньги; 3 — развлечения и отдых; 4 — наука и исследования; 5 — совершенствование мира; 6 — саморазвитие; 7 — красота и искусство; 8 — слава и влияние; smoking (integer) — отношение к курению. Возможные значения: 1 — резко негативное; 2 — негативное; 3 — компромиссное; 4 — нейтральное; 5 — положительное. alcohol (integer) — отношение к алкоголю. Возможные значения: 1 — резко негативное; 2 — негативное; 3 — компромиссное; 4 — нейтральное; 5 — положительное.
	Photo                  string          `json:"photo"`                     //
	PhotoMedium            string          `json:"photo_medium"`              //
	PhotoBig               string          `json:"photo_big"`                 //
	Photo50                string          `json:"photo_50"`                  // URL квадратной фотографии пользователя, имеющей ширину 50 пикселей. В случае отсутствия у пользователя фотографии возвращается https://vk.com/images/camera_50.png.
	Photo100               string          `json:"photo_100"`                 // URL квадратной фотографии пользователя, имеющей ширину 100 пикселей. В случае отсутствия у пользователя фотографии возвращается https://vk.com/images/camera_100.png.
	Photo200Orig           string          `json:"photo_200_orig"`            // URL фотографии пользователя, имеющей ширину 200 пикселей. В случае отсутствия у пользователя фотографии возвращается https://vk.com/images/camera_200.png.
	Photo200               string          `json:"photo_200"`                 // URL квадратной фотографии, имеющей ширину 200 пикселей. Если у пользователя отсутствует фотография таких размеров, в ответе вернется https://vk.com/images/camera_200.png
	Photo400Orig           string          `json:"photo_400_orig"`            // URL фотографии, имеющей ширину 400 пикселей. Если у пользователя отсутствует фотография такого размера, в ответе вернется https://vk.com/images/camera_400.png.
	PhotoId                string          `json:"photo_id"`                  // Строковый идентификатор главной фотографии профиля пользователя в формате {user_id}_{photo_id}, например, 6492_192164258. Обратите внимание, это поле может отсутствовать в ответе.
	PhotoMax               string          `json:"photo_max"`                 // URL квадратной фотографии с максимальной шириной. Может быть возвращена фотография, имеющая ширину как 200, так и 100 пикселей. В случае отсутствия у пользователя фотографии возвращается https://vk.com/images/camera_200.png.
	PhotoMaxOrig           string          `json:"photo_max_orig"`            // URL фотографии максимального размера. Может быть возвращена фотография, имеющая ширину как 400, так и 200 пикселей. В случае отсутствия у пользователя фотографии возвращается https://vk.com/images/camera_400.png.
	Quotes                 string          `json:"quotes"`                    // Любимые цитаты.
	Relatives              []Relative      `json:"relatives"`                 // Список родственников. Массив объектов, каждый из которых содержит поля: id (integer) — идентификатор пользователя; name (string) — имя родственника (если родственник не является пользователем ВКонтакте, то предыдущее значение id возвращено не будет); type (string) — тип родственной связи. Возможные значения: child — сын/дочь; sibling — брат/сестра; parent — отец/мать; grandparent — дедушка/бабушка; grandchild — внук/внучка.
	Relation               int             `json:"relation"`                  // Семейное положение. Возможные значения: 1 — не женат/не замужем; 2 — есть друг/есть подруга; 3 — помолвлен/помолвлена; 4 — женат/замужем; 5 — всё сложно; 6 — в активном поиске; 7 — влюблён/влюблена; 8 — в гражданском браке; 0 — не указано. Если в семейном положении указан другой пользователь, дополнительно возвращается объект relation_partner, содержащий id и имя этого человека.
	Schools                []School        `json:"schools"`                   // Список школ, в которых учился пользователь. Массив объектов, описывающих школы. Каждый объект содержит следующие поля: id (string) — идентификатор школы; country (integer) — идентификатор страны, в которой расположена школа; city (integer) — идентификатор города, в котором расположена школа; name (string) — наименование школы year_from (integer) — год начала обучения; year_to (integer) — год окончания обучения; year_graduated (integer) — год выпуска; class (string) — буква класса; speciality (string) — специализация; type (integer) — идентификатор типа; type_str (string) — название типа. Возможные значения для пар type-typeStr: 0 — "школа"; 1 — "гимназия"; 2 — "лицей"; 3 — "школа-интернат"; 4 — "школа вечерняя"; 5 — "школа музыкальная"; 6 — "школа спортивная"; 7 — "школа художественная"; 8 — "колледж"; 9 — "профессиональный лицей"; 10 — "техникум"; 11 — "ПТУ"; 12 — "училище"; 13 — "школа искусств".
	ScreenName             string          `json:"screen_name"`               // Короткое имя страницы.
	Sex                    int             `json:"sex"`                       // Пол. Возможные значения: 1 — женский; 2 — мужской; 0 — пол не указан.
	Site                   string          `json:"site"`                      // Адрес сайта, указанный в профиле.
	Status                 string          `json:"status"`                    // Статус пользователя. Возвращается строка, содержащая текст статуса, расположенного в профиле под именем. Если включена опция «Транслировать в статус играющую музыку», возвращается дополнительное поле status_audio, содержащее информацию о композиции.
	Timezone               int             `json:"timezone"`                  // Временная зона. Только при запросе информации о текущем пользователе.
	Trending               int             `json:"trending"`                  // Информация о том, есть ли на странице пользователя «огонёк».
	Tv                     string          `json:"tv"`                        // Любимые телешоу.
	Universities           []University    `json:"universities"`              // Список вузов, в которых учился пользователь. Массив объектов, описывающих университеты. Каждый объект содержит следующие поля: id (integer)— идентификатор университета; country (integer) — идентификатор страны, в которой расположен университет; city (integer) — идентификатор города, в котором расположен университет; name (string) — наименование университета; faculty (integer) — идентификатор факультета; faculty_name (string) — наименование факультета; chair (integer) — идентификатор кафедры; chair_name (string) — наименование кафедры; graduation (integer) — год окончания обучения; education_form (string) — форма обучения; education_status (string) — статус (например, «Выпускник (специалист)»).
	Verified               int             `json:"verified"`                  // Возвращается 1, если страница пользователя верифицирована, 0 — если нет.
	WallDefault            string          `json:"wall_default"`              // Режим стены по умолчанию. Возможные значения: owner, all.
}

type UserMin struct {
	UID             int    `json:"id"`                // Идентификатор пользователя.
	FirstName       string `json:"first_name"`        // Имя.
	LastName        string `json:"last_name"`         // Фамилия.
	Deactivated     string `json:"deactivated"`       // Поле возвращается, если страница пользователя удалена или заблокирована, содержит значение deleted или banned. В этом случае опциональные поля не возвращаются.
	IsClosed        bool   `json:"is_closed"`         // Скрыт ли профиль пользователя настройками приватности.
	CanAccessClosed bool   `json:"can_access_closed"` // Может ли текущий пользователь видеть профиль при is_closed = 1 (например, он есть в друзьях).
}

// UserCity содержит id и название населенного пункта пользователя ВК
// Информация о городе, указанном на странице пользователя в разделе «Контакты».
// Возвращаются следующие поля:
// id (integer) — идентификатор города, который можно использовать для получения его названия с помощью метода database.getCitiesById;
// title (string) — название города.
type UserCity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// UserCountry содержит id и название страны пользователя ВК
// Информация о стране, указанной на странице пользователя в разделе «Контакты».
// Возвращаются следующие поля:
// id (integer) — идентификатор страны, который можно использовать для получения ее названия с помощью метода database.getCountriesById;
// title (string) — название страны.
type UserCountry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// LastSeen содержит информацию о последнем посещении пользователя (время и тип устройства).
// time (integer) — время последнего посещения в формате Unixtime.
// platform (integer) — тип платформы. Возможные значения: 1 — мобильная версия; 2 — приложение для iPhone; 3 — приложение для iPad; 4 — приложение для Android; 5 — приложение для Windows Phone; 6 — приложение для Windows 10; 7 — полная версия сайта.
type LastSeen struct {
	Time     int64 `json:"time"`
	Platform int   `json:"platform"`
}

// OnlineInfo содержит информацию о статусе "онлайн" пользователя
type OnlineInfo struct {
	Visible  bool  `json:"visible"`
	LastSeen int64 `json:"last_seen"`
}

// Career содержит информацию о карьере пользователя.
// Объект, содержащий следующие поля:
// group_id (integer) — идентификатор сообщества (если доступно, иначе company);
// company (string) — название компании (если доступно, иначе group_id);
// country_id (integer) — идентификатор страны;
// city_id (integer) — идентификатор города (если доступно, иначе city_name);
// city_name (string) — название города (если доступно, иначе city_id);
// from (integer) — год начала работы;
// until (integer) — год окончания работы;
// position (string) — должность.
type Career struct {
	GroupID   int    `json:"group_id"`
	Company   string `json:"company"`
	CountryID int    `json:"country_id"`
	CityID    int    `json:"city_id"`
	CityName  string `json:"city_name"`
	From      int    `json:"from"`
	Until     int    `json:"until"`
	Position  string `json:"position"`
}

// Contacts содержит информацию о телефонных номерах пользователя.
// Если данные указаны и не скрыты настройками приватности, возвращаются следующие поля:
// mobile_phone (string) — номер мобильного телефона пользователя (только для Standalone-приложений);
// home_phone (string) — дополнительный номер телефона пользователя.
type UserContacts struct {
	MobilePhone string `json:"mobile_phone"`
	HomePhone   string `json:"home_phone"`
}

// Counters - количество различных объектов у пользователя.
// Поле возвращается только в методе users.get при запросе информации об одном пользователе,
// с передачей пользовательского access_token.
// Объект, содержащий следующие поля:
// albums (integer) — количество фотоальбомов;
// videos (integer) — количество видеозаписей;
// audios (integer) — количество аудиозаписей;
// photos (integer) — количество фотографий;
// notes (integer) — количество заметок;
// friends (integer) — количество друзей;
// groups (integer) — количество сообществ;
// online_friends (integer) — количество друзей онлайн;
// mutual_friends (integer) — количество общих друзей;
// user_videos (integer) — количество видеозаписей с пользователем;
// followers (integer) — количество подписчиков;
// pages (integer) — количество объектов в блоке «Интересные страницы».
type UserCounters struct {
	Albums         int `json:"albums"`
	Videos         int `json:"videos"`
	Audios         int `json:"audios"`
	Photos         int `json:"photos"`
	Notes          int `json:"notes"`
	Friends        int `json:"friends"`
	Groups         int `json:"groups"`
	OnlineFriends  int `json:"online_friends"`
	Mutual_friends int `json:"mutual_friends"`
	UserVideos     int `json:"user_videos"`
	Followers      int `json:"followers"`
	Pages          int `json:"pages"`
}

// UserEducation - информация о высшем учебном заведении пользователя.
// Возвращаются поля:
// university (integer) — идентификатор университета;
// university_name (string) — название университета;
// faculty (integer) — идентификатор факультета;
// faculty_name (string)— название факультета;
// graduation (integer) — год окончания.
type UserEducation struct {
	UniversityID   int    `json:"university"`
	UniversityName string `json:"university_name"`
	FacultyID      int    `json:"faculty"`
	FacultyName    string `json:"faculty_name"`
	Graduation     int    `json:"graduation"`
}

// Military - информация о военной службе пользователя.
// Объект, содержащий следующие поля:
// unit (string) — номер части;
// unit_id (integer) — идентификатор части в базе данных;
// country_id (integer) — идентификатор страны, в которой находится часть;
// from (integer) — год начала службы;
// until (integer) — год окончания службы.
type Military struct {
	Unit      string `json:"unit"`
	UnitID    int    `json:"unit_id"`
	CountryID int    `json:"country_id"`
	From      int    `json:"from"`
	Until     int    `json:"until"`
}

// UserOccupation - информация о текущем роде занятия пользователя.
// Объект, содержащий следующие поля:
// type (string) — тип. Возможные значения: work — работа; school — среднее образование; university — высшее образование$
// id (integer) — идентификатор школы, вуза, сообщества компании (в которой пользователь работает);
// name (string) — название школы, вуза или места работы.
type UserOccupation struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Personal - информация о полях из раздела «Жизненная позиция».
// political (integer) — политические предпочтения.
// Возможные значения: 1 — коммунистические; 2 — социалистические; 3 — умеренные; 4 — либеральные; 5 — консервативные; 6 — монархические; 7 — ультраконсервативные; 8 — индифферентные; 9 — либертарианские.
// langs (array) — языки.
// religion (string) — мировоззрение.
// inspired_by (string) — источники вдохновения.
// people_main (integer) — главное в людях.
// Возможные значения: 1 — ум и креативность; 2 — доброта и честность; 3 — красота и здоровье; 4 — власть и богатство; 5 — смелость и упорство; 6 — юмор и жизнелюбие.
// life_main (integer) — главное в жизни.
// Возможные значения: 1 — семья и дети; 2 — карьера и деньги; 3 — развлечения и отдых; 4 — наука и исследования; 5 — совершенствование мира; 6 — саморазвитие; 7 — красота и искусство; 8 — слава и влияние;
// smoking (integer) — отношение к курению.
// Возможные значения: 1 — резко негативное; 2 — негативное; 3 — компромиссное; 4 — нейтральное; 5 — положительное.
// alcohol (integer) — отношение к алкоголю.
// Возможные значения: 1 — резко негативное; 2 — негативное; 3 — компромиссное; 4 — нейтральное; 5 — положительное.
type Personal struct {
	Political  int    `json:"political"`
	Langs      []int  `json:"langs"`
	Religion   string `json:"religion"`
	InspiredBy string `json:"inspired_by"`
	PeopleMain int    `json:"people_main"`
	LifeMain   int    `json:"life_main"`
	Smoking    int    `json:"smoking"`
	Alcohol    int    `json:"alcohol"`
}

// Relative - родственник. Выдается в массиве - списке родсвенников.
// Объект содержит поля:
// id (integer) — идентификатор пользователя;
// name (string) — имя родственника (если родственник не является пользователем ВКонтакте, то предыдущее значение id возвращено не будет);
// type (string) — тип родственной связи.
// Возможные значения: child — сын/дочь; sibling — брат/сестра; parent — отец/мать; grandparent — дедушка/бабушка; grandchild — внук/внучка.
type Relative struct {
	UID  int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// School - школа. Выдается в массиве - списке школ, в которых учился пользователь.
// Объект содержит следующие поля:
// id (string?) — идентификатор школы;
// country (integer) — идентификатор страны, в которой расположена школа;
// city (integer) — идентификатор города, в котором расположена школа;
// name (string) — наименование школы;
// year_from (integer) — год начала обучения;
// year_to (integer) — год окончания обучения;
// year_graduated (integer) — год выпуска;
// class (string) — буква класса;
// speciality (string) — специализация;
// type (integer) — идентификатор типа;
// type_str (string) — название типа.
// Возможные значения для пар type-typeStr: 0 — "школа"; 1 — "гимназия"; 2 — "лицей"; 3 — "школа-интернат"; 4 — "школа вечерняя"; 5 — "школа музыкальная"; 6 — "школа спортивная"; 7 — "школа художественная"; 8 — "колледж"; 9 — "профессиональный лицей"; 10 — "техникум"; 11 — "ПТУ"; 12 — "училище"; 13 — "школа искусств".
type School struct {
	SchoolID      int    `json:"id"`
	CountryID     int    `json:"country"`
	CityID        int    `json:"city"`
	Name          string `json:"name"`
	YearFrom      int    `json:"year_from"`
	YearTo        int    `json:"year_to"`
	YearGraduated int    `json:"year_graduated"`
	Class         string `json:"class"`
	Speciality    string `json:"speciality"`
	Type          int    `json:"type"`
	TypeStr       string `json:"type_str"`
}

// University - университет. Выдается в массиве - списке вузов, в которых учился пользователь.
// Объект содержит следующие поля:
// id (integer)— идентификатор университета;
// country (integer) — идентификатор страны, в которой расположен университет;
// city (integer) — идентификатор города, в котором расположен университет;
// name (string) — наименование университета;
// faculty (integer) — идентификатор факультета;
// faculty_name (string) — наименование факультета;
// chair (integer) — идентификатор кафедры;
// chair_name (string) — наименование кафедры;
// graduation (integer) — год окончания обучения;
// education_form (string) — форма обучения;
// education_status (string) — статус (например, «Выпускник (специалист)»).
type University struct {
	UniversityID    int    `json:"id"`
	CountryID       int    `json:"country"`
	CityID          int    `json:"city"`
	Name            string `json:"name"`
	Faculty         int    `json:"faculty"`
	FacultyName     string `json:"faculty_name"`
	Chair           int    `json:"chair"`
	ChairName       string `json:"chair_name"`
	Graduation      int    `json:"graduation"`
	EducationForm   string `json:"education_form"`
	EducationStatus string `json:"education_status"`
}

func (client *VKClient) UsersGet(users []int, fields string) ([]*User, error) {
	userIds := ArrayToStr(users)
	params := url.Values{}
	params.Set("user_ids", userIds)
	if fields != "" {
		params.Set("fields", fields)
	}

	resp, err := client.MakeRequest("users.get", params)
	if err != nil {
		return nil, err
	}

	var userList []*User
	json.Unmarshal(resp.Response, &userList)

	return userList, nil
}

// func (client *VKClient) UsersGet(users []int) ([]*User, error) {
// 	idsString := ArrayToStr(users)
// 	params := url.Values{}
// 	params.Set("user_ids", idsString)
// 	params.Set("fields", userFields)

// 	resp, err := client.MakeRequest("users.get", params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var userList []*User
// 	json.Unmarshal(resp.Response, &userList)

// 	return userList, nil
// }
