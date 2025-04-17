package service

import (
	"dedobot/internal/models"
	"dedobot/internal/repositories"
	"fmt"
	"math/rand"
	"time"
)

type SkufService struct {
	repo *repositories.SkufRepo
}

func NewSkufService(repo *repositories.SkufRepo) *SkufService {
	return &SkufService{repo}
}

func (s *SkufService) InitSkuf(userID int64) (string, error) {
	_, err := s.repo.GetByUserID(userID)
	if err == nil {
		return "You already have a skuf!", nil
	}
	skuf := &models.Skuf{UserID: userID, Weight: 1}
	err = s.repo.Create(skuf)
	return "Your skuf is born! Initial weight: 1kg", err
}

func (s *SkufService) FeedSkuf(userID int64) (string, error) {
	skuf, err := s.repo.GetByUserID(userID)
	if err != nil {
		return "You need to /init your skuf first!", nil
	}
	rand.Seed(time.Now().UnixNano())
	gain := float64(rand.Intn(8) + 1)
	skuf.Weight += gain
	err = s.repo.Update(skuf)
	return fmt.Sprintf("You fed your skuf! 🍖 It gained %.2f kg. Total weight: %.2f kg.", gain, skuf.Weight), err
}

func (s *SkufService) RenameSkuf(userID int64, newName string) (string, error) {
	skuf, err := s.repo.GetByUserID(userID)
	if err != nil {
		return "You need to /init your skuf first!", nil
	}
	oldName := skuf.Name
	skuf.Name = newName
	err = s.repo.Update(skuf)
	return fmt.Sprintf("Your Skuf has been renamed from %s to %s!", oldName, newName), err
}

func (s *SkufService) ListSkufs() (string, error) {
	skufs, err := s.repo.GetAll()
	if err != nil {
		return "", err
	}
	if len(skufs) == 0 {
		return "No Skufs found yet. Be the first to /init one!", nil
	}

	result := "🐽 *List of Skufs*:\n"
	for i, skuf := range skufs {
		result += fmt.Sprintf("%d. %s — %.2f kg (user ID: %d)\n", i+1, skuf.Name, skuf.Weight, skuf.UserID)
	}
	return result, nil
}
