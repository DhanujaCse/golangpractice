package model

import "golangprg/views"

func ReadData() ([]views.PostData, error) {

	rows, err := con.Query("SELECT * FROM table1")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	todos := []views.PostData{}
	for rows.Next() {
		data := views.PostData{}
		rows.Scan(&data.FirstName, &data.LastName)
		todos = append(todos, data)
	}
	return todos, nil

}
func ReadSpecificData(name string) ([]views.PostData, error) {

	rows, err := con.Query("SELECT * FROM table1 WHERE firstname=?", name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	todos := []views.PostData{}
	for rows.Next() {
		data := views.PostData{}
		rows.Scan(&data.FirstName, &data.LastName)
		todos = append(todos, data)
	}
	return todos, nil

}
func DeleteSpecificData1(name string) error {

	rows, err := con.Query("DELETE FROM table1 WHERE firstname=?", name)
	defer rows.Close()
	if err != nil {
		return err
	}
	todos := []views.PostData{}
	for rows.Next() {
		data := views.PostData{}
		rows.Scan(&data.FirstName, &data.LastName)
		todos = append(todos, data)
	}
	return nil

}
func DeleteSpecificData3(name string) ([]views.PostData, error) {

	rows, err := con.Query("DELETE FROM table1 WHERE firstname=?", name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	todos := []views.PostData{}
	for rows.Next() {
		data := views.PostData{}
		rows.Scan(&data.FirstName, &data.LastName)
		todos = append(todos, data)
	}
	return todos, nil

}
