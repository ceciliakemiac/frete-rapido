package database

import (
	"fmt"
	"strconv"

	"github.com/ceciliakemiac/frete-rapido/model"
)

func (db *Database) GetMetrics(filters map[string]string) (*model.Metrics, error) {
	lastQuotes := filters["last_quotes"]
	lastQuotesNum, _ := strconv.Atoi(lastQuotes)

	transportersMetrics, err := db.getTransportersMetrics(lastQuotesNum)
	if err != nil {
		return nil, err
	}

	minFreight, err := db.getMinimumPrice(lastQuotesNum)
	if err != nil {
		return nil, err
	}

	maxFreight, err := db.getMaximumPrice(lastQuotesNum)
	if err != nil {
		return nil, err
	}

	metrics := &model.Metrics{
		Fretes:          transportersMetrics,
		FreteMaisBarato: minFreight,
		FreteMaisCaro:   maxFreight,
	}

	return metrics, nil
}

func (db *Database) getTransportersMetrics(lastQuotes int) ([]model.Metric, error) {
	var metrics []model.Metric
	var query = queryGetMetrics

	if lastQuotes > 0 {
		query = fmt.Sprintf(queryGetMetricsLastQuotes, lastQuotes)
	}

	if err := db.PG.Raw(query).Scan(&metrics).Error; err != nil {
		return nil, fmt.Errorf("Error Get Transporters Metrics: %v", err)
	}

	return metrics, nil
}

func (db *Database) getMinimumPrice(lastQuotes int) (model.ValueFreight, error) {
	var minFreight model.ValueFreight
	query := fmt.Sprintf(queryPrice, "min")

	if lastQuotes > 0 {
		query = fmt.Sprintf(queryPriceLastQuotes, lastQuotes, "min", lastQuotes)
	}

	if err := db.PG.Raw(query).Scan(&minFreight).Error; err != nil {
		return model.ValueFreight{}, fmt.Errorf("Error Get Minimum Freight Price: %v", err)
	}

	return minFreight, nil
}

func (db *Database) getMaximumPrice(lastQuotes int) (model.ValueFreight, error) {
	var maxFreight model.ValueFreight
	query := fmt.Sprintf(queryPrice, "max")

	if lastQuotes > 0 {
		query = fmt.Sprintf(queryPriceLastQuotes, lastQuotes, "max", lastQuotes)
	}

	if err := db.PG.Raw(query).Scan(&maxFreight).Error; err != nil {
		return model.ValueFreight{}, fmt.Errorf("Error Get Maximum Freight Price: %v", err)
	}

	return maxFreight, nil
}
