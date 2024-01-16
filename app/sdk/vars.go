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
	MessageStart        = "Рад тебя видеть! Создадим новый вишлист?"
	MessageWishlistNew  = "Напиши название нового вишлиста ниже"
	MessageWishlistList = "Вот список твоих вишлистов.\nВыбери тот, с которым сейчам будешь работать. Ты всегда сможешь изменить его в этом меню"
)

const (
	StateHome        = "home"
	StateWishlistNew = "wishlist_new"
)

const (
	ButtonNewItem          = "📌 Добавить айтем в текущий вишлист"
	ButtonNewWishlist      = "📋 Создать новый вишлист"
	ButtonExistingWishlist = "🗄 Сменить вишлист"
)

const (
	CallbackWishlist = "wishlist"
)
