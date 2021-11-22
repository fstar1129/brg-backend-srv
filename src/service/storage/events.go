package storage

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*
1. CREATE event IN ./tx.go|ConfirmWorkerTx(...)
2. GET event
3. UPDATE event
*/

// GetEventsByTypeAndStatuses ...
func (d *DataBase) GetEventsByTypeAndStatuses(statuses []EventStatus) []*Event {
	swaps := make([]*Event, 0)
	if err := d.db.Where("status in (?)", statuses).
		Find(&swaps).Error; err != nil {
		return nil
	}

	return swaps
}

// UpdateEventStatus ...
func (d *DataBase) UpdateEventStatus(event *Event, status EventStatus) {
	toUpdate := map[string]interface{}{
		"status":      status,
		"update_time": time.Now().Unix(),
	}

	d.db.Model(event).Update(toUpdate)
}

// CompensateNewEvent ...
func (d *DataBase) CompensateNewEvent(chain string, tx *gorm.DB, newEvents []*Event) error {
	for _, event := range newEvents {
		txLogs, err := d.GetConfirmedTxsLog(chain, event, tx)
		if err != nil {
			continue
		}

		if len(txLogs) == 0 {
			continue
		}

		if err = d.ConfirmTx(tx, txLogs[0]); err != nil {
			return err
		}
	}

	return nil
}

// UpdateEventStatusWhenConfirmTx ...
func (d *DataBase) UpdateEventStatusWhenConfirmTx(tx *gorm.DB, txLog *TxLog,
	inStatuses []EventStatus, notInStatuses []EventStatus, updateStatus EventStatus) error {
	query := tx.Model(Event{})

	query = query.Where("event_id = ?", txLog.EventID)

	if len(inStatuses) != 0 {
		query = query.Where("status in (?)", inStatuses)
	}

	if len(notInStatuses) != 0 {
		query = query.Where("status not in (?)", notInStatuses)
	}

	toUpdate := map[string]interface{}{
		"status":      updateStatus,
		"update_time": time.Now().Unix(),
	}

	return query.Updates(toUpdate).Error
}
