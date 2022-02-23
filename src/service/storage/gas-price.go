package storage

func (d *DataBase) SaveGasPriceInfo(priceLogs []*GasPrice) {
	for _, priceLog := range priceLogs {
		previousLog := d.GetGasPrice(priceLog.ChainName)
		if previousLog.UpdateTime == 0 {
			if e := d.savePrice(priceLog); e != nil {
				continue
			}
		} else if priceLog.UpdateTime > previousLog.UpdateTime {
			if e := d.updatePrice(priceLog); e != nil {
				continue
			}
		}
	}
}

func (d *DataBase) GetGasPrice(name string) (priceLog GasPrice) {
	d.db.Model(GasPrice{}).Where("chain_name = ?", name).Order("update_time desc").First(&priceLog)
	return priceLog
}

func (d *DataBase) updatePrice(priceLog *GasPrice) error {
	if err := d.db.Model(GasPrice{}).Where("chain_name = ?", priceLog.ChainName).Update(&priceLog).Error; err != nil {
		return err
	}
	return nil
}

func (d *DataBase) savePrice(priceLog *GasPrice) error {
	if err := d.db.Model(GasPrice{}).Create(&priceLog).Error; err != nil {
		return err
	}
	return nil
}
