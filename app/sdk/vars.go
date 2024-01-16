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
)

const (
	StateHome           = "home"
	StateWishlistNew    = "wishlist_new"
	StateWishlistChoose = "wishlist_choose"
)

const (
	ButtonNewItem          = "📌 Добавить айтем в текущий вишлист"
	ButtonNewWishlist      = "📋 Создать новый вишлист"
	ButtonExistingWishlist = "🗄 Сменить вишлист"
)

const (
	CallbackWishlist = "wishlist"
)
