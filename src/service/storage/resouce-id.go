package storage

import "strings"

func (d *DataBase) SaveResourceIDs(resourceIDs []*ResourceId) {
	for _, rID := range resourceIDs {
		rID.ID = strings.ToLower(rID.ID)
		existing := d.FetchResourceIDByName(rID.Name)
		if existing.Name != "" {
			if err := d.updateResourceID(rID); err != nil {
				continue
			}
		} else {
			if err := d.saveResourceID(rID); err != nil {
				continue
			}
		}
	}
}

func (d *DataBase) FetchResourceID(resourceId string) (rID ResourceId) {
	d.db.Model(ResourceId{}).Where("id = ?", resourceId).First(&rID)
	return rID
}

func (d *DataBase) saveResourceID(resourceID *ResourceId) error {
	if err := d.db.Model(ResourceId{}).Create(resourceID).Error; err != nil {
		return err
	}
	return nil
}

func (d *DataBase) updateResourceID(rID *ResourceId) error {
	if err := d.db.Model(ResourceId{}).Where("name = ?", rID.Name).Update(&rID).Error; err != nil {
		return err
	}
	return nil
}

func (d *DataBase) FetchResourceIDByName(name string) (rID ResourceId) {
	d.db.Model(ResourceId{}).Where("name = ?", name).First(&rID)
	return rID
}
