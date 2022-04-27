package usecase

import (
	"math/rand"
	"twitter-lofi/entities"
	"twitter-lofi/services"
)

type LofiUseCase struct {
	provider services.SearchLofiOnYoutube
	receiver services.PostLofi
	listLofi []entities.Lofi
}

func (l *LofiUseCase) Execute(search string) entities.Lofi {
	provider := services.SearchLofiOnYoutube{}
	l.addLofiProvider(provider)

	receiver := services.PostLofi{}
	l.addReceiverLofi(receiver)

	lofiSelected := l.selectLofi(search)

	l.receiver.Post(lofiSelected)

	return lofiSelected
}

func (l *LofiUseCase) addLofiProvider(provider services.SearchLofiOnYoutube) {
	l.provider = provider
}

func (l *LofiUseCase) addReceiverLofi(receiver services.PostLofi) {
	l.receiver = receiver
}

func (l *LofiUseCase) selectLofi(search string) entities.Lofi {
	l.listLofi = l.provider.Search(search)

	if len(l.listLofi) == 0 {
		return entities.Lofi{}
	}

	randomIndex := rand.Intn(len(l.listLofi))

	return l.listLofi[randomIndex]
}
