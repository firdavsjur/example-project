package jsonDb

import (
	"app/models"
	"encoding/json"
	"os"
)

type commissionRepo struct {
	fileName string
}

func NewCommissionRepo(fileName string) *commissionRepo {
	return &commissionRepo{
		fileName: fileName,
	}
}

func (c *commissionRepo) AddCommission(req *models.Commission) error {
	data, err := os.ReadFile(c.fileName)
	if err != nil {
		return err
	}

	var commissions []models.Commission
	err = json.Unmarshal(data, &commissions)
	if err != nil {
		return err
	}

	commissions = append(commissions, models.Commission{
		Balance: req.Balance,
	})

	body, err := json.MarshalIndent(commissions, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(c.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
