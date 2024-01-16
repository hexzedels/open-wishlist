package sdk

const (
	EnvToken    = "TOKEN"
	EnvPostgres = "POSTGRES"
	EnvDebug    = "DEBUG"
)

const (
	CommandStart = "start"
)

const (
	MessageStart           = "Рад тебя видеть! Создадим новый вишлист?"
	MessageWishlistNew     = "Напиши название нового вишлиста ниже"
	MessageWishlistCreated = `Создал новый вишлист "%s"`
	MessageWishlistList    = "Вот список твоих вишлистов.\nВыбери тот, с которым сейчам будешь работать. Ты всегда сможешь изменить его в этом меню"
	MessageWishlistSet     = `Поставил вишлист "%s" текущим`
	MessageItemNew         = "Добавляем новый предмет в вишлист\nНапиши название предмета ниже"
	MessageItemURL         = "У тебя есть ссылка на сайт, где можно купить добавляемый предмет?"
	MessageCreatedItem     = "Добавил"
	MessagedListItems      = `
	| **N** | **id** | **name** |
	|-------|--------|----------|
	`
	MessageTableFormat = "| %d | %d | %s |\n"
)

const (
	StateHome           = "home"
	StateWishlistNew    = "wishlist_new"
	StateWishlistChoose = "wishlist_choose"
	StateItemNew        = "item_new"
	StateItemName       = "item_name"
)

const (
	ButtonNewItem          = "📌 Добавить айтем в текущий вишлист"
	ButtonListItems        = "📝 Вывести текущий вишлист"
	ButtonNewWishlist      = "🎁 Создать новый вишлист"
	ButtonExistingWishlist = "🗂 Сменить вишлист"
)

const (
	CallbackWishlist  = "wishlist"
	CallbackItemURLNo = "item_url_no"
)
