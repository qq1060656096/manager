package gormm

import "fmt"

type DataSourceName struct {

}

func (o *DataSourceName) GetMysql(host, port, user, pass, dbName, charset string) string {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		dbName,
		charset,
	)
	return dataSourceName
}

func (o *DataSourceName) GetSqlite(path string) string {
	dataSourceName := path
	return dataSourceName
}

func (o *DataSourceName) GetPostgres(host, port, user, pass, dbName string) string {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		host,
		port,
		user,
		dbName,
		pass,
	)
	return dataSourceName
}

func (o *DataSourceName) GetSqlServer(host, port, user, pass, dbName string) string {
	dataSourceName := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		user,
		pass,
		host,
		port,
		dbName,
	)
	return dataSourceName
}