package pets

type Pet interface {
	GetName() string
	Feed(food string) string
	GiveAttention(activity string) string
	IsHungry() bool
}
