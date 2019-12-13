package postgres

import (
	"database/sql"

	"github.com/nagamocha3000/contact_list/pkg/models"
)

//ContactModel ...
type ContactModel struct {
	DB *sql.DB
}

//Insert ... TODO
func (m *ContactModel) Insert(name, phone, address string, favorites []string) (int, error) {
	return -1, nil
}

//GetSingle ... TODO
func (m *ContactModel) GetSingle(id int) (*models.Contact, error) {
	return nil, nil
}

//GetAll ...
func (m *ContactModel) GetAll() ([]*models.Contact, error) {
	stmt := "select id, name, phone, address, favorites from contacts"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	contacts := []*models.Contact{}
	for rows.Next() {
		c := &models.Contact{}
		err = rows.Scan(&c.ID, &c.Name, &c.Phone, &c.Address, &c.Favourites)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return contacts, nil
}
